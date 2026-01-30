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
