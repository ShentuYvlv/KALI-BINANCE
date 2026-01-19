[跳到主要内容](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2#__docusaurus_skipToContent_fallback)
[![Binance Logo](https://developers.binance.com/docs/zh-CN/img/logo.svg)](https://developers.binance.com/en)
产品▼
搜索`⌘``K`
当前全站
[](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
  * [English](https://developers.binance.com/docs/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
  * [简体中文](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)


衍生品交易
  * [更新日志](https://developers.binance.com/docs/zh-CN/derivatives/change-log)
  * [概述](https://developers.binance.com/docs/zh-CN/derivatives/Introduction)
  * [快速开始](https://developers.binance.com/docs/zh-CN/derivatives/quick-start)
  * [U本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
    * [基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/general-info)
    * [WebSocket API基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [通用枚举定义](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/common-definition)
    * [行情接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
    * [交易接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
    * [Websocket行情推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
    * [Websocket账户信息推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
    * [账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
      * [REST API](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
      * [WebSocket API](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
        * [账户余额(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api)
        * [账户余额(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)
        * [账户信息V2(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
        * [账户信息(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information)
    * [合约钱包闪兑](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
    * [经典统一账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
    * [常见问题](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
    * [错误代码](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/error-code)
  * [币本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
  * [统一账户](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
  * [统一账户专业版](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)
  * [期权](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2)


  * [](https://developers.binance.com/docs/zh-CN/)
  * U本位合约
  * 账户接口
  * WebSocket API
  * 账户信息V2(USER-DATA)


本页总览
# 账户信息V2 (USER_DATA)
## 接口描述[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2#%E6%8E%A5%E5%8F%A3%E6%8F%8F%E8%BF%B0 "接口描述的直接链接")
现有账户信息。 用户在单资产模式和多资产模式下会看到不同结果，响应部分的注释解释了两种模式下的不同。
## 方式[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2#%E6%96%B9%E5%BC%8F "方式的直接链接")
`v2/account.status`
## 请求[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2#%E8%AF%B7%E6%B1%82 "请求的直接链接")
```
{  
    "id": "605a6d20-6588-4cb9-afa0-b0ab087507ba",  
    "method": "v2/account.status",  
    "params": {  
        "apiKey": "xTaDyrmvA9XT2oBHHjy39zyPzKCvMdtH3b9q4xadkAg2dNSJXQGCxzui26L823W2",  
        "timestamp": 1702620814781,  
        "signature": "6bb98ef84170c70ba3d01f44261bfdf50fef374e551e590de22b5c3b729b1d8c"  
    }  
}  

```

## 请求权重[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D "请求权重的直接链接")
**5**
## 请求参数[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0 "请求参数的直接链接")
名称 | 类型 | 是否必需 | 描述  
---|---|---|---  
recvWindow | LONG | NO |   
timestamp | LONG | YES |   
## 响应示例[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B "响应示例的直接链接")
> 单资产模式
```
{  
  "id": "605a6d20-6588-4cb9-afa0-b0ab087507ba",  
  "status": 200,  
  "result": {     
	"totalInitialMargin": "0.00000000",            // 当前所需起始保证金总额(存在逐仓请忽略), 仅计算usdt资产positions), only for USDT asset  
	"totalMaintMargin": "0.00000000",  	           // 维持保证金总额, 仅计算usdt资产  
	"totalWalletBalance": "103.12345678",          // 账户总余额, 仅计算usdt资产  
	"totalUnrealizedProfit": "0.00000000",         // 持仓未实现盈亏总额, 仅计算usdt资产  
	"totalMarginBalance": "103.12345678",          // 保证金总余额, 仅计算usdt资产  
	"totalPositionInitialMargin": "0.00000000",    // 持仓所需起始保证金(基于最新标记价格), 仅计算usdt资产  
	"totalOpenOrderInitialMargin": "0.00000000",   // 当前挂单所需起始保证金(基于最新标记价格), 仅计算usdt资产  
	"totalCrossWalletBalance": "103.12345678",     // 全仓账户余额, 仅计算usdt资产  
	"totalCrossUnPnl": "0.00000000",	           // 全仓持仓未实现盈亏总额, 仅计算usdt资产  
	"availableBalance": "103.12345678",            // 可用余额, 仅计算usdt资产  
	"maxWithdrawAmount": "103.12345678"            // 最大可转出余额, 仅计算usdt资产  
	"assets": [   
		{  
			"asset": "USDT",			            // 资产  
			"walletBalance": "23.72469206",         // 余额  
			"unrealizedProfit": "0.00000000",       // 未实现盈亏  
			"marginBalance": "23.72469206",         // 保证金余额  
			"maintMargin": "0.00000000",	        // 维持保证金  
			"initialMargin": "0.00000000",          // 当前所需起始保证金  
			"positionInitialMargin": "0.00000000",  // 持仓所需起始保证金(基于最新标记价格)  
			"openOrderInitialMargin": "0.00000000", // 当前挂单所需起始保证金(基于最新标记价格)  
			"crossWalletBalance": "23.72469206",    // 全仓账户余额  
			"crossUnPnl": "0.00000000"              // 全仓持仓未实现盈亏  
			"availableBalance": "23.72469206",      // 可用余额  
			"maxWithdrawAmount": "23.72469206",     // 最大可转出余额  
			"updateTime": 1625474304765             // 更新时间  
		},     
		{  
			"asset": "USDC",			            
			"walletBalance": "103.12345678",        
			"unrealizedProfit": "0.00000000",       
			"marginBalance": "103.12345678",        
			"maintMargin": "0.00000000",	       
			"initialMargin": "0.00000000",     
			"positionInitialMargin": "0.00000000",     
			"openOrderInitialMargin": "0.00000000",    
			"crossWalletBalance": "103.12345678",      
			"crossUnPnl": "0.00000000"        
			"availableBalance": "126.72469206",   
			"maxWithdrawAmount": "103.12345678",   
			"updateTime": 1625474304765   
		}  
    ],  
	"positions": [  // 仅有仓位或挂单的交易对会被返回  
		            // 根据用户持仓模式展示持仓方向，即单向模式下只返回BOTH持仓情况，双向模式下只返回 LONG 和 SHORT 持仓情况  
   	  {  
           "symbol": "BTCUSDT",               // 交易对  
           "positionSide": "BOTH",            // 持仓方向  
           "positionAmt": "1.000",            // 持仓数量  
           "unrealizedProfit": "0.00000000",  // 持仓未实现盈亏     
           "isolatedMargin": "0.00000000",	  
           "notional": "0",  
           "isolatedWallet": "0",  
           "initialMargin": "0",              // 持仓所需起始保证金(基于最新标记价格)  
           "maintMargin": "0",                // 当前杠杆下用户可用的最大名义价值  
           "updateTime": 0                    // 更新时间   
  	  }   
	]  
  },  
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

> 多资产模式
```
{  
  "id": "605a6d20-6588-4cb9-afa0-b0ab087507ba",  
  "status": 200,  
  "result": {     
	"totalInitialMargin": "0.00000000",          // 以USD计价的所需起始保证金总额  
	"totalMaintMargin": "0.00000000",  	         // 以USD计价的维持保证金总额  
	"totalWalletBalance": "126.72469206",        // 以USD计价的账户总余额  
	"totalUnrealizedProfit": "0.00000000",       // 以USD计价的持仓未实现盈亏总额  
	"totalMarginBalance": "126.72469206",        // 以USD计价的保证金总余额   
	"totalPositionInitialMargin": "0.00000000",  // 以USD计价的持仓所需起始保证金(基于最新标记价格)  
	"totalOpenOrderInitialMargin": "0.00000000", // 以USD计价的当前挂单所需起始保证金(基于最新标记价格)  
	"totalCrossWalletBalance": "126.72469206",   // 以USD计价的全仓账户余额  
	"totalCrossUnPnl": "0.00000000",	         // 以USD计价的全仓持仓未实现盈亏总额  
	"availableBalance": "126.72469206",          // 以USD计价的可用余额  
	"maxWithdrawAmount": "126.72469206"          // 以USD计价的最大可转出余额  
	"assets": [  
		{  
			"asset": "USDT",			            // 资产  
			"walletBalance": "23.72469206",         // 余额  
			"unrealizedProfit": "0.00000000",       // 未实现盈亏  
			"marginBalance": "23.72469206",         // 保证金余额  
			"maintMargin": "0.00000000",	        // 维持保证金  
			"initialMargin": "0.00000000",          // 当前所需起始保证金  
			"positionInitialMargin": "0.00000000",  // 持仓所需起始保证金(基于最新标记价格)  
			"openOrderInitialMargin": "0.00000000", // 当前挂单所需起始保证金(基于最新标记价格)  
			"crossWalletBalance": "23.72469206",    // 全仓账户余额  
			"crossUnPnl": "0.00000000"              // 全仓持仓未实现盈亏  
			"availableBalance": "23.72469206",      // 可用余额  
			"maxWithdrawAmount": "23.72469206",     // 最大可转出余额  
			"updateTime": 1625474304765             // 更新时间  
		},     
		{  
			"asset": "USDC",			            
			"walletBalance": "103.12345678",        
			"unrealizedProfit": "0.00000000",       
			"marginBalance": "103.12345678",        
			"maintMargin": "0.00000000",	       
			"initialMargin": "0.00000000",     
			"positionInitialMargin": "0.00000000",     
			"openOrderInitialMargin": "0.00000000",    
			"crossWalletBalance": "103.12345678",      
			"crossUnPnl": "0.00000000"        
			"availableBalance": "126.72469206",   
			"maxWithdrawAmount": "103.12345678",   
			"updateTime": 1625474304765   
		}  
	],  
	"positions": [  // 仅有仓位或挂单的交易对会被返回  
		            // 根据用户持仓模式展示持仓方向，即单向模式下只返回BOTH持仓情况，双向模式下只返回 LONG 和 SHORT 持仓情况  
   	  {  
           "symbol": "BTCUSDT",               // 交易对  
           "positionSide": "BOTH",            // 持仓方向  
           "positionAmt": "1.000",            // 持仓数量  
           "unrealizedProfit": "0.00000000",  // 持仓未实现盈亏     
           "isolatedMargin": "0.00000000",	  
           "notional": "0",  
           "isolatedWallet": "0",  
           "initialMargin": "0",              // 持仓所需起始保证金(基于最新标记价格)  
           "maintMargin": "0",                // 当前杠杆下用户可用的最大名义价值  
           "updateTime": 0                    // 更新时间   
  	  }   
	]  
  },  
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

[上一页 账户余额(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Futures-Account-Balance)[下一页 账户信息(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information)
  * [接口描述](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2#%E6%8E%A5%E5%8F%A3%E6%8F%8F%E8%BF%B0)
  * [方式](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2#%E6%96%B9%E5%BC%8F)
  * [请求](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2#%E8%AF%B7%E6%B1%82)
  * [请求权重](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D)
  * [请求参数](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0)
  * [响应示例](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/account/websocket-api/Account-Information-V2#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B)


Copyright © 2026 Binance.
