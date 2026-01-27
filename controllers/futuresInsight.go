package controllers

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"go_binance_futures/feature/api/binance"
	spotbinance "go_binance_futures/spot/api/binance"
	"go_binance_futures/utils"

	spotlib "github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/server/web"
)

type FuturesInsightController struct {
	web.Controller
}

type insightCacheItem struct {
	data      map[string]interface{}
	expiredAt time.Time
}

type insightCacheStore struct {
	mu    sync.RWMutex
	items map[string]insightCacheItem
}

var futuresInsightCache = &insightCacheStore{items: map[string]insightCacheItem{}}

func getInsightCache(key string) (map[string]interface{}, bool) {
	futuresInsightCache.mu.RLock()
	item, ok := futuresInsightCache.items[key]
	futuresInsightCache.mu.RUnlock()
	if !ok || time.Now().After(item.expiredAt) {
		return nil, false
	}
	return item.data, true
}

func setInsightCache(key string, data map[string]interface{}) {
	futuresInsightCache.mu.Lock()
	futuresInsightCache.items[key] = insightCacheItem{
		data:      data,
		expiredAt: time.Now().Add(90 * time.Second),
	}
	futuresInsightCache.mu.Unlock()
}

type OIHistoryPoint struct {
	Time              int64   `json:"time"`
	OpenInterest      float64 `json:"open_interest"`
	OpenInterestValue float64 `json:"open_interest_value"`
}

type RatioPoint struct {
	Time  int64   `json:"time"`
	Ratio float64 `json:"ratio"`
}

type TakerPoint struct {
	Time   int64   `json:"time"`
	BuyVol float64 `json:"buy_vol"`
	SellVol float64 `json:"sell_vol"`
}

type BasisPoint struct {
	Time      int64   `json:"time"`
	Basis     float64 `json:"basis"`
	MarkPrice float64 `json:"mark_price"`
	SpotPrice float64 `json:"spot_price"`
}

type FundingPoint struct {
	Time int64   `json:"time"`
	Rate float64 `json:"rate"`
}

func parseFloat(v string) float64 {
	if v == "" {
		return 0
	}
	val, _ := strconv.ParseFloat(v, 64)
	return val
}

