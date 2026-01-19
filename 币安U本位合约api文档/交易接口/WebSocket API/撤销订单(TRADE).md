[跳到主要内容](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order#__docusaurus_skipToContent_fallback)
[![Binance Logo](https://developers.binance.com/docs/zh-CN/img/logo.svg)](https://developers.binance.com/en)
产品▼
搜索`⌘``K`
当前全站
[](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
  * [English](https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
  * [简体中文](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)


衍生品交易
  * [更新日志](https://developers.binance.com/docs/zh-CN/derivatives/change-log)
  * [概述](https://developers.binance.com/docs/zh-CN/derivatives/Introduction)
  * [快速开始](https://developers.binance.com/docs/zh-CN/derivatives/quick-start)
  * [U本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
    * [基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/general-info)
    * [WebSocket API基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [通用枚举定义](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/common-definition)
    * [行情接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
    * [交易接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
      * [REST API](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
      * [WebSocket API](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
        * [下单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api)
        * [修改订单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Modify-Order)
        * [撤销订单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
        * [查询订单(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
        * [用户持仓风险V2(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
        * [用户持仓风险(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Information)
        * [条件单下单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/New-Algo-Order)
        * [条件单撤销订单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Algo-Order)
    * [Websocket行情推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
    * [Websocket账户信息推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
    * [账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
    * [合约钱包闪兑](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
    * [经典统一账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
    * [常见问题](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
    * [错误代码](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/error-code)
  * [币本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
  * [统一账户](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
  * [统一账户专业版](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
  * [期权](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)


  * [](https://developers.binance.com/docs/zh-CN/)
  * U本位合约
  * 交易接口
  * WebSocket API
  * 撤销订单(TRADE)


本页总览
# 撤销订单 (TRADE)
## 接口描述[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order#%E6%8E%A5%E5%8F%A3%E6%8F%8F%E8%BF%B0 "接口描述的直接链接")
撤销订单
## 方式[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order#%E6%96%B9%E5%BC%8F "方式的直接链接")
`order.cancel`
## 请求[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order#%E8%AF%B7%E6%B1%82 "请求的直接链接")
```
{  
   	"id": "5633b6a2-90a9-4192-83e7-925c90b6a2fd",  
    "method": "order.cancel",   
    "params": {   
      "apiKey": "HsOehcfih8ZRxnhjp2XjGXhsOBd6msAhKz9joQaWwZ7arcJTlD2hGOGQj1lGdTjR",   
      "orderId": 283194212,  
      "symbol": "BTCUSDT",   
      "timestamp": 1703439070722,   
      "signature": "b09c49815b4e3f1f6098cd9fbe26a933a9af79803deaaaae03c29f719c08a8a8"   
    }  
}  

```

## 请求权重[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D "请求权重的直接链接")
**1**
## 请求参数[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0 "请求参数的直接链接")
名称 | 类型 | 是否必需 | 描述  
---|---|---|---  
symbol | STRING | YES | 交易对  
orderId | LONG | NO | 系统订单号  
origClientOrderId | STRING | NO | 用户自定义的订单号  
recvWindow | LONG | NO |   
timestamp | LONG | YES |   
>   * `orderId` 与 `origClientOrderId` 必须至少发送一个
> 

## 响应示例[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B "响应示例的直接链接")
```
{  
  "id": "5633b6a2-90a9-4192-83e7-925c90b6a2fd",  
  "status": 200,  
  "result": {  
    "clientOrderId": "myOrder1",  
    "cumQty": "0",  
    "cumQuote": "0",  
    "executedQty": "0",  
    "orderId": 283194212,  
    "origQty": "11",  
    "origType": "TRAILING_STOP_MARKET",  
    "price": "0",  
    "reduceOnly": false,  
    "side": "BUY",  
    "positionSide": "SHORT",  
    "status": "CANCELED",  
    "stopPrice": "9300",                  
    "closePosition": false,    
    "symbol": "BTCUSDT",  
    "timeInForce": "GTC",  
    "type": "TRAILING_STOP_MARKET",  
    "activatePrice": "9020",              
    "priceRate": "0.3",                  
    "updateTime": 1571110484038,  
    "workingType": "CONTRACT_PRICE",  
    "priceProtect": false,             
    "priceMatch": "NONE",                
    "selfTradePreventionMode": "NONE",  
    "goodTillDate": 0                   
  },  
  "rateLimits": [  
    {  
      "rateLimitType": "REQUEST_WEIGHT",  
      "interval": "MINUTE",  
      "intervalNum": 1,  
      "limit": 2400,  
      "count": 1  
    }  
  ]  
}  

```

[上一页 修改订单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Modify-Order)[下一页 查询订单(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
  * [接口描述](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order#%E6%8E%A5%E5%8F%A3%E6%8F%8F%E8%BF%B0)
  * [方式](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order#%E6%96%B9%E5%BC%8F)
  * [请求](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order#%E8%AF%B7%E6%B1%82)
  * [请求权重](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D)
  * [请求参数](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0)
  * [响应示例](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B)


Copyright © 2026 Binance.
