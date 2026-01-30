# KLine UI 变更记录

> 目标：记录 K 线页面与相关入口的所有 UI/功能变更（后续继续追加）。

## 2026-01-29
### 已完成
- 新增合约 K 线页面：`/#/futures/symbols/kline`，点击合约列表币种打开新标签页。
- K 线数据接入后端：新增 REST `/futures/kline` 与 WS `/ws/futures/kline`。
- 前端接入 KLineChart 并提供完整 UI 控件：周期切换、指标、时区、设置、截图、全屏、画线工具栏。
- UI 精修（对齐官方 demo 结构）：顶栏布局分区、周期按钮对齐、工具栏图标化、暗色主题适配。

### 2026-01-29（续）
- 顶栏布局调整为左右/中间分区，周期按钮样式更贴近官方 demo。
- 左侧绘图工具栏改为图标化紧凑布局。
- 动作按钮改为图标+文本形式，配色与暗色主题统一。

### 2026-01-29（细节对齐）
- 新增图表顶部信息行（时间/开高低收/成交量）与 MA 数值展示。
- 调整图表暗色主题细节（网格、轴线、K线色彩）。
- 周期列表对齐官方 demo（1m/5m/15m/1H/2H/4H/D/W/M/Y）。
- Y 周期保留展示但禁用（币安不支持 1y）。

### 相关文件
- 后端：`controllers/futuresKline.go`，`feature/api/binance/index.go`，`routers/router.go`
- 前端：`frontend/src/views/futures/Kline.vue`，`frontend/src/views/futures/index.vue`
- API：`frontend/src/api/trade.js`
- 路由与文案：`frontend/src/router/modules/trade.js`，`frontend/src/lang/zh.js`，`frontend/src/lang/en.js`
- 依赖：`frontend/package.json`（新增 `klinecharts`）

## 2026-01-30
- K 线页数据加载方向修正：历史数据仅在 forward（向左拖动）加载，避免旧数据追加到末尾导致时间回跳。
- K 线页面全屏化：隐藏全局左侧菜单与顶部导航/标签栏，内容区高度调整为 100vh。
- K 线页交互优化：右侧价格轴支持滚轮缩放；内置 KLineChart tooltip 关闭，避免与自定义信息栏重复。
- K 线页新增左侧行情面板：订单簿（50档）与最新成交（100条），使用 WS 推送。
- 行情面板细化：订单簿改为 20 档、增加“最新价”分隔行与累计量列，深度条背景按累计量渲染。
- 最新成交优化：按成交额 ≥ 50,000 USDT 标记大单，新增 1s 闪烁高亮，时间格式调整为 HH:mm。
- 行情面板可用性增强：增加 WS 连接状态提示；WS 断开时自动重连；REST 兜底轮询（3s）补数据。
- 行情面板布局调整：绘图工具栏改为悬浮叠加，左侧区域优先展示订单簿/成交数据。
- 行情面板排版优化：订单簿/成交采用三列栅格与等宽数字，缓解数字挤压。
- 面板重排：最新成交独立左列，订单簿独立右列，绘图工具栏贴在订单簿右侧、紧邻图表。
- 订单簿/成交样式重构：顶部 Tabs/动作区、盘口精度胶囊、表头单位化与网格对齐、盘口上下均分并突出中间价。
