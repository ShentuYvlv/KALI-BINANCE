package feature

import (
	"encoding/json"
	"go_binance_futures/feature/api/binance"
	"go_binance_futures/models"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2/futures"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
)

type scanConfig struct {
	Enabled        bool
	RunOnStart     bool
	TopN           int
	MinQuoteVolume float64
	MaxWorkers     int
	Period         string
	TsLimit        int
	FundingLimit   int
	UtcHour        int
	UtcMinute      int
	RetryMax       int
	RetryBackoffMs int
	RequestGapMs   int
}

type moverItem struct {
	Symbol             string
	PriceChangePercent float64
	QuoteVolume        float64
	OpenTime           int64
	CloseTime          int64
}

func StartDailyScanScheduler() {
	ranOnce := false
	for {
		cfg := getScanConfig()
		if !cfg.Enabled {
			time.Sleep(time.Hour)
			continue
		}
		if cfg.RunOnStart && !ranOnce {
			scanDate := time.Now().UTC().Format("2006-01-02")
			if err := RunDailyScanPipeline(scanDate, cfg); err != nil {
				logs.Error("daily scan run_on_start error:", err)
			}
			ranOnce = true
		}
		next := nextScanTimeUTC(cfg.UtcHour, cfg.UtcMinute)
		time.Sleep(time.Until(next))
		scanDate := next.UTC().Format("2006-01-02")
		if err := RunDailyScanPipeline(scanDate, cfg); err != nil {
			logs.Error("daily scan error:", err)
		}
	}
}

func RunDailyScanPipeline(scanDate string, cfg scanConfig) error {
	startAt := time.Now()
	logs.Info("daily scan start:", scanDate)
	allowMap, err := loadTradableUSDTPerp(cfg)
	if err != nil {
		return err
	}
	var stats []*futures.PriceChangeStats
	err = withRetry(cfg, func() error {
		var reqErr error
		stats, reqErr = binance.Get24hPriceChangeStats("")
		return reqErr
	})
	if err != nil {
		return err
	}
	gainers, losers := pickTopMovers(stats, allowMap, cfg.TopN, cfg.MinQuoteVolume)
	if len(gainers) == 0 && len(losers) == 0 {
		logs.Error("daily scan empty after filter")
		return nil
	}

	if err := upsertDailyScan(scanDate, gainers, losers); err != nil {
		logs.Error("upsert daily_scan error:", err)
	}
	logs.Info("daily scan list: date=%s gainers=%d losers=%d", scanDate, len(gainers), len(losers))

	symbols := buildScanSymbolList(gainers, losers)
	featureStats := runFeatureCollection(scanDate, symbols, cfg)

	backtestOK, backtestFail, err := backtestT1(scanDate, cfg)
	if err != nil {
		logs.Error("backtest t1 error:", err)
	}
	logs.Info("daily scan summary: date=%s symbols=%d feature_ok=%d feature_fail=%d backtest_ok=%d backtest_fail=%d cost=%s",
		scanDate,
		len(symbols),
		featureStats.success,
		featureStats.fail,
		backtestOK,
		backtestFail,
		time.Since(startAt).String(),
	)
	return nil
}

