# Futures Insight Drawer (单币详情指标)

## 目标
在 `futures/symbols` 表格的“操作”列新增【指标】按钮，点击后打开右侧抽屉展示单币多维指标图表，覆盖截图中的 8 个指标。

## UI 入口与布局
- 入口：`futures/symbols` 表格行操作区新增【指标】按钮。
- 展示：右侧抽屉 `FuturesInsightDrawer`，顶部含币种与时间粒度切换（5m/15m/30m/1h/4h）。
- 布局：3x3 栅格（8 张卡片，最后一行 2 张）。

## 指标与数据来源
- 合约持仓量（OI）+ 名义价值（USDT）：USDM Open Interest History
- 大户账户数多空比：Top Long/Short Account Ratio
- 大户持仓量多空比：Top Long/Short Position Ratio
- 多空账户数比：Global Long/Short Account Ratio
- 合约主动买卖量：Taker Buy/Sell Volume
- 基差：markPrice - spotPrice（使用 mark price Kline + spot Kline 同周期对齐）
- 资金费率历史：Funding Rate History
- OI/市值比：暂不接市值源，返回空值/提示

## API 设计
新增聚合接口：
`GET /futures/insight?symbol=BTCUSDT&interval=5m&limit=60`

返回结构（简化）：
- `oi_hist`: [{time, open_interest, open_interest_value}]
- `top_account_ratio`: [{time, ratio}]
- `top_position_ratio`: [{time, ratio}]
- `global_account_ratio`: [{time, ratio}]
- `taker_volume`: [{time, buy_vol, sell_vol}]
- `basis`: [{time, basis, mark_price, spot_price}]
- `funding_rate`: [{time, rate}]
- `oi_marketcap_ratio`: []

## 性能与缓存
- 统一聚合接口减少前端请求数。
- 服务端缓存 60-120 秒（key: symbol+interval+limit）。

## 前端实现
- 新组件：`frontend/src/views/futures/components/InsightDrawer.vue`
- ECharts 渲染 8 张图表，统一暗色主题、tooltip、空态。
- 切换粒度后重新请求接口。

## 约束
- 市值不接外部源，OI/市值比先空置。
- 基差按 markPrice - spotPrice 计算。
