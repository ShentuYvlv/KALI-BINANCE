package controllers

import (
	"net/http"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"

	"go_binance_futures/feature/api/binance"
	"go_binance_futures/utils"

	"github.com/adshao/go-binance/v2/futures"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/gorilla/websocket"
)

type FuturesMarketController struct {
	web.Controller
}

type orderLevel struct {
	Price float64 `json:"price"`
	Qty   float64 `json:"qty"`
}

var marketWsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func buildDepthLevels(levels map[string]float64, limit int, desc bool) []orderLevel {
	items := make([]orderLevel, 0, len(levels))
	for priceStr, qty := range levels {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			continue
		}
		items = append(items, orderLevel{Price: price, Qty: qty})
	}
	sort.Slice(items, func(i, j int) bool {
		if desc {
			return items[i].Price > items[j].Price
		}
		return items[i].Price < items[j].Price
	})
	if limit > 0 && len(items) > limit {
		return items[:limit]
	}
	return items
}

func (ctrl *FuturesMarketController) DepthWs() {
	symbol := ctrl.GetString("symbol")
	limit, _ := ctrl.GetInt("limit", 50)
	if symbol == "" {
		ctrl.CustomAbort(400, "symbol required")
		return
	}
	if limit <= 0 {
		limit = 50
	}

	conn, err := marketWsUpgrader.Upgrade(ctrl.Ctx.ResponseWriter, ctrl.Ctx.Request, nil)
	if err != nil {
		logs.Error("depth ws upgrade error:", err)
		return
	}
	defer conn.Close()

	var writeMu sync.Mutex
	var closed int32
	var stopOnce sync.Once
	var stopC chan struct{}
	var doneC chan struct{}

	closeStop := func(stopC chan struct{}) {
		stopOnce.Do(func() {
			if stopC != nil {
				close(stopC)
			}
		})
	}

	bookMu := &sync.Mutex{}
	bids := map[string]float64{}
	asks := map[string]float64{}
	var lastUpdateID int64

	sendDepth := func(msgType string) {
		bookMu.Lock()
		bidLevels := buildDepthLevels(bids, limit, true)
		askLevels := buildDepthLevels(asks, limit, false)
		bookMu.Unlock()
		payload := map[string]interface{}{
			"type": msgType,
			"data": map[string]interface{}{
				"bids": bidLevels,
				"asks": askLevels,
			},
		}
		if atomic.LoadInt32(&closed) == 1 {
			return
		}
		writeMu.Lock()
		if err := conn.WriteJSON(payload); err != nil {
			atomic.StoreInt32(&closed, 1)
			closeStop(stopC)
		}
		writeMu.Unlock()
	}

	// initial snapshot
	snapshot, err := binance.GetDepth(symbol, limit)
	if err != nil {
		logs.Error("depth snapshot error:", err)
		return
	}
	bookMu.Lock()
	bids = map[string]float64{}
	asks = map[string]float64{}
	for _, b := range snapshot.Bids {
		qty, _ := strconv.ParseFloat(b.Quantity, 64)
		if qty > 0 {
			bids[b.Price] = qty
		}
	}
	for _, a := range snapshot.Asks {
		qty, _ := strconv.ParseFloat(a.Quantity, 64)
		if qty > 0 {
			asks[a.Price] = qty
		}
	}
	lastUpdateID = snapshot.LastUpdateID
	bookMu.Unlock()
	sendDepth("snapshot")

	doneC, stopC, err = futures.WsDiffDepthServe(symbol, func(event *futures.WsDepthEvent) {
		bookMu.Lock()
		if event.LastUpdateID <= lastUpdateID {
			bookMu.Unlock()
			return
		}
		if event.FirstUpdateID > lastUpdateID+1 {
			bookMu.Unlock()
			snap, snapErr := binance.GetDepth(symbol, limit)
			if snapErr != nil {
				logs.Error("depth resync error:", snapErr)
				return
			}
			bookMu.Lock()
			bids = map[string]float64{}
			asks = map[string]float64{}
			for _, b := range snap.Bids {
				qty, _ := strconv.ParseFloat(b.Quantity, 64)
				if qty > 0 {
					bids[b.Price] = qty
				}
			}
			for _, a := range snap.Asks {
				qty, _ := strconv.ParseFloat(a.Quantity, 64)
				if qty > 0 {
					asks[a.Price] = qty
				}
			}
			lastUpdateID = snap.LastUpdateID
			bookMu.Unlock()
			sendDepth("snapshot")
			return
		}
		for _, b := range event.Bids {
			qty, _ := strconv.ParseFloat(b.Quantity, 64)
			if qty <= 0 {
				delete(bids, b.Price)
			} else {
				bids[b.Price] = qty
			}
		}
		for _, a := range event.Asks {
			qty, _ := strconv.ParseFloat(a.Quantity, 64)
			if qty <= 0 {
				delete(asks, a.Price)
			} else {
				asks[a.Price] = qty
			}
		}
		lastUpdateID = event.LastUpdateID
		bookMu.Unlock()
		sendDepth("depth")
	}, func(err error) {
		logs.Error("depth ws error:", err)
	})
	if err != nil {
		logs.Error("depth ws start error:", err)
		return
	}

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			atomic.StoreInt32(&closed, 1)
			closeStop(stopC)
			break
		}
	}
	<-doneC
}