func getScanConfig() scanConfig {
	cfg := scanConfig{
		Enabled:        true,
		RunOnStart:     false,
		TopN:           10,
		MinQuoteVolume: 50000000,
		MaxWorkers:     5,
		Period:         "1d",
		TsLimit:        30,
		FundingLimit:   30,
		UtcHour:        0,
		UtcMinute:      5,
		RetryMax:       4,
		RetryBackoffMs: 800,
		RequestGapMs:   0,
	}
	if v, err := config.Int("scan::enabled"); err == nil {
		cfg.Enabled = v == 1
	}
	if v, err := config.Int("scan::run_on_start"); err == nil {
		cfg.RunOnStart = v == 1
	}
	if v, err := config.Int("scan::top_n"); err == nil && v > 0 {
		cfg.TopN = v
	}
	if v, err := config.Int("scan::min_quote_volume"); err == nil && v > 0 {
		cfg.MinQuoteVolume = float64(v)
	}
	if v, err := config.Int("scan::max_workers"); err == nil && v > 0 {
		cfg.MaxWorkers = v
	}
	if v, err := config.String("scan::period"); err == nil && v != "" {
		cfg.Period = v
	}
	if v, err := config.Int("scan::ts_limit"); err == nil && v > 0 {
		cfg.TsLimit = v
	}
	if v, err := config.Int("scan::funding_limit"); err == nil && v > 0 {
		cfg.FundingLimit = v
	}
	if v, err := config.Int("scan::utc_hour"); err == nil {
		cfg.UtcHour = v
	}
	if v, err := config.Int("scan::utc_minute"); err == nil {
		cfg.UtcMinute = v
	}
	if v, err := config.Int("scan::retry_max"); err == nil && v >= 0 {
		cfg.RetryMax = v
	}
	if v, err := config.Int("scan::retry_backoff_ms"); err == nil && v > 0 {
		cfg.RetryBackoffMs = v
	}
	if v, err := config.Int("scan::request_gap_ms"); err == nil && v >= 0 {
		cfg.RequestGapMs = v
	}
	return cfg
}

func nextScanTimeUTC(hour int, minute int) time.Time {
	now := time.Now().UTC()
	run := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, time.UTC)
	if !run.After(now) {
		run = run.Add(24 * time.Hour)
	}
	return run
}

func loadTradableUSDTPerp(cfg scanConfig) (map[string]bool, error) {
	var info *futures.ExchangeInfo
	err := withRetry(cfg, func() error {
		var reqErr error
		info, reqErr = binance.GetExchangeInfo()
		return reqErr
	})
	if err != nil {
		return nil, err
	}
	allow := make(map[string]bool)
	for _, s := range info.Symbols {
		if s.Status != "TRADING" {
			continue
		}
		if s.QuoteAsset != "USDT" {
			continue
		}
		if s.ContractType != futures.ContractTypePerpetual {
			continue
		}
		allow[s.Symbol] = true
	}
	return allow, nil
}

func pickTopMovers(stats []*futures.PriceChangeStats, allow map[string]bool, topN int, minQuote float64) ([]moverItem, []moverItem) {
	items := make([]moverItem, 0, len(stats))
	for _, s := range stats {
		if s == nil || s.Symbol == "" {
			continue
		}
		if !allow[s.Symbol] {
			continue
		}
		pct, err1 := strconv.ParseFloat(s.PriceChangePercent, 64)
		vol, err2 := strconv.ParseFloat(s.QuoteVolume, 64)
		if err1 != nil || err2 != nil {
			continue
		}
		if vol < minQuote {
			continue
		}
		items = append(items, moverItem{
			Symbol:             s.Symbol,
			PriceChangePercent: pct,
			QuoteVolume:        vol,
			OpenTime:           s.OpenTime,
			CloseTime:          s.CloseTime,
		})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].PriceChangePercent > items[j].PriceChangePercent
	})
	gainers := append([]moverItem{}, items...)
	losers := append([]moverItem{}, items...)
	sort.Slice(losers, func(i, j int) bool {
		return losers[i].PriceChangePercent < losers[j].PriceChangePercent
	})
	if len(gainers) > topN {
		gainers = gainers[:topN]
	}
	if len(losers) > topN {
		losers = losers[:topN]
	}
	return gainers, losers
}

func buildScanSymbolList(gainers []moverItem, losers []moverItem) []string {
	uniq := make(map[string]bool)
	list := make([]string, 0, len(gainers)+len(losers))
	for _, item := range gainers {
		if !uniq[item.Symbol] {
			uniq[item.Symbol] = true
			list = append(list, item.Symbol)
		}
	}
	for _, item := range losers {
		if !uniq[item.Symbol] {
			uniq[item.Symbol] = true
			list = append(list, item.Symbol)
		}
	}
	return list
}

