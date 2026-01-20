启动教程：
## 配置文件

```
cp ./conf/app.conf.example app.conf
```

### 安装 bee
> 记得将`GOPATH/bin`添加到环境变量`PATH`，否则 `bee` 命令无法全局使用
> 使用 `go env GOPATH` 查看 `GOPATH` 路径

```
go install github.com/beego/bee/v2@latest
```

### 安装依赖
> 进入项目根目录下执行

```
go mod tidy
```

### 启动
> http://localhost:3333/zmkm/index.html

```
bee run
```

### 打包

#### 打包到`windows`平台
> 其它平台需要参考 bee 文档
> 此项目的 github 的 workflows 实现了 linux amd64 和 window amd64 下的编译打包

```
bee pack -be GOOS=windows
```

## web ui 开发
> https://github.com/sorry510/go_binance_futures_ui

### TODO

- [X] 完成新币抢购功能
- [X] 完成挖矿新币抛售功能
- [X] 添加独立的币种配置收益率
- [X] 添加一键修改所有币种的配置
- [X] 系统首页显示(那些服务开启和关闭)
- [X] 监听币种的价格突变情况，报警通知
- [X] 学习蜡烛图结合其它数据，报警通知
- [X] 添加新的自动交易策略
- [X] 批量配置自定义策略，使用模板方式导入，增加一个模板页面(可以导入策略模板)
- [X] 更新代码提示功能
- [X] 将配置从 `conf` 缩减，改为可视化配置实时生效
- [X] 数据库更新方式，第一次自动生成 db 文件，如果存在数据库，就根据版本号进行更新语句更新
- [X] 现货和合约添加 tab，添加 usdt 之外的交易对
- [X] 抢购添加手动设定价格挂单的功能，不设定才采用市价抢单
- [ ] api 重构，添加一层适配器用来统一支持其它交易所的接口
- [ ] 替换现有的币安 sdk 换用官方的
- [ ] 添加币本位合约
- [ ] 接入币安的订单接口(优点:精确化操控合约的每一次开仓，不再对当前仓位的合约全部平仓, 缺点:手动操作的合约，无法自动根据策略完成平仓)
- [ ] 添加定时自动交易(现货买入和x倍合约等值对冲，吃资金费用)
- [ ] 监控资金流入流出，报警通知
- [X] 自动伸缩功能，当连续亏损时，是否需要自动缩减 最大持仓数量，反之亦然，类似于 tcp 窗口缩放
- [ ] 当买入30min中内无明显变化时，及时平仓


### 检查 binance 和 x 的新闻
> https://www.binance.com/zh-CN/square/profile/xxx

#### 关键词
- 马斯克
- 活动
- 竞赛
- 比赛

#### binance 账号
- [ ] binance_announcement
- [ ] binance_news