func (ctrl *FuturesInsightController) Get() {
	symbol := strings.ToUpper(ctrl.GetString("symbol"))
	if symbol == "" {
		ctrl.Ctx.Resp(utils.ResJson(400, nil, "symbol required"))
		return
	}
	interval := ctrl.GetString("interval")
	if interval == "" {
		interval = "5m"
	}
	allowed := map[string]bool{
		"5m": true,
		"15m": true,
		"30m": true,
		"1h": true,
		"4h": true,
	}
	if !allowed[interval] {
		interval = "5m"
	}
	limit, err := ctrl.GetInt("limit")
	if err != nil || limit <= 0 {
		limit = 60
	}
	if limit > 200 {
		limit = 200
	}

	cacheKey := fmt.Sprintf("%s|%s|%d", symbol, interval, limit)
	if data, ok := getInsightCache(cacheKey); ok {
		ctrl.Ctx.Resp(map[string]interface{}{
			"code": 200,
			"data": data,
			"msg":  "success",
		})
		return
	}

	var (
		oiHist       = make([]*futures.OpenInterestStatistic, 0)
		topAccount   = make([]*futures.TopLongShortAccountRatio, 0)
		topPosition  = make([]*futures.TopLongShortPositionRatio, 0)
		globalRatio  = make([]*futures.LongShortRatio, 0)
		takerVolume  = make([]*futures.TakerLongShortRatio, 0)
		markKlines   = make([]*futures.Kline, 0)
		spotKlines   = make([]*spotlib.Kline, 0)
		fundingRates = make([]*futures.FundingRate, 0)
	)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		res, err := binance.GetOpenInterestHistory(symbol, interval, limit)
		if err == nil {
			oiHist = res
		} else {
			logs.Error("GetOpenInterestHistory:", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res, err := binance.GetTopLongShortAccountRatio(symbol, interval, limit)
		if err == nil {
			topAccount = res
		} else {
			logs.Error("GetTopLongShortAccountRatio:", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res, err := binance.GetTopLongShortPositionRatio(symbol, interval, limit)
		if err == nil {
			topPosition = res
		} else {
			logs.Error("GetTopLongShortPositionRatio:", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res, err := binance.GetGlobalLongShortAccountRatio(symbol, interval, limit)
		if err == nil {
			globalRatio = res
		} else {
			logs.Error("GetGlobalLongShortAccountRatio:", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res, err := binance.GetTakerLongShortRatio(symbol, interval, limit)
		if err == nil {
			takerVolume = res
		} else {
			logs.Error("GetTakerLongShortRatio:", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res, err := binance.GetMarkPriceKlines(symbol, interval, limit)
		if err == nil {
			markKlines = res
		} else {
			logs.Error("GetMarkPriceKlines:", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res, err := spotbinance.GetKlineData(symbol, interval, limit)
		if err == nil {
			spotKlines = res
		} else {
			logs.Error("GetSpotKlines:", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res, err := binance.GetFundingRateHistory(binance.FundingRateParams{Symbol: symbol, Limit: limit})
		if err == nil {
			fundingRates = res
		} else {
			logs.Error("GetFundingRateHistory:", err)
		}
	}()

	wg.Wait()

	oiPoints := make([]OIHistoryPoint, 0, len(oiHist))
	for _, item := range oiHist {
		oiPoints = append(oiPoints, OIHistoryPoint{
			Time:              item.Timestamp,
			OpenInterest:      parseFloat(item.SumOpenInterest),
			OpenInterestValue: parseFloat(item.SumOpenInterestValue),
		})
	}
	sort.Slice(oiPoints, func(i, j int) bool { return oiPoints[i].Time < oiPoints[j].Time })

	topAccountPoints := make([]RatioPoint, 0, len(topAccount))
	for _, item := range topAccount {
		topAccountPoints = append(topAccountPoints, RatioPoint{
			Time:  int64(item.Timestamp),
			Ratio: parseFloat(item.LongShortRatio),
		})
	}
	sort.Slice(topAccountPoints, func(i, j int) bool { return topAccountPoints[i].Time < topAccountPoints[j].Time })

	topPositionPoints := make([]RatioPoint, 0, len(topPosition))
	for _, item := range topPosition {
		topPositionPoints = append(topPositionPoints, RatioPoint{
			Time:  int64(item.Timestamp),
			Ratio: parseFloat(item.LongShortRatio),
		})
	}
	sort.Slice(topPositionPoints, func(i, j int) bool { return topPositionPoints[i].Time < topPositionPoints[j].Time })

	globalRatioPoints := make([]RatioPoint, 0, len(globalRatio))
	for _, item := range globalRatio {
		globalRatioPoints = append(globalRatioPoints, RatioPoint{
			Time:  item.Timestamp,
			Ratio: parseFloat(item.LongShortRatio),
		})
	}
	sort.Slice(globalRatioPoints, func(i, j int) bool { return globalRatioPoints[i].Time < globalRatioPoints[j].Time })

	takerPoints := make([]TakerPoint, 0, len(takerVolume))
	for _, item := range takerVolume {
		takerPoints = append(takerPoints, TakerPoint{
			Time:   int64(item.Timestamp),
			BuyVol: parseFloat(item.BuyVol),
			SellVol: parseFloat(item.SellVol),
		})
	}
	sort.Slice(takerPoints, func(i, j int) bool { return takerPoints[i].Time < takerPoints[j].Time })

	fundingPoints := make([]FundingPoint, 0, len(fundingRates))
	for _, item := range fundingRates {
		fundingPoints = append(fundingPoints, FundingPoint{
			Time: item.FundingTime,
			Rate: parseFloat(item.FundingRate),
		})
	}
	sort.Slice(fundingPoints, func(i, j int) bool { return fundingPoints[i].Time < fundingPoints[j].Time })

	spotMap := make(map[int64]float64, len(spotKlines))
	for _, k := range spotKlines {
		spotMap[k.OpenTime] = parseFloat(k.Close)
	}
	basisPoints := make([]BasisPoint, 0, len(markKlines))
	for _, mk := range markKlines {
		spotPrice, ok := spotMap[mk.OpenTime]
		if !ok {
			continue
		}
		markPrice := parseFloat(mk.Close)
		basisPoints = append(basisPoints, BasisPoint{
			Time:      mk.OpenTime,
			MarkPrice: markPrice,
			SpotPrice: spotPrice,
			Basis:     markPrice - spotPrice,
		})
	}
	sort.Slice(basisPoints, func(i, j int) bool { return basisPoints[i].Time < basisPoints[j].Time })

	data := map[string]interface{}{
		"symbol":              symbol,
		"interval":            interval,
		"oi_hist":             oiPoints,
		"top_account_ratio":   topAccountPoints,
		"top_position_ratio":  topPositionPoints,
		"global_account_ratio": globalRatioPoints,
		"taker_volume":        takerPoints,
		"basis":               basisPoints,
		"funding_rate":        fundingPoints,
		"oi_marketcap_ratio":  []interface{}{},
	}

	setInsightCache(cacheKey, data)

	ctrl.Ctx.Resp(map[string]interface{}{
		"code": 200,
		"data": data,
		"msg":  "success",
	})
}