func upsertDailyScan(scanDate string, gainers []moverItem, losers []moverItem) error {
	o := orm.NewOrm()
	now := time.Now().Unix() * 1000
	for idx, item := range gainers {
		record := models.DailyScan{
			ScanDate:           scanDate,
			Symbol:             item.Symbol,
			Direction:          "gain",
			Rank:               int64(idx + 1),
			PriceChangePercent: item.PriceChangePercent,
			QuoteVolume:        item.QuoteVolume,
			OpenTime:           item.OpenTime,
			CloseTime:          item.CloseTime,
			CreateTime:         now,
		}
		if err := upsertDailyScanRecord(o, record); err != nil {
			logs.Error("upsert daily_scan gain:", item.Symbol, err)
		}
	}
	for idx, item := range losers {
		record := models.DailyScan{
			ScanDate:           scanDate,
			Symbol:             item.Symbol,
			Direction:          "lose",
			Rank:               int64(idx + 1),
			PriceChangePercent: item.PriceChangePercent,
			QuoteVolume:        item.QuoteVolume,
			OpenTime:           item.OpenTime,
			CloseTime:          item.CloseTime,
			CreateTime:         now,
		}
		if err := upsertDailyScanRecord(o, record); err != nil {
			logs.Error("upsert daily_scan lose:", item.Symbol, err)
		}
	}
	return nil
}

func upsertDailyScanRecord(o orm.Ormer, record models.DailyScan) error {
	var exist models.DailyScan
	err := o.QueryTable(record.TableName()).
		Filter("scan_date", record.ScanDate).
		Filter("symbol", record.Symbol).
		One(&exist)
	if err == nil {
		exist.Direction = record.Direction
		exist.Rank = record.Rank
		exist.PriceChangePercent = record.PriceChangePercent
		exist.QuoteVolume = record.QuoteVolume
		exist.OpenTime = record.OpenTime
		exist.CloseTime = record.CloseTime
		_, err = o.Update(&exist, "direction", "rank", "price_change_percent", "quote_volume", "open_time", "close_time")
		return err
	}
	if err == orm.ErrNoRows {
		_, err = o.Insert(&record)
		return err
	}
	return err
}

type featureRunStats struct {
	success int
	fail    int
}

func runFeatureCollection(scanDate string, symbols []string, cfg scanConfig) featureRunStats {
	stats := featureRunStats{}
	if len(symbols) == 0 {
		return stats
	}
	workerCount := cfg.MaxWorkers
	if workerCount < 1 {
		workerCount = 1
	}
	jobs := make(chan string)
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for symbol := range jobs {
				if err := collectSymbolFeatures(scanDate, symbol, cfg); err != nil {
					logs.Error("collect symbol features:", symbol, err)
					mu.Lock()
					stats.fail++
					mu.Unlock()
				} else {
					mu.Lock()
					stats.success++
					mu.Unlock()
				}
			}
		}()
	}
	for _, symbol := range symbols {
		jobs <- symbol
	}
	close(jobs)
	wg.Wait()
	return stats
}

