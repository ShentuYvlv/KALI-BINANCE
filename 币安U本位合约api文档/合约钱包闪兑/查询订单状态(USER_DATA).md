[跳到主要内容](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status#__docusaurus_skipToContent_fallback)
[![Binance Logo](https://developers.binance.com/docs/zh-CN/img/logo.svg)](https://developers.binance.com/en)
产品▼
搜索`⌘``K`
当前全站
[](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
  * [English](https://developers.binance.com/docs/derivatives/usds-margined-futures/convert/Order-Status)
  * [简体中文](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)


衍生品交易
  * [更新日志](https://developers.binance.com/docs/zh-CN/derivatives/change-log)
  * [概述](https://developers.binance.com/docs/zh-CN/derivatives/Introduction)
  * [快速开始](https://developers.binance.com/docs/zh-CN/derivatives/quick-start)
  * [U本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
    * [基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/general-info)
    * [WebSocket API基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [通用枚举定义](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/common-definition)
    * [行情接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
    * [交易接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
    * [Websocket行情推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
    * [Websocket账户信息推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
    * [账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
    * [合约钱包闪兑](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
      * [查询可交易币对信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert)
      * [发送获取报价请求(USER_DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)
      * [接受报价(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
      * [查询订单状态(USER_DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
    * [经典统一账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
    * [常见问题](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
    * [错误代码](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/error-code)
  * [币本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
  * [统一账户](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
  * [统一账户专业版](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
  * [期权](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)


  * [](https://developers.binance.com/docs/zh-CN/)
  * U本位合约
  * 合约钱包闪兑
  * 查询订单状态(USER_DATA)


本页总览
# 查询订单状态(USER_DATA)
## 接口描述[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status#%E6%8E%A5%E5%8F%A3%E6%8F%8F%E8%BF%B0 "接口描述的直接链接")
通过 order ID 来查询订单状态。
## HTTP请求[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status#http%E8%AF%B7%E6%B1%82 "HTTP请求的直接链接")
GET `/fapi/v1/convert/orderStatus`
## 请求权重[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D "请求权重的直接链接")
**50(IP)**
## 请求参数[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0 "请求参数的直接链接")
名称 | 类型 | 是否必需 | 描述  
---|---|---|---  
orderId | STRING | NO | orderId 和quoteId需要填其中一个  
quoteId | STRING | NO | orderId 和quoteId需要填其中一个  
## 响应示例[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B "响应示例的直接链接")
```
{  
  "orderId":933256278426274426,  
  "orderStatus":"SUCCESS",  
  "fromAsset":"BTC",  
  "fromAmount":"0.00054414",  
  "toAsset":"USDT",  
  "toAmount":"20",  
  "ratio":"36755",  
  "inverseRatio":"0.00002721",  
  "createTime":1623381330472  
}  

```

[上一页 接受报价(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)[下一页 查询经典统一账户账户信息(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)
  * [接口描述](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status#%E6%8E%A5%E5%8F%A3%E6%8F%8F%E8%BF%B0)
  * [HTTP请求](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status#http%E8%AF%B7%E6%B1%82)
  * [请求权重](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D)
  * [请求参数](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0)
  * [响应示例](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B)


Copyright © 2026 Binance.
