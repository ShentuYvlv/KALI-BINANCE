[跳到主要内容](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints#__docusaurus_skipToContent_fallback)
[![Binance Logo](https://developers.binance.com/docs/zh-CN/img/logo.svg)](https://developers.binance.com/en)
产品▼
搜索`⌘``K`
当前全站
[](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)
  * [English](https://developers.binance.com/docs/derivatives/usds-margined-futures/portfolio-margin-endpoints)
  * [简体中文](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)


衍生品交易
  * [更新日志](https://developers.binance.com/docs/zh-CN/derivatives/change-log)
  * [概述](https://developers.binance.com/docs/zh-CN/derivatives/Introduction)
  * [快速开始](https://developers.binance.com/docs/zh-CN/derivatives/quick-start)
  * [U本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)
    * [基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/general-info)
    * [WebSocket API基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [通用枚举定义](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/common-definition)
    * [行情接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)
    * [交易接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)
    * [Websocket行情推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)
    * [Websocket账户信息推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)
    * [账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)
    * [合约钱包闪兑](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)
    * [经典统一账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)
      * [查询经典统一账户账户信息(USER-DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)
    * [常见问题](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)
    * [错误代码](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/error-code)
  * [币本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)
  * [统一账户](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)
  * [统一账户专业版](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)
  * [期权](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints)


  * [](https://developers.binance.com/docs/zh-CN/)
  * U本位合约
  * 经典统一账户接口
  * 查询经典统一账户账户信息(USER-DATA)


本页总览
# Classic Portfolio Margin Account Information
## 查询经典统一账户账户信息 (USER_DATA)[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints#%E6%9F%A5%E8%AF%A2%E7%BB%8F%E5%85%B8%E7%BB%9F%E4%B8%80%E8%B4%A6%E6%88%B7%E8%B4%A6%E6%88%B7%E4%BF%A1%E6%81%AF-user_data "查询经典统一账户账户信息 \(USER_DATA\)的直接链接")
## 接口描述[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints#%E6%8E%A5%E5%8F%A3%E6%8F%8F%E8%BF%B0 "接口描述的直接链接")
查询经典统一账户当前账户信息
## HTTP请求[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints#http%E8%AF%B7%E6%B1%82 "HTTP请求的直接链接")
GET `/fapi/v1/pmAccountInfo`
## 请求权重[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D "请求权重的直接链接")
**5**
## 请求参数[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0 "请求参数的直接链接")
名称 | 类型 | 是否必需 | 描述  
---|---|---|---  
asset | STRING | YES |   
recvWindow | LONG | NO |   
timestamp | LONG | YES |   
>   * 最大可转出余额指可以转出到现货钱包到金额。
> 

## 响应示例[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B "响应示例的直接链接")
```
{  
    "maxWithdrawAmountUSD": "1627523.32459208",   //经典统一账户以USD计价的最大可转出余额  
    "asset": "BTC",            // 资产  
    "maxWithdrawAmount": "27.43689636",        //最大可转出余额  
}  

```

[上一页 查询订单状态(USER_DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/convert/Order-Status)[下一页 自我交易预防常见问题](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/faq/stp-faq)
  * [查询经典统一账户账户信息 (USER_DATA)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints#%E6%9F%A5%E8%AF%A2%E7%BB%8F%E5%85%B8%E7%BB%9F%E4%B8%80%E8%B4%A6%E6%88%B7%E8%B4%A6%E6%88%B7%E4%BF%A1%E6%81%AF-user_data)
  * [接口描述](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints#%E6%8E%A5%E5%8F%A3%E6%8F%8F%E8%BF%B0)
  * [HTTP请求](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints#http%E8%AF%B7%E6%B1%82)
  * [请求权重](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints#%E8%AF%B7%E6%B1%82%E6%9D%83%E9%87%8D)
  * [请求参数](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0)
  * [响应示例](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/portfolio-margin-endpoints#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B)


Copyright © 2026 Binance.