func collectSymbolFeatures(scanDate string, symbol string, cfg scanConfig) error {
	now := time.Now().Unix() * 1000
	o := orm.NewOrm()
	feature := models.DailyFeature{
		ScanDate:   scanDate,
		Symbol:     symbol,
		CreateTime: now,
	}
	raw := map[string]any{}

	var premium []*futures.PremiumIndex
	err := withRetry(cfg, func() error {
		var reqErr error
		premium, reqErr = binance.GetFundingRate(binance.FundingRateParams{Symbol: symbol})
		return reqErr
	})
	if err != nil {
		logs.Error("premiumIndex:", symbol, err)
	} else if len(premium) > 0 {
		item := premium[0]
		feature.MarkPrice = parseFloat(item.MarkPrice)
		feature.IndexPrice = parseFloat(item.IndexPrice)
		feature.LastFundingRate = parseFloat(item.LastFundingRate)
		feature.NextFundingTime = item.NextFundingTime
		raw["premiumIndex"] = item
	}

	var openInterest *futures.OpenInterest
	err = withRetry(cfg, func() error {
		var reqErr error
		openInterest, reqErr = binance.GetOpenInterest(symbol)
		return reqErr
	})
	if err != nil {
		logs.Error("openInterest:", symbol, err)
	} else if openInterest != nil {
		feature.OpenInterest = parseFloat(openInterest.OpenInterest)
		raw["openInterest"] = openInterest
	}

	var openInterestHist []*futures.OpenInterestStatistic
	err = withRetry(cfg, func() error {
		var reqErr error
		openInterestHist, reqErr = binance.GetOpenInterestHistory(symbol, cfg.Period, cfg.TsLimit)
		return reqErr
	})
	if err != nil {
		logs.Error("openInterestHist:", symbol, err)
	} else {
		raw["openInterestHist"] = openInterestHist
		appendOpenInterestTimeseries(o, symbol, openInterestHist, now)
	}

	var topPos []*futures.TopLongShortPositionRatio
	err = withRetry(cfg, func() error {
		var reqErr error
		topPos, reqErr = binance.GetTopLongShortPositionRatio(symbol, cfg.Period, cfg.TsLimit)
		return reqErr
	})
	if err != nil {
		logs.Error("topLongShortPosition:", symbol, err)
	} else {
		raw["topLongShortPositionRatio"] = topPos
		feature.TopLongShortPositionRatio = lastRatioValue(topPos)
		appendTopLongShortPositionTimeseries(o, symbol, topPos, now)
	}

	var topAcc []*futures.TopLongShortAccountRatio
	err = withRetry(cfg, func() error {
		var reqErr error
		topAcc, reqErr = binance.GetTopLongShortAccountRatio(symbol, cfg.Period, cfg.TsLimit)
		return reqErr
	})
	if err != nil {
		logs.Error("topLongShortAccount:", symbol, err)
	} else {
		raw["topLongShortAccountRatio"] = topAcc
		feature.TopLongShortAccountRatio = lastRatioValue(topAcc)
		appendTopLongShortAccountTimeseries(o, symbol, topAcc, now)
	}

	var globalRatio []*futures.LongShortRatio
	err = withRetry(cfg, func() error {
		var reqErr error
		globalRatio, reqErr = binance.GetGlobalLongShortAccountRatio(symbol, cfg.Period, cfg.TsLimit)
		return reqErr
	})
	if err != nil {
		logs.Error("globalLongShortRatio:", symbol, err)
	} else {
		raw["globalLongShortRatio"] = globalRatio
		feature.GlobalLongShortRatio = lastRatioValue(globalRatio)
		appendGlobalLongShortTimeseries(o, symbol, globalRatio, now)
	}

	var takerRatio []*futures.TakerLongShortRatio
	err = withRetry(cfg, func() error {
		var reqErr error
		takerRatio, reqErr = binance.GetTakerLongShortRatio(symbol, cfg.Period, cfg.TsLimit)
		return reqErr
	})
	if err != nil {
		logs.Error("takerLongShortRatio:", symbol, err)
	} else {
		raw["takerLongShortRatio"] = takerRatio
		feature.TakerBuySellRatio = lastTakerRatioValue(takerRatio)
		appendTakerRatioTimeseries(o, symbol, takerRatio, now)
	}

	var basis []*futures.Basis
	err = withRetry(cfg, func() error {
		var reqErr error
		basis, reqErr = binance.GetBasis(symbol, "PERPETUAL", cfg.Period, cfg.TsLimit)
		return reqErr
	})
	if err != nil {
		logs.Error("basis:", symbol, err)
	} else {
		raw["basis"] = basis
		if len(basis) > 0 {
			last := basis[len(basis)-1]
			feature.Basis = parseFloat(last.Basis)
			feature.AnnualizedBasisRate = parseFloat(last.AnnualizedBasisRate)
			feature.FuturesPrice = parseFloat(last.FuturesPrice)
		}
		appendBasisTimeseries(o, symbol, basis, now)
	}

	var fundingHist []*futures.FundingRate
	err = withRetry(cfg, func() error {
		var reqErr error
		fundingHist, reqErr = binance.GetFundingRateHistory(binance.FundingRateParams{
			Symbol: symbol,
			Limit:  cfg.FundingLimit,
		})
		return reqErr
	})
	if err != nil {
		logs.Error("fundingRateHistory:", symbol, err)
	} else {
		raw["fundingRateHistory"] = fundingHist
		appendFundingTimeseries(o, symbol, fundingHist, now)
	}

	rawBytes, _ := json.Marshal(raw)
	feature.RawJson = string(rawBytes)
	return upsertDailyFeature(o, feature)
}

