[跳到主要内容](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance#__docusaurus_skipToContent_fallback)
[![Binance Logo](https://developers.binance.com/docs/zh-CN/img/logo.svg)](https://developers.binance.com/en)
产品▼
搜索`⌘``K`
当前全站
[](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
  * [English](https://developers.binance.com/docs/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
  * [简体中文](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)


衍生品交易
  * [更新日志](https://developers.binance.com/docs/zh-CN/derivatives/change-log)
  * [概述](https://developers.binance.com/docs/zh-CN/derivatives/Introduction)
  * [快速开始](https://developers.binance.com/docs/zh-CN/derivatives/quick-start)
  * [U本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
    * [基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/general-info)
    * [WebSocket API基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [通用枚举定义](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/common-definition)
    * [行情接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
    * [交易接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
    * [Websocket行情推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
    * [Websocket账户信息推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
    * [账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
      * [REST API](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
      * [WebSocket API](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
        * [账户余额(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api)
        * [账户余额(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
        * [账户信息V2(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
        * [账户信息(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information)
    * [合约钱包闪兑](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
    * [经典统一账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
    * [常见问题](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
    * [错误代码](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/error-code)
  * [币本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
  * [统一账户](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
  * [统一账户专业版](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
  * [期权](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)


  * [](https://developers.binance.com/docs/zh-CN/)
  * U本位合约
  * 账户接口
  * WebSocket API
  * 账户余额(USER-DATA)


本页总览
# 账户余额 (USER_DATA)
查询账户余额
## 方式[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance#%E6%96%B9%E5%BC%8F "方式的直接链接")
`account.balance`
## 请求[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance#%E8%AF%B7%E6%B1%82 "请求的直接链接")
```
{  
    "id": "605a6d20-6588-4cb9-afa0-b0ab087507ba",  
    "method": "account.balance",  
    "params": {  
        "apiKey": "xTaDyrmvA9XT2oBHHjy39zyPzKCvMdtH3b9q4xadkAg2dNSJXQGCxzui26L823W2",  
        "timestamp": 1702561978458,  
        "signature": "208bb94a26f99aa122b1319490ca9cb2798fccc81d9b6449521a26268d53217a"  
    }  
}  

```

## 请求权重[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D "请求权重的直接链接")
**5**
## 请求参数[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0 "请求参数的直接链接")
名称 | 类型 | 是否必需 | 描述  
---|---|---|---  
recvWindow | LONG | NO |   
timestamp | LONG | YES |   
## 响应示例[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B "响应示例的直接链接")
```
{  
    "id": "605a6d20-6588-4cb9-afa0-b0ab087507ba",  
    "status": 200,  
    "result": [  
      {  
        "accountAlias": "SgsR",    // 账户唯一识别码  
        "asset": "USDT",		// 资产  
        "balance": "122607.35137903",	// 总余额  
        "crossWalletBalance": "23.72469206", // 全仓余额  
          "crossUnPnl": "0.00000000"  // 全仓持仓未实现盈亏  
          "availableBalance": "23.72469206",       // 下单可用余额  
          "maxWithdrawAmount": "23.72469206",     // 最大可转出余额  
          "marginAvailable": true,    // 是否可用作联合保证金  
          "updateTime": 1617939110373  
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

[上一页 账户余额(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api)[下一页 账户信息V2(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
  * [方式](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance#%E6%96%B9%E5%BC%8F)
  * [请求](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance#%E8%AF%B7%E6%B1%82)
  * [请求权重](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D)
  * [请求参数](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0)
  * [响应示例](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B)


Copyright © 2026 Binance.
