[跳到主要内容](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order#__docusaurus_skipToContent_fallback)
[![Binance Logo](https://developers.binance.com/docs/zh-CN/img/logo.svg)](https://developers.binance.com/en)
产品▼
搜索`⌘``K`
当前全站
[](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
  * [English](https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
  * [简体中文](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)


衍生品交易
  * [更新日志](https://developers.binance.com/docs/zh-CN/derivatives/change-log)
  * [概述](https://developers.binance.com/docs/zh-CN/derivatives/Introduction)
  * [快速开始](https://developers.binance.com/docs/zh-CN/derivatives/quick-start)
  * [U本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
    * [基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/general-info)
    * [WebSocket API基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [通用枚举定义](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/common-definition)
    * [行情接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
    * [交易接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
      * [REST API](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
      * [WebSocket API](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
        * [下单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api)
        * [修改订单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Modify-Order)
        * [撤销订单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
        * [查询订单(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
        * [用户持仓风险V2(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
        * [用户持仓风险(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Information)
        * [条件单下单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/New-Algo-Order)
        * [条件单撤销订单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Algo-Order)
    * [Websocket行情推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
    * [Websocket账户信息推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
    * [账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
    * [合约钱包闪兑](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
    * [经典统一账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
    * [常见问题](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
    * [错误代码](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/error-code)
  * [币本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
  * [统一账户](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
  * [统一账户专业版](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
  * [期权](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)


  * [](https://developers.binance.com/docs/zh-CN/)
  * U本位合约
  * 交易接口
  * WebSocket API
  * 查询订单(USER-DATA)


本页总览
# 查询订单 (USER_DATA)
## 接口描述[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order#%E6%8E%A5%E5%8F%A3%E6%8F%8F%E8%BF%B0 "接口描述的直接链接")
查询订单状态
  * 请注意，如果订单满足如下条件，不会被查询到： 
    * 订单的最终状态为 `CANCELED` 或者 `EXPIRED` **并且** 订单没有任何的成交记录 **并且** 订单生成时间 + 3天 < 当前时间
    * 订单创建时间 + 90天 < 当前时间


## 方式[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order#%E6%96%B9%E5%BC%8F "方式的直接链接")
`order.status`
## 请求[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order#%E8%AF%B7%E6%B1%82 "请求的直接链接")
```
{  
    "id": "0ce5d070-a5e5-4ff2-b57f-1556741a4204",  
    "method": "order.status",  
    "params": {  
        "apiKey": "HMOchcfii9ZRZnhjp2XjGXhsOBd6msAhKz9joQaWwZ7arcJTlD2hGPHQj1lGdTjR",  
        "orderId": 328999071,  
        "symbol": "BTCUSDT",  
        "timestamp": 1703441060152,  
        "signature": "ba48184fc38a71d03d2b5435bd67c1206e3191e989fe99bda1bc643a880dfdbf"  
    }  
}  

```

## 请求权重[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D "请求权重的直接链接")
**1**
## 请求参数[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0 "请求参数的直接链接")
名称 | 类型 | 是否必需 | 描述  
---|---|---|---  
symbol | STRING | YES | 交易对  
orderId | LONG | NO | 系统订单号  
origClientOrderId | STRING | NO | 用户自定义的订单号  
recvWindow | LONG | NO |   
timestamp | LONG | YES |   
注意:
>   * 至少需要发送 `orderId` 与 `origClientOrderId`中的一个
>   * `orderId`在`symbol`维度是自增的
> 

## 响应示例[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B "响应示例的直接链接")
```
{  
  	"avgPrice": "0.00000",				// 平均成交价  
  	"clientOrderId": "abc",				// 用户自定义的订单号  
  	"cumQuote": "0",					// 成交金额  
  	"executedQty": "0",					// 成交量  
  	"orderId": 1573346959,				// 系统订单号  
  	"origQty": "0.40",					// 原始委托数量  
  	"origType": "TRAILING_STOP_MARKET",	// 触发前订单类型  
  	"price": "0",						// 委托价格  
  	"reduceOnly": false,				// 是否仅减仓  
  	"side": "BUY",						// 买卖方向  
  	"positionSide": "SHORT", 			// 持仓方向  
  	"status": "NEW",					// 订单状态  
  	"stopPrice": "9300",			    // 触发价，对`TRAILING_STOP_MARKET`无效  
  	"closePosition": false,             // 是否条件全平仓  
  	"symbol": "BTCUSDT",				// 交易对  
  	"time": 1579276756075,				// 订单时间  
  	"timeInForce": "GTC",				// 有效方法  
  	"type": "TRAILING_STOP_MARKET",		// 订单类型  
  	"activatePrice": "9020",			// 跟踪止损激活价格, 仅`TRAILING_STOP_MARKET` 订单返回此字段  
  	"priceRate": "0.3",					// 跟踪止损回调比例, 仅`TRAILING_STOP_MARKET` 订单返回此字段  
  	"updateTime": 1579276756075,		// 更新时间  
  	"workingType": "CONTRACT_PRICE",    // 条件价格触发类型  
 	"priceProtect": false               // 是否开启条件单触发保护  
}  

```

[上一页 撤销订单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)[下一页 用户持仓风险V2(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
  * [接口描述](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order#%E6%8E%A5%E5%8F%A3%E6%8F%8F%E8%BF%B0)
  * [方式](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order#%E6%96%B9%E5%BC%8F)
  * [请求](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order#%E8%AF%B7%E6%B1%82)
  * [请求权重](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D)
  * [请求参数](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0)
  * [响应示例](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B)


Copyright © 2026 Binance.
