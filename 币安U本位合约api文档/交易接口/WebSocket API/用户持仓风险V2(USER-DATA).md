[跳到主要内容](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2#__docusaurus_skipToContent_fallback)
[![Binance Logo](https://developers.binance.com/docs/zh-CN/img/logo.svg)](https://developers.binance.com/en)
产品▼
搜索`⌘``K`
当前全站
[](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
  * [English](https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
  * [简体中文](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)


衍生品交易
  * [更新日志](https://developers.binance.com/docs/zh-CN/derivatives/change-log)
  * [概述](https://developers.binance.com/docs/zh-CN/derivatives/Introduction)
  * [快速开始](https://developers.binance.com/docs/zh-CN/derivatives/quick-start)
  * [U本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
    * [基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/general-info)
    * [WebSocket API基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [通用枚举定义](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/common-definition)
    * [行情接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
    * [交易接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
      * [REST API](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
      * [WebSocket API](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
        * [下单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api)
        * [修改订单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Modify-Order)
        * [撤销订单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order)
        * [查询订单(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)
        * [用户持仓风险V2(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
        * [用户持仓风险(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Information)
        * [条件单下单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/New-Algo-Order)
        * [条件单撤销订单(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Algo-Order)
    * [Websocket行情推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
    * [Websocket账户信息推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
    * [账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
    * [合约钱包闪兑](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
    * [经典统一账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
    * [常见问题](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
    * [错误代码](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/error-code)
  * [币本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
  * [统一账户](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
  * [统一账户专业版](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)
  * [期权](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2)


  * [](https://developers.binance.com/docs/zh-CN/)
  * U本位合约
  * 交易接口
  * WebSocket API
  * 用户持仓风险V2(USER-DATA)


本页总览
# 用户持仓风险V2 (USER_DATA)
## 接口描述[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2#%E6%8E%A5%E5%8F%A3%E6%8F%8F%E8%BF%B0 "接口描述的直接链接")
查询持仓风险，仅返回有仓位或挂单的交易对
## 方式[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2#%E6%96%B9%E5%BC%8F "方式的直接链接")
`v2/account.position`
## 请求[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2#%E8%AF%B7%E6%B1%82 "请求的直接链接")
```
{  
   	"id": "605a6d20-6588-4cb9-afa0-b0ab087507ba",  
    "method": "v2/account.position",  
    "params": {  
        "apiKey": "xTaDyrmvA9XT2oBHHjy39zyPzKCvMdtH3b9q4xadkAg2dNSJXQGCxzui26L823W2",  
        "symbol": "BTCUSDT",  
        "timestamp": 1702920680303,  
        "signature": "31ab02a51a3989b66c29d40fcdf78216978a60afc6d8dc1c753ae49fa3164a2a"  
    }  
}  

```

## 请求权重[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D "请求权重的直接链接")
**5**
## 请求参数[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0 "请求参数的直接链接")
名称 | 类型 | 是否必需 | 描述  
---|---|---|---  
symbol | STRING | NO |   
recvWindow | LONG | NO |   
timestamp | LONG | YES |   
**注意**
>   * 请与账户推送信息`ACCOUNT_UPDATE`配合使用，以满足您的及时性和准确性需求。
> 

## 响应示例[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B "响应示例的直接链接")
> 单向持仓模式下：
```
{  
  "id": "605a6d20-6588-4cb9-afa0-b0ab087507ba",  
  "status": 200,  
  "result":[  
    {  
        "symbol": "ADAUSDT",  
        "positionSide": "BOTH",               // 持仓方向  
        "positionAmt": "30",  
        "entryPrice": "0.385",  
        "breakEvenPrice": "0.385077",  
        "markPrice": "0.41047590",  
        "unRealizedProfit": "0.76427700",     // 持仓未实现盈亏   
        "liquidationPrice": "0",  
        "isolatedMargin": "0",  
        "notional": "12.31427700",  
        "marginAsset": "USDT",  
        "isolatedWallet": "0",  
        "initialMargin": "0.61571385",        // 初始保证金  
        "maintMargin": "0.08004280",          // 维持保证金  
        "positionInitialMargin": "0.61571385",// 仓位初始保证金  
        "openOrderInitialMargin": "0",        // 订单初始保证金  
        "adl": 2,  
        "bidNotional": "0",                     
        "askNotional": "0",                     
        "updateTime": 1720736417660           // 更新时间  
    }  
  ],  
  "rateLimits": [  
    {  
      "rateLimitType": "REQUEST_WEIGHT",  
      "interval": "MINUTE",  
      "intervalNum": 1,  
      "limit": 2400,  
      "count": 20  
    }  
  ]  
}  

```

> 双向持仓模式下：
```
{  
  "id": "605a6d20-6588-4cb9-afa0-b0ab087507ba",  
  "status": 200,  
  "result": [  
  {  
        "symbol": "ADAUSDT",  
        "positionSide": "LONG",                 
        "positionAmt": "30",  
        "entryPrice": "0.385",  
        "breakEvenPrice": "0.385077",  
        "markPrice": "0.41047590",  
        "unRealizedProfit": "0.76427700",      
        "liquidationPrice": "0",  
        "isolatedMargin": "0",  
        "notional": "12.31427700",  
        "marginAsset": "USDT",  
        "isolatedWallet": "0",  
        "initialMargin": "0.61571385",          
        "maintMargin": "0.08004280",            
        "positionInitialMargin": "0.61571385",  
        "openOrderInitialMargin": "0",         
        "adl": 2,  
        "bidNotional": "0",                    
        "askNotional": "0",                  
        "updateTime": 1720736417660  
  },  
  {  
        "symbol": "COMPUSDT",  
        "positionSide": "SHORT",  
        "positionAmt": "-1.000",  
        "entryPrice": "70.92841",  
        "breakEvenPrice": "70.900038636",  
        "markPrice": "49.72023376",  
        "unRealizedProfit": "21.20817624",  
        "liquidationPrice": "2260.56757210",  
        "isolatedMargin": "0",  
        "notional": "-49.72023376",  
        "marginAsset": "USDT",  
        "isolatedWallet": "0",  
        "initialMargin": "2.48601168",  
        "maintMargin": "0.49720233",  
        "positionInitialMargin": "2.48601168",  
        "openOrderInitialMargin": "0",  
        "adl": 2,  
        "bidNotional": "0",  
        "askNotional": "0",  
        "updateTime": 1708943511656  
  }  
  ],  
  "rateLimits": [  
    {  
      "rateLimitType": "REQUEST_WEIGHT",  
      "interval": "MINUTE",  
      "intervalNum": 1,  
      "limit": 2400,  
      "count": 20  
    }  
  ]  
}  

```

[上一页 查询订单(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Query-Order)[下一页 用户持仓风险(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Information)
  * [接口描述](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2#%E6%8E%A5%E5%8F%A3%E6%8F%8F%E8%BF%B0)
  * [方式](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2#%E6%96%B9%E5%BC%8F)
  * [请求](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2#%E8%AF%B7%E6%B1%82)
  * [请求权重](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D)
  * [请求参数](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0)
  * [响应示例](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B)


Copyright © 2026 Binance.
