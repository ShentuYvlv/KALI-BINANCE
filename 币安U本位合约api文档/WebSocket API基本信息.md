[跳到主要内容](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#__docusaurus_skipToContent_fallback)
[![Binance Logo](https://developers.binance.com/docs/zh-CN/img/logo.svg)](https://developers.binance.com/en)
产品▼
搜索`⌘``K`
当前全站
[](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
  * [English](https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-api-general-info)
  * [简体中文](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)


衍生品交易
  * [更新日志](https://developers.binance.com/docs/zh-CN/derivatives/change-log)
  * [概述](https://developers.binance.com/docs/zh-CN/derivatives/Introduction)
  * [快速开始](https://developers.binance.com/docs/zh-CN/derivatives/quick-start)
  * [U本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/general-info)
    * [WebSocket API基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [通用枚举定义](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/common-definition)
    * [行情接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [交易接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [Websocket行情推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [Websocket账户信息推送](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [合约钱包闪兑](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [经典统一账户接口](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [常见问题](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
    * [错误代码](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/error-code)
  * [币本位合约](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
  * [统一账户](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
  * [统一账户专业版](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)
  * [期权](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info)


  * [](https://developers.binance.com/docs/zh-CN/)
  * U本位合约
  * WebSocket API基本信息


本页总览
# WebSocket API基本信息
  * Base url为：**'wss://ws-fapi.binance.com/ws-fapi/v1'**
    * 测试网的Base url为：`wss://testnet.binancefuture.com/ws-fapi/v1`
  * 单次连接API有效期仅为24小时;预计在 24 小时标记后断开连接。
  * Websocket服务器每3分钟发送一个ping消息。 
    * 如果 websocket 服务器在10分钟内没有收到来自连接的`pong frame`，则连接将断开。
    * 当客户收到ping消息，必需尽快回复pong消息，同时payload需要和ping消息一致。
    * 未经请求的pong消息是被允许的，但是不会保证连接不断开。**对于这些pong消息，建议payload为空** 。
  * 必须通过获取除签名之外的所有请求参数并按字母顺序按名称排序来生成签名payload
  * 除非另有说明，否则列表按**时间顺序** 返回。
  * 除非另有说明，否则所有时间戳均以 UTC 中的毫秒为单位。
  * 除非另有说明，否则所有字段名称和值均区分大小写。
  * **`INT`参数（如时间戳）应为 JSON 整数，而不是字符串** 。
  * **`DECIMAL`参数（如 price）应为 JSON 字符串，而不是浮点数** 。
  * **用户数据流请求 - 您需要建立单独的WebSocket连接来获得 [账户信息推送]（[https://binance-docs.github.io/apidocs/futures/cn/#websocket-2）](https://binance-docs.github.io/apidocs/futures/cn/#websocket-2%EF%BC%89)**


> 请求示例:
```
{  
  "id": "9ca10e58-7452-467e-9454-f669bb9c764e",  
  "method": "order.place",  
  "params": {  
    "apiKey": "yeqKcXjtA9Eu4Tr3nJk61UJAGzXsEmFqqfVterxpMpR4peNfqE7Zl7oans8Qj089",  
    "price": "42088.0",  
    "quantity": "0.1",  
    "recvWindow": 5000,  
    "side": "BUY",  
    "signature": "996962a19802b5a09d7bc6ab1524227894533322a2f8a1f8934991689cabf8fe",  
    "symbol": "BTCUSDT",  
    "timeInForce": "GTC",  
    "timestamp": 1705311512994,  
    "type": "LIMIT"  
  }  
}  

```

请求字段:
名称 | 类型 | 是否必需 | 描述  
---|---|---|---  
`id` | INT / STRING / `null` | YES | 任意的 ID 用于匹配对请求的响应  
`method` | STRING | YES | 请求函数名称  
`params` | OBJECT | NO | 请求参数。如果没有参数可以省略  
  * 请求 `id` 是任意的。可以使用 UUID、顺次 ID、当前时间戳等。 服务器不会以任何方式解释 `id`，只是在响应中回显它。
可以在一个会话中自由重复使用 ID，不过请注意不要一次发送多个具有相同 ID 的请求，因为否则可能无法区分响应。
  * 请求函数名称可以以显式版本为前缀，例如：`"v3/order.place"`
  * `params` 的顺序不重要。


## WebSocket API响应格式[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#websocket-api%E5%93%8D%E5%BA%94%E6%A0%BC%E5%BC%8F "WebSocket API响应格式的直接链接")
响应在 **text 帧** 中以 JSON 格式返回，每帧一个响应。
> 成功响应示例:
```
{  
  "id": "43a3843a-2321-4e45-8f79-351e5c354563",  
  "status": 200,  
  "result": {  
    "orderId": 336829446,  
    "symbol": "BTCUSDT",  
    "status": "NEW",  
    "clientOrderId": "FqEw6cn0vDhrkmfiwLYPeo",  
    "price": "42088.00",  
    "avgPrice": "0.00",  
    "origQty": "0.100",  
    "executedQty": "0.000",  
    "cumQty": "0.000",  
    "cumQuote": "0.00000",  
    "timeInForce": "GTC",  
    "type": "LIMIT",  
    "reduceOnly": false,  
    "closePosition": false,  
    "side": "BUY",  
    "positionSide": "BOTH",  
    "stopPrice": "0.00",  
    "workingType": "CONTRACT_PRICE",  
    "priceProtect": false,  
    "origType": "LIMIT",  
    "priceMatch": "NONE",  
    "selfTradePreventionMode": "NONE",  
    "goodTillDate": 0,  
    "updateTime": 1705385954229  
  },  
  "rateLimits": [  
    {  
      "rateLimitType": "REQUEST_WEIGHT",  
      "interval": "MINUTE",  
      "intervalNum": 1,  
      "limit": 2400,  
      "count": 1  
    },  
    {  
      "rateLimitType": "ORDERS",  
      "interval": "SECOND",  
      "intervalNum": 10,  
      "limit": 300,  
      "count": 1  
    },  
    {  
      "rateLimitType": "ORDERS",  
      "interval": "MINUTE",  
      "intervalNum": 1,  
      "limit": 1200,  
      "count": 0  
    }  
  ]  
}  

```

> 失败响应示例:
```
{  
  "id": "5761b939-27b1-4948-ab87-4a372a3f6b72",  
  "status": 400,  
  "error": {  
    "code": -1102,  
    "msg": "Mandatory parameter 'quantity' was not sent, was empty/null, or malformed."  
  },  
  "rateLimits": [  
    {  
      "rateLimitType": "REQUEST_WEIGHT",  
      "interval": "MINUTE",  
      "intervalNum": 1,  
      "limit": 2400,  
      "count": 1  
    },  
    {  
      "rateLimitType": "ORDERS",  
      "interval": "SECOND",  
      "intervalNum": 10,  
      "limit": 300,  
      "count": 1  
    },  
    {  
      "rateLimitType": "ORDERS",  
      "interval": "MINUTE",  
      "intervalNum": 1,  
      "limit": 1200,  
      "count": 1  
    }  
  ]  
}  

```

响应字段:
名称 | 类型 | 是否必需 | 描述  
---|---|---|---  
`id` | INT / STRING / `null` | YES | 与原来请求的ID一样  
`status` | INT | YES | 响应状态。请看 [状态代码](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#%E7%8A%B6%E6%80%81%E4%BB%A3%E7%A0%81)  
`result` | OBJECT / ARRAY | YES | 响应内容。请求成功则显示  
`error` | OBJECT | 错误描述。请求失败则显示  
`rateLimits` | ARRAY | NO | 速率限制状态。请看 [速率限制](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#%E9%80%9F%E7%8E%87%E9%99%90%E5%88%B6)  
## WebSocket API访问限制[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#websocket-api%E8%AE%BF%E9%97%AE%E9%99%90%E5%88%B6 "WebSocket API访问限制的直接链接")
  * 速率限制与 REST API 相同，并且与 REST API 共享。
  * WebSocket 握手尝试消耗权重5。
  * ping/pong 帧的速率限制：每秒最多5次。
  * 默认情况下，响应中包含速率限制信息，参见 `rateLimits` 字段。
  * 可以通过在连接字符串或单个请求中的 `returnRateLimits` 布尔参数来控制 `rateLimits` 字段的可见性。
  * 例如，使用 `wss://ws-fapi.binance.com/ws-fapi/v1?returnRateLimits=false` 默认在响应中隐藏 `rateLimits`。在这样的情况下，您可以在请求中传递额外的 `"returnRateLimits": true` 参数，当默认隐藏时，就可以在响应中显示速率限制。


## WebSocket API连接后进行身份验证[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#websocket-api%E8%BF%9E%E6%8E%A5%E5%90%8E%E8%BF%9B%E8%A1%8C%E8%BA%AB%E4%BB%BD%E9%AA%8C%E8%AF%81 "WebSocket API连接后进行身份验证的直接链接")
你可以使用会话身份验证请求对已经建立的连接进行身份验证：
  * [`session.logon`](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#log-in-with-api-key-signed) – 进行身份验证，或更改与连接相关联的API密钥。
  * [`session.status`](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#query-session-status) – 检查连接状态和当前API密钥。
  * [`session.logout`](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#log-out-of-the-session) – 忘记与连接关联的API密钥。


## WebSocket API关于吊销API密钥[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#websocket-api%E5%85%B3%E4%BA%8E%E5%90%8A%E9%94%80api%E5%AF%86%E9%92%A5 "WebSocket API关于吊销API密钥的直接链接")
如果在活动会话期间，由于 _任何_ 原因（例如IP地址未被加入白名单、API密钥被删除、API密钥没有正确的权限等），在下一个请求后，会话将被吊销，并显示以下错误消息:
```
{  
  "id": null,  
  "status": 401,  
  "error": {  
    "code": -2015,  
    "msg": "Invalid API-key, IP, or permissions for action."   
  }  
}  

```

## WebSocket API授权 _临时_ 请求[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#websocket-api%E6%8E%88%E6%9D%83-%E4%B8%B4%E6%97%B6-%E8%AF%B7%E6%B1%82 "websocket-api授权-临时-请求的直接链接")
WebSocket连接只能通过一个API密钥进行身份验证。 默认情况下，经过身份验证的API密钥将用于需要`apiKey`参数的请求。 但是，你始终可以为单个请求明确指定`apiKey`和`signature`，覆盖已认证的API密钥，以使用不同的API密钥授权特定请求。
例如，你可能希望用默认密钥来验证 `USER_DATA`，但在下单时使用`TRADE`密钥来签名。
## WebSocket API身份验证请求[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#websocket-api%E8%BA%AB%E4%BB%BD%E9%AA%8C%E8%AF%81%E8%AF%B7%E6%B1%82 "WebSocket API身份验证请求的直接链接")
**注意：** 仅支持 _Ed25519_ 密钥用于此功能。
### 用API key登录 (SIGNED)[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#%E7%94%A8api-key%E7%99%BB%E5%BD%95-signed "用API key登录 \(SIGNED\)的直接链接")
> **请求**
```
{  
  "id": "c174a2b1-3f51-4580-b200-8528bd237cb7",  
  "method": "session.logon",  
  "params": {  
    "apiKey": "vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",  
    "signature": "1cf54395b336b0a9727ef27d5d98987962bc47aca6e13fe978612d0adee066ed",  
    "timestamp": 1649729878532  
  }  
}  

```

> **响应:**
```
{  
  "id": "c174a2b1-3f51-4580-b200-8528bd237cb7",  
  "status": 200,  
  "result": {  
    "apiKey": "vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",  
    "authorizedSince": 1649729878532,  
    "connectedSince": 1649729873021,  
    "returnRateLimits": false,  
    "serverTime": 1649729878630  
  }  
}  

```

使用提供的API密钥进行WebSocket连接身份验证。
在调用`session.logon`后，将来的需要`apiKey`和`signature`参数的请求可以省略它们。
请注意，只能认证一个API密钥。 多次调用`session.logon`将更改当前已认证的API密钥。
**权重:** 2
**参数:**
参数名 | 类型 | 是否必需 | 描述  
---|---|---|---  
`apiKey` | STRING | YES |   
`recvWindow` | INT | NO | The value cannot be greater than `60000`  
`signature` | STRING | YES |   
`timestamp` | INT | YES |   
**数据源:** 缓存
### 查询会话状态[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#%E6%9F%A5%E8%AF%A2%E4%BC%9A%E8%AF%9D%E7%8A%B6%E6%80%81 "查询会话状态的直接链接")
> **请求**
```
{  
  "id": "b50c16cd-62c9-4e29-89e4-37f10111f5bf",  
  "method": "session.status"  
}  

```

> **响应:**
```
{  
  "id": "b50c16cd-62c9-4e29-89e4-37f10111f5bf",  
  "status": 200,  
  "result": {  
    // 如果连接未经身份验证，"apiKey" 和 "authorizedSince" 将显示为 null。  
    "apiKey": "vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",  
    "authorizedSince": 1649729878532,  
    "connectedSince": 1649729873021,  
    "returnRateLimits": false,  
    "serverTime": 1649730611671  
  }  
}  

```

查询WebSocket连接的状态，检查用于授权请求的API密钥（如果有的话）。
**权重:** 2
**参数:** NONE
**数据源:** 缓存
### 退出会话[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#%E9%80%80%E5%87%BA%E4%BC%9A%E8%AF%9D "退出会话的直接链接")
> **请求**
```
{  
  "id": "c174a2b1-3f51-4580-b200-8528bd237cb7",  
  "method": "session.logout"  
}  

```

> **响应:**
```
{  
  "id": "c174a2b1-3f51-4580-b200-8528bd237cb7",  
  "status": 200,  
  "result": {  
    "apiKey": null,  
    "authorizedSince": null,  
    "connectedSince": 1649729873021,  
    "returnRateLimits": false,  
    "serverTime": 1649730611671  
  }  
}  

```

忘记之前认证的API密钥。 如果连接未经身份验证，此请求不会有任何作用。
请注意，`session.logout`请求后，WebSocket连接仍然保持打开状态。 你可以继续使用连接，但现在必须在需要的地方明确提供`apiKey`和`signature`参数。
**权重:** 2
**参数:** NONE
**数据源:** 缓存
## SIGNED (TRADE and USER_DATA) 请求安全[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#signed-trade-and-user_data-%E8%AF%B7%E6%B1%82%E5%AE%89%E5%85%A8 "SIGNED \(TRADE and USER_DATA\) 请求安全的直接链接")
### SIGNED 请求示例 (Ed25519)[​](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#signed-%E8%AF%B7%E6%B1%82%E7%A4%BA%E4%BE%8B-ed25519 "SIGNED 请求示例 \(Ed25519\)的直接链接")
*我们建议使用 Ed25519 API keys**，因为它在所有受支持的 API key 类型中提供最佳性能和安全性。  参数 | 取值  
---|---  
symbol | BTCUSDT  
side | SELL  
type | LIMIT  
timeInForce | GTC  
quantity | 1  
price | 0.2  
timestamp | 1668481559918  
```
#!/usr/bin/env python3  
import base64  
import time  
import json  
from cryptography.hazmat.primitives.serialization import load_pem_private_key  
from websocket import create_connection  
# 设置身份验证：  
API_KEY='替换成您的 API Key'  
PRIVATE_KEY_PATH='test-prv-key.pem'  
# 加载 private key。  
# 在这个例子中，private key 没有加密，但我们建议使用强密码以提高安全性。  
with open(PRIVATE_KEY_PATH, 'rb') as f:  
    private_key = load_pem_private_key(data=f.read(), password=None)  
# 设置请求参数：  
params = {  
    'apiKey':        API_KEY,	  
    'symbol':       'BTCUSDT',  
    'side':         'SELL',  
    'type':         'LIMIT',  
    'timeInForce':  'GTC',  
    'quantity':     '1.0000000',  
    'price':        '0.20'  
}  
# 参数中加时间戳：  
timestamp = int(time.time() * 1000) # 以毫秒为单位的 UNIX 时间戳  
params['timestamp'] = timestamp  
# 参数中加签名：  
payload = '&'.join([f'{param}={value}' for param, value in sorted(params.items())])  
signature = base64.b64encode(private_key.sign(payload.encode('ASCII')))  
params['signature'] = signature.decode('ASCII')  
# 发送请求：  
request = {	  
    'id': 'my_new_order',	  
    'method': 'order.place',	  
    'params': params  
}  
ws = create_connection('wss://ws-fapi.binance.com/ws-fapi/v1')	  
ws.send(json.dumps(request))	  
result =  ws.recv()	  
ws.close()	  
print(result)  

```

右边有 Python 脚本来示例如何使用 Ed25519 key 签名。
[上一页 基本信息](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/general-info)[下一页 通用枚举定义](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/common-definition)
  * [WebSocket API响应格式](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#websocket-api%E5%93%8D%E5%BA%94%E6%A0%BC%E5%BC%8F)
  * [WebSocket API访问限制](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#websocket-api%E8%AE%BF%E9%97%AE%E9%99%90%E5%88%B6)
  * [WebSocket API连接后进行身份验证](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#websocket-api%E8%BF%9E%E6%8E%A5%E5%90%8E%E8%BF%9B%E8%A1%8C%E8%BA%AB%E4%BB%BD%E9%AA%8C%E8%AF%81)
  * [WebSocket API关于吊销API密钥](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#websocket-api%E5%85%B3%E4%BA%8E%E5%90%8A%E9%94%80api%E5%AF%86%E9%92%A5)
  * [WebSocket API授权 _临时_ 请求](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#websocket-api%E6%8E%88%E6%9D%83-%E4%B8%B4%E6%97%B6-%E8%AF%B7%E6%B1%82)
  * [WebSocket API身份验证请求](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#websocket-api%E8%BA%AB%E4%BB%BD%E9%AA%8C%E8%AF%81%E8%AF%B7%E6%B1%82)
    * [用API key登录 (SIGNED)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#%E7%94%A8api-key%E7%99%BB%E5%BD%95-signed)
    * [查询会话状态](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#%E6%9F%A5%E8%AF%A2%E4%BC%9A%E8%AF%9D%E7%8A%B6%E6%80%81)
    * [退出会话](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#%E9%80%80%E5%87%BA%E4%BC%9A%E8%AF%9D)
  * [SIGNED (TRADE and USER_DATA) 请求安全](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#signed-trade-and-user_data-%E8%AF%B7%E6%B1%82%E5%AE%89%E5%85%A8)
    * [SIGNED 请求示例 (Ed25519)](https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/websocket-api-general-info#signed-%E8%AF%B7%E6%B1%82%E7%A4%BA%E4%BE%8B-ed25519)


Copyright © 2026 Binance.
