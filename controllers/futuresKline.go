package controllers

import (
	"net/http"
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

type FuturesKlineController struct {
	web.Controller
}

func (ctrl *FuturesKlineController) Get() {
	symbol := ctrl.GetString("symbol")
	interval := ctrl.GetString("interval", "1m")
	limit, _ := ctrl.GetInt("limit", 500)
	startTime, _ := ctrl.GetInt64("startTime", 0)
	endTime, _ := ctrl.GetInt64("endTime", 0)

	if symbol == "" {
		ctrl.Ctx.Resp(utils.ResJson(400, nil, "symbol required"))
		return
	}
	if limit <= 0 {
		limit = 500
	}
	if limit > 1500 {
		limit = 1500
	}

	klines, err := binance.GetKlineDataRange(symbol, interval, limit, startTime, endTime)
	if err != nil {
		ctrl.Ctx.Resp(utils.ResJson(500, nil, err.Error()))
		return
	}

	data := make([]map[string]interface{}, 0, len(klines))
	for _, k := range klines {
		open, _ := strconv.ParseFloat(k.Open, 64)
		high, _ := strconv.ParseFloat(k.High, 64)
		low, _ := strconv.ParseFloat(k.Low, 64)
		closePrice, _ := strconv.ParseFloat(k.Close, 64)
		volume, _ := strconv.ParseFloat(k.Volume, 64)
		turnover, _ := strconv.ParseFloat(k.QuoteAssetVolume, 64)
		data = append(data, map[string]interface{}{
			"timestamp": k.OpenTime,
			"open":      open,
			"high":      high,
			"low":       low,
			"close":     closePrice,
			"volume":    volume,
			"turnover":  turnover,
		})
	}

	ctrl.Ctx.Resp(map[string]interface{}{
		"code": 200,
		"data": data,
		"msg":  "success",
	})
}

var klineWsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (ctrl *FuturesKlineController) Ws() {
	symbol := ctrl.GetString("symbol")
	interval := ctrl.GetString("interval", "1m")
	if symbol == "" {
		ctrl.CustomAbort(400, "symbol required")
		return
	}

	conn, err := klineWsUpgrader.Upgrade(ctrl.Ctx.ResponseWriter, ctrl.Ctx.Request, nil)
	if err != nil {
		logs.Error("kline ws upgrade error:", err)
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
			close(stopC)
		})
	}

	doneC, stopC, err = futures.WsKlineServe(symbol, interval, func(event *futures.WsKlineEvent) {
		k := event.Kline
		open, _ := strconv.ParseFloat(k.Open, 64)
		high, _ := strconv.ParseFloat(k.High, 64)
		low, _ := strconv.ParseFloat(k.Low, 64)
		closePrice, _ := strconv.ParseFloat(k.Close, 64)
		volume, _ := strconv.ParseFloat(k.Volume, 64)
		turnover, _ := strconv.ParseFloat(k.QuoteVolume, 64)

		msg := map[string]interface{}{
			"type": "kline",
			"data": map[string]interface{}{
				"timestamp": k.StartTime,
				"open":      open,
				"high":      high,
				"low":       low,
				"close":     closePrice,
				"volume":    volume,
				"turnover":  turnover,
				"final":     k.IsFinal,
			},
		}

		if atomic.LoadInt32(&closed) == 1 {
			return
		}
		writeMu.Lock()
		if err := conn.WriteJSON(msg); err != nil {
			atomic.StoreInt32(&closed, 1)
			closeStop(stopC)
		}
		writeMu.Unlock()
	}, func(err error) {
		logs.Error("kline ws error:", err)
	})
	if err != nil {
		logs.Error("kline ws start error:", err)
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