func upsertDailyFeature(o orm.Ormer, record models.DailyFeature) error {
	var exist models.DailyFeature
	err := o.QueryTable(record.TableName()).
		Filter("scan_date", record.ScanDate).
		Filter("symbol", record.Symbol).
		One(&exist)
	if err == nil {
		exist.MarkPrice = record.MarkPrice
		exist.IndexPrice = record.IndexPrice
		exist.LastFundingRate = record.LastFundingRate
		exist.NextFundingTime = record.NextFundingTime
		exist.OpenInterest = record.OpenInterest
		exist.TopLongShortPositionRatio = record.TopLongShortPositionRatio
		exist.TopLongShortAccountRatio = record.TopLongShortAccountRatio
		exist.GlobalLongShortRatio = record.GlobalLongShortRatio
		exist.TakerBuySellRatio = record.TakerBuySellRatio
		exist.Basis = record.Basis
		exist.AnnualizedBasisRate = record.AnnualizedBasisRate
		exist.FuturesPrice = record.FuturesPrice
		exist.FundingRateInterval = record.FundingRateInterval
		exist.RawJson = record.RawJson
		_, err = o.Update(&exist,
			"mark_price",
			"index_price",
			"last_funding_rate",
			"next_funding_time",
			"open_interest",
			"top_long_short_position_ratio",
			"top_long_short_account_ratio",
			"global_long_short_ratio",
			"taker_buy_sell_ratio",
			"basis",
			"annualized_basis_rate",
			"futures_price",
			"funding_rate_interval",
			"raw_json",
		)
		return err
	}
	if err == orm.ErrNoRows {
		_, err = o.Insert(&record)
		return err
	}
	return err
}

func backtestT1(scanDate string, cfg scanConfig) (int, int, error) {
	scanTime, err := time.Parse("2006-01-02", scanDate)
	if err != nil {
		return 0, 0, err
	}
	targetDate := scanTime.AddDate(0, 0, -1)
	targetDateStr := targetDate.Format("2006-01-02")
	nextDate := targetDate.AddDate(0, 0, 1)

	o := orm.NewOrm()
	var scans []models.DailyScan
	_, err = o.QueryTable("daily_scan").Filter("scan_date", targetDateStr).All(&scans)
	if err != nil || len(scans) == 0 {
		return 0, 0, nil
	}
	okCount := 0
	failCount := 0
	for _, item := range scans {
		open, close, ret, err := calcT1Return(item.Symbol, nextDate, cfg)
		if err != nil {
			logs.Error("calc t1 return:", item.Symbol, err)
			failCount++
			continue
		}
		record := models.BacktestT1{
			ScanDate:   targetDateStr,
			Symbol:     item.Symbol,
			Direction:  item.Direction,
			Rank:       item.Rank,
			T1Open:     open,
			T1Close:    close,
			T1Return:   ret,
			CreateTime: time.Now().Unix() * 1000,
		}
		if err := upsertBacktestT1(o, record); err != nil {
			logs.Error("upsert backtest t1:", item.Symbol, err)
			failCount++
		} else {
			okCount++
		}
	}
	return okCount, failCount, nil
}

