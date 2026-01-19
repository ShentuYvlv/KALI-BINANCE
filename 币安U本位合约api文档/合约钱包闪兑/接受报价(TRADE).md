[跳到主要内容](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote#__docusaurus_skipToContent_fallback)
[![Binance Logo](https://developers.binance.com/docs/zh-CN/img/logo.svg)](https://developers.binance.com/en)
产品▼
搜索`⌘``K`
当前全站
[](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
  * [English](https://developers.binance.com/docs/derivatives/usds-margined-futures/convert/Accept-Quote)
  * [简体中文](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)


衍生品交易
  * [更新日志](https://developers.binance.com/docs/zh-CN/derivatives/change-log)
  * [概述](https://developers.binance.com/docs/zh-CN/derivatives/Introduction)
  * [快速开始](https://developers.binance.com/docs/zh-CN/derivatives/quick-start)
  * [U本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
    * [基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/general-info)
    * [WebSocket API基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [通用枚举定义](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/common-definition)
    * [行情接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
    * [交易接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
    * [Websocket行情推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
    * [Websocket账户信息推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
    * [账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
    * [合约钱包闪兑](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
      * [查询可交易币对信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert)
      * [发送获取报价请求(USER_DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)
      * [接受报价(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
      * [查询订单状态(USER_DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
    * [经典统一账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
    * [常见问题](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
    * [错误代码](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/error-code)
  * [币本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
  * [统一账户](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
  * [统一账户专业版](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
  * [期权](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)


  * [](https://developers.binance.com/docs/zh-CN/)
  * U本位合约
  * 合约钱包闪兑
  * 接受报价(TRADE)


本页总览
# 接受报价(TRADE)
## 接口描述[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote#%E6%8E%A5%E5%8F%A3%E6%8F%8F%E8%BF%B0 "接口描述的直接链接")
通过 quote ID 来接受报价。
## HTTP请求[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote#http%E8%AF%B7%E6%B1%82 "HTTP请求的直接链接")
POST `/fapi/v1/convert/acceptQuote`
## 请求权重[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D "请求权重的直接链接")
**200(IP)**
## 请求参数[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0 "请求参数的直接链接")
名称 | 类型 | 是否必需 | 描述  
---|---|---|---  
quoteId | STRING | YES |   
recvWindow | LONG | NO | 此值不能大于 60000  
timestamp | LONG | YES |   
## 响应示例[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B "响应示例的直接链接")
```
{  
  "orderId":"933256278426274426",  
  "createTime":1623381330472,  
  "orderStatus":"PROCESS" //PROCESS/ACCEPT_SUCCESS/SUCCESS/FAIL  
}  

```

[上一页 发送获取报价请求(USER_DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)[下一页 查询订单状态(USER_DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
  * [接口描述](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote#%E6%8E%A5%E5%8F%A3%E6%8F%8F%E8%BF%B0)
  * [HTTP请求](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote#http%E8%AF%B7%E6%B1%82)
  * [请求权重](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D)
  * [请求参数](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0)
  * [响应示例](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B)


Copyright © 2026 Binance.
