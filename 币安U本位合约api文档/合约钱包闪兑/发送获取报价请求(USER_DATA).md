[跳到主要内容](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request#__docusaurus_skipToContent_fallback)
[![Binance Logo](https://developers.binance.com/docs/zh-CN/img/logo.svg)](https://developers.binance.com/en)
产品▼
搜索`⌘``K`
当前全站
[](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)
  * [English](https://developers.binance.com/docs/derivatives/usds-margined-futures/convert/Send-quote-request)
  * [简体中文](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)


衍生品交易
  * [更新日志](https://developers.binance.com/docs/zh-CN/derivatives/change-log)
  * [概述](https://developers.binance.com/docs/zh-CN/derivatives/Introduction)
  * [快速开始](https://developers.binance.com/docs/zh-CN/derivatives/quick-start)
  * [U本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)
    * [基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/general-info)
    * [WebSocket API基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [通用枚举定义](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/common-definition)
    * [行情接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)
    * [交易接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)
    * [Websocket行情推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)
    * [Websocket账户信息推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)
    * [账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)
    * [合约钱包闪兑](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)
      * [查询可交易币对信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert)
      * [发送获取报价请求(USER_DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)
      * [接受报价(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
      * [查询订单状态(USER_DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)
    * [经典统一账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)
    * [常见问题](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)
    * [错误代码](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/error-code)
  * [币本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)
  * [统一账户](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)
  * [统一账户专业版](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)
  * [期权](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request)


  * [](https://developers.binance.com/docs/zh-CN/)
  * U本位合约
  * 合约钱包闪兑
  * 发送获取报价请求(USER_DATA)


本页总览
# 发送获取报价请求(USER_DATA)
## 接口描述[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request#%E6%8E%A5%E5%8F%A3%E6%8F%8F%E8%BF%B0 "接口描述的直接链接")
对所需的币对发送获取报价请求
## HTTP请求[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request#http%E8%AF%B7%E6%B1%82 "HTTP请求的直接链接")
POST `/fapi/v1/convert/getQuote`
## 请求权重[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D "请求权重的直接链接")
**50(IP)** **每小时360次，每天500次**
## 请求参数[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0 "请求参数的直接链接")
名称 | 类型 | 是否必需 | 描述  
---|---|---|---  
fromAsset | STRING | YES |   
toAsset | STRING | YES |   
fromAmount | DECIMAL | EITHER | 这是成交后将被扣除的金额  
toAmount | DECIMAL | EITHER | 这是成交后将会获得的金额  
validTime | ENUM | NO | 可以支持10s，默认值为10s  
recvWindow | LONG | NO | 此值不能大于 60000  
timestamp | LONG | YES |   
>   * 参数fromAmount或者toAmount只需要提供其中一个。
>   * `quoteId`仅在账户余额充足时返回。
> 

## 响应示例[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B "响应示例的直接链接")
```
{  
   "quoteId":"12415572564",  
   "ratio":"38163.7",  
   "inverseRatio":"0.0000262",  
   "validTimestamp":1623319461670,  
   "toAmount":"3816.37",  
   "fromAmount":"0.1"  
}  

```

[上一页 查询可交易币对信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert)[下一页 接受报价(TRADE)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Accept-Quote)
  * [接口描述](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request#%E6%8E%A5%E5%8F%A3%E6%8F%8F%E8%BF%B0)
  * [HTTP请求](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request#http%E8%AF%B7%E6%B1%82)
  * [请求权重](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D)
  * [请求参数](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0)
  * [响应示例](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Send-quote-request#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B)


Copyright © 2026 Binance.