func calcT1Return(symbol string, day time.Time, cfg scanConfig) (float64, float64, float64, error) {
	start := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.UTC).Unix() * 1000
	end := start + 24*60*60*1000 - 1
	var klines []*futures.Kline
	err := withRetry(cfg, func() error {
		var reqErr error
		klines, reqErr = binance.GetKlineDataRange(symbol, "1d", 2, start, end)
		return reqErr
	})
	if err != nil || len(klines) == 0 {
		return 0, 0, 0, err
	}
	k := klines[0]
	open := parseFloat(k.Open)
	close := parseFloat(k.Close)
	if open == 0 {
		return 0, 0, 0, nil
	}
	ret := (close / open) - 1
	return open, close, ret, nil
}

func upsertBacktestT1(o orm.Ormer, record models.BacktestT1) error {
	var exist models.BacktestT1
	err := o.QueryTable(record.TableName()).
		Filter("scan_date", record.ScanDate).
		Filter("symbol", record.Symbol).
		One(&exist)
	if err == nil {
		exist.Direction = record.Direction
		exist.Rank = record.Rank
		exist.T1Open = record.T1Open
		exist.T1Close = record.T1Close
		exist.T1Return = record.T1Return
		_, err = o.Update(&exist, "direction", "rank", "t1_open", "t1_close", "t1_return")
		return err
	}
	if err == orm.ErrNoRows {
		_, err = o.Insert(&record)
		return err
	}
	return err
}

func appendOpenInterestTimeseries(o orm.Ormer, symbol string, items []*futures.OpenInterestStatistic, now int64) {
	for _, item := range items {
		if item == nil {
			continue
		}
		openInterest := parseFloat(item.SumOpenInterest)
		openInterestValue := parseFloat(item.SumOpenInterestValue)
		if openInterest != 0 {
			_ = insertTimeseries(o, models.FeatureTimeseries{
				Symbol:     symbol,
				Metric:     "open_interest_sum",
				Timestamp:  item.Timestamp,
				Value:      openInterest,
				CreateTime: now,
			})
		}
		if openInterestValue != 0 {
			_ = insertTimeseries(o, models.FeatureTimeseries{
				Symbol:     symbol,
				Metric:     "open_interest_value",
				Timestamp:  item.Timestamp,
				Value:      openInterestValue,
				CreateTime: now,
			})
		}
	}
}

func appendTopLongShortPositionTimeseries(o orm.Ormer, symbol string, items []*futures.TopLongShortPositionRatio, now int64) {
	for _, item := range items {
		if item == nil {
			continue
		}
		value := parseFloat(item.LongShortRatio)
		if value == 0 {
			continue
		}
		_ = insertTimeseries(o, models.FeatureTimeseries{
			Symbol:     symbol,
			Metric:     "top_long_short_position_ratio",
			Timestamp:  int64(item.Timestamp),
			Value:      value,
			CreateTime: now,
		})
	}
}

func appendTopLongShortAccountTimeseries(o orm.Ormer, symbol string, items []*futures.TopLongShortAccountRatio, now int64) {
	for _, item := range items {
		if item == nil {
			continue
		}
		value := parseFloat(item.LongShortRatio)
		if value == 0 {
			continue
		}
		_ = insertTimeseries(o, models.FeatureTimeseries{
			Symbol:     symbol,
			Metric:     "top_long_short_account_ratio",
			Timestamp:  int64(item.Timestamp),
			Value:      value,
			CreateTime: now,
		})
	}
}

func appendGlobalLongShortTimeseries(o orm.Ormer, symbol string, items []*futures.LongShortRatio, now int64) {
	for _, item := range items {
		if item == nil {
			continue
		}
		value := parseFloat(item.LongShortRatio)
		if value == 0 {
			continue
		}
		_ = insertTimeseries(o, models.FeatureTimeseries{
			Symbol:     symbol,
			Metric:     "global_long_short_ratio",
			Timestamp:  item.Timestamp,
			Value:      value,
			CreateTime: now,
		})
	}
}