func (ctrl *FuturesMarketController) DepthSnapshot() {
	symbol := ctrl.GetString("symbol")
	limit, _ := ctrl.GetInt("limit", 50)
	if symbol == "" {
		ctrl.Ctx.Resp(utils.ResJson(400, nil, "symbol required"))
		return
	}
	if limit <= 0 {
		limit = 50
	}
	if limit > 200 {
		limit = 200
	}

	snapshot, err := binance.GetDepth(symbol, limit)
	if err != nil {
		ctrl.Ctx.Resp(utils.ResJson(400, nil, err.Error()))
		return
	}

	bids := make([]orderLevel, 0, len(snapshot.Bids))
	asks := make([]orderLevel, 0, len(snapshot.Asks))
	for _, b := range snapshot.Bids {
		qty, _ := strconv.ParseFloat(b.Quantity, 64)
		price, _ := strconv.ParseFloat(b.Price, 64)
		bids = append(bids, orderLevel{Price: price, Qty: qty})
	}
	for _, a := range snapshot.Asks {
		qty, _ := strconv.ParseFloat(a.Quantity, 64)
		price, _ := strconv.ParseFloat(a.Price, 64)
		asks = append(asks, orderLevel{Price: price, Qty: qty})
	}
	ctrl.Ctx.Resp(utils.ResJson(200, map[string]interface{}{
		"bids": bids,
		"asks": asks,
	}))
}

func (ctrl *FuturesMarketController) TradesWs() {
	symbol := ctrl.GetString("symbol")
	limit, _ := ctrl.GetInt("limit", 100)
	if symbol == "" {
		ctrl.CustomAbort(400, "symbol required")
		return
	}
	if limit <= 0 {
		limit = 100
	}
	if limit > 200 {
		limit = 200
	}

	conn, err := marketWsUpgrader.Upgrade(ctrl.Ctx.ResponseWriter, ctrl.Ctx.Request, nil)
	if err != nil {
		logs.Error("trades ws upgrade error:", err)
		return
	}
	defer conn.Close()

	var writeMu sync.Mutex
	var closed int32
	var stopOnce sync.Once
	var stopC chan struct{}
	var doneC chan struct{}

	closeStop := func(stopC chan struct{}) {
		stopOnce.Do(func() {
			if stopC != nil {
				close(stopC)
			}
		})
	}

	sendJSON := func(payload map[string]interface{}) {
		if atomic.LoadInt32(&closed) == 1 {
			return
		}
		writeMu.Lock()
		if err := conn.WriteJSON(payload); err != nil {
			atomic.StoreInt32(&closed, 1)
			closeStop(stopC)
		}
		writeMu.Unlock()
	}

	trades, err := binance.GetRecentTrades(symbol, limit)
	if err == nil {
		items := make([]map[string]interface{}, 0, len(trades))
		for i := len(trades) - 1; i >= 0; i-- { // latest first
			t := trades[i]
			items = append(items, map[string]interface{}{
				"id":     t.ID,
				"price":  t.Price,
				"qty":    t.Quantity,
				"time":   t.Time,
				"maker":  t.IsBuyerMaker,
				"symbol": symbol,
			})
		}
		sendJSON(map[string]interface{}{
			"type": "snapshot",
			"data": items,
		})
	}

	doneC, stopC, err = futures.WsAggTradeServe(symbol, func(event *futures.WsAggTradeEvent) {
		item := map[string]interface{}{
			"id":     event.AggregateTradeID,
			"price":  event.Price,
			"qty":    event.Quantity,
			"time":   event.TradeTime,
			"maker":  event.Maker,
			"symbol": event.Symbol,
		}
		sendJSON(map[string]interface{}{
			"type": "update",
			"data": item,
		})
	}, func(err error) {
		logs.Error("trades ws error:", err)
	})
	if err != nil {
		logs.Error("trades ws start error:", err)
		return
	}

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			atomic.StoreInt32(&closed, 1)
			closeStop(stopC)
			break
		}
	}
	<-doneC
}

func (ctrl *FuturesMarketController) TradesSnapshot() {
	symbol := ctrl.GetString("symbol")
	limit, _ := ctrl.GetInt("limit", 100)
	if symbol == "" {
		ctrl.Ctx.Resp(utils.ResJson(400, nil, "symbol required"))
		return
	}
	if limit <= 0 {
		limit = 100
	}
	if limit > 200 {
		limit = 200
	}

	trades, err := binance.GetRecentTrades(symbol, limit)
	if err != nil {
		ctrl.Ctx.Resp(utils.ResJson(400, nil, err.Error()))
		return
	}
	items := make([]map[string]interface{}, 0, len(trades))
	for i := len(trades) - 1; i >= 0; i-- {
		t := trades[i]
		items = append(items, map[string]interface{}{
			"id":     t.ID,
			"price":  t.Price,
			"qty":    t.Quantity,
			"time":   t.Time,
			"maker":  t.IsBuyerMaker,
			"symbol": symbol,
		})
	}
	ctrl.Ctx.Resp(utils.ResJson(200, map[string]interface{}{
		"items": items,
	}))
}