func appendTakerRatioTimeseries(o orm.Ormer, symbol string, items []*futures.TakerLongShortRatio, now int64) {
	for _, item := range items {
		if item == nil {
			continue
		}
		value := parseFloat(item.BuySellRatio)
		extra := map[string]string{
			"buyVol":  item.BuyVol,
			"sellVol": item.SellVol,
		}
		extraBytes, _ := json.Marshal(extra)
		_ = insertTimeseries(o, models.FeatureTimeseries{
			Symbol:     symbol,
			Metric:     "taker_buy_sell_ratio",
			Timestamp:  int64(item.Timestamp),
			Value:      value,
			ExtraJson:  string(extraBytes),
			CreateTime: now,
		})
	}
}

func appendBasisTimeseries(o orm.Ormer, symbol string, items []*futures.Basis, now int64) {
	for _, item := range items {
		if item == nil {
			continue
		}
		value := parseFloat(item.Basis)
		extraBytes, _ := json.Marshal(item)
		_ = insertTimeseries(o, models.FeatureTimeseries{
			Symbol:     symbol,
			Metric:     "basis",
			Timestamp:  int64(item.Timestamp),
			Value:      value,
			ExtraJson:  string(extraBytes),
			CreateTime: now,
		})
	}
}

func appendFundingTimeseries(o orm.Ormer, symbol string, items []*futures.FundingRate, now int64) {
	for _, item := range items {
		if item == nil {
			continue
		}
		value := parseFloat(item.FundingRate)
		extraBytes, _ := json.Marshal(item)
		_ = insertTimeseries(o, models.FeatureTimeseries{
			Symbol:     symbol,
			Metric:     "funding_rate",
			Timestamp:  item.FundingTime,
			Value:      value,
			ExtraJson:  string(extraBytes),
			CreateTime: now,
		})
	}
}

func insertTimeseries(o orm.Ormer, record models.FeatureTimeseries) error {
	if record.Symbol == "" || record.Metric == "" || record.Timestamp == 0 {
		return nil
	}
	_, err := o.Insert(&record)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "unique") || strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return nil
		}
	}
	return err
}

func lastRatioValue(items interface{}) float64 {
	switch v := items.(type) {
	case []*futures.TopLongShortPositionRatio:
		if len(v) == 0 {
			return 0
		}
		return parseFloat(v[len(v)-1].LongShortRatio)
	case []*futures.TopLongShortAccountRatio:
		if len(v) == 0 {
			return 0
		}
		return parseFloat(v[len(v)-1].LongShortRatio)
	case []*futures.LongShortRatio:
		if len(v) == 0 {
			return 0
		}
		return parseFloat(v[len(v)-1].LongShortRatio)
	default:
		return 0
	}
}

func lastTakerRatioValue(items []*futures.TakerLongShortRatio) float64 {
	if len(items) == 0 {
		return 0
	}
	return parseFloat(items[len(items)-1].BuySellRatio)
}

func parseFloat(val string) float64 {
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0
	}
	return f
}

func withRetry(cfg scanConfig, fn func() error) error {
	if fn == nil {
		return nil
	}
	retryMax := cfg.RetryMax
	if retryMax < 0 {
		retryMax = 0
	}
	var err error
	for attempt := 0; attempt <= retryMax; attempt++ {
		if attempt > 0 {
			time.Sleep(backoffDuration(cfg, attempt))
		}
		sleepGap(cfg)
		err = fn()
		if err == nil {
			return nil
		}
	}
	return err
}

func sleepGap(cfg scanConfig) {
	if cfg.RequestGapMs <= 0 {
		return
	}
	time.Sleep(time.Duration(cfg.RequestGapMs) * time.Millisecond)
}

func backoffDuration(cfg scanConfig, attempt int) time.Duration {
	base := cfg.RetryBackoffMs
	if base <= 0 {
		base = 800
	}
	backoff := time.Duration(base) * time.Millisecond
	for i := 1; i < attempt; i++ {
		backoff *= 2
		if backoff > 10*time.Second {
			backoff = 10 * time.Second
			break
		}
	}
	jitter := time.Duration(time.Now().UnixNano()%200) * time.Millisecond
	return backoff + jitter
}
