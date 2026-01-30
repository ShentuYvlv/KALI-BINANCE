<template>
  <div class="kline-page">
    <div class="kline-header">
      <div class="kline-header-left">
        <button class="icon-btn" @click="goBack" title="返回">
          <i class="el-icon-back" />
        </button>
        <div class="symbol-pill">
          <span class="symbol-dot" />
          <span class="symbol-text">{{ symbol }}</span>
        </div>
      </div>
      <div class="kline-header-center">
        <div class="intervals">
          <button
            v-for="item in intervalOptions"
            :key="item.value"
            class="interval-btn"
            :class="{ active: currentInterval === item.value, disabled: item.disabled }"
            @click="changeInterval(item.value)"
          >
            {{ item.label }}
          </button>
        </div>
      </div>
      <div class="kline-header-right">
        <el-dropdown trigger="click" @command="toggleIndicator">
          <button class="header-btn">
            <i class="el-icon-view" />
            <span>指标</span>
            <i class="el-icon-arrow-down el-icon--right" />
          </button>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item v-for="item in indicatorOptions" :key="item.value" :command="item.value">
              <span class="dropdown-label">{{ item.label }}</span>
              <i v-if="activeIndicators[item.value]" class="el-icon-check dropdown-check" />
            </el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
        <el-dropdown trigger="click" @command="changeTimezone">
          <button class="header-btn">
            <i class="el-icon-time" />
            <span>时区</span>
            <i class="el-icon-arrow-down el-icon--right" />
          </button>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item v-for="tz in timezoneOptions" :key="tz.value" :command="tz.value">
              <span class="dropdown-label">{{ tz.label }}</span>
              <i v-if="timezone === tz.value" class="el-icon-check dropdown-check" />
            </el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
        <el-popover placement="bottom" width="240" trigger="click">
          <div class="settings-row">
            <span>网格</span>
            <el-switch v-model="settings.showGrid" @change="applySettings" />
          </div>
          <div class="settings-row">
            <span>十字线</span>
            <el-switch v-model="settings.showCrosshair" @change="applySettings" />
          </div>
          <div class="settings-row">
            <span>成交量</span>
            <el-switch v-model="settings.showVolume" @change="applySettings" />
          </div>
          <div class="settings-row">
            <span>MACD</span>
            <el-switch v-model="settings.showMacd" @change="applySettings" />
          </div>
          <button class="header-btn" slot="reference">
            <i class="el-icon-setting" />
            <span>设置</span>
          </button>
        </el-popover>
        <button class="header-btn" @click="downloadSnapshot">
          <i class="el-icon-picture-outline" />
          <span>截图</span>
        </button>
        <button class="header-btn" @click="toggleFullscreen">
          <i class="el-icon-rank" />
          <span>{{ isFullscreen ? '退出全屏' : '全屏' }}</span>
        </button>
      </div>
    </div>
    <div class="kline-body" ref="fullscreenRoot">
      <div class="kline-orderbook">
        <div class="side-section">
          <div class="side-header">
            <div class="side-title-text">订单簿</div>
            <div class="side-actions">
              <span v-if="depthStatus !== 'connected'" class="status-pill">{{ statusLabel(depthStatus) }}</span>
              <i class="el-icon-more" />
              <i class="el-icon-close" />
            </div>
          </div>
          <div class="orderbook">
            <div class="orderbook-top">
              <span class="precision-pill">{{ tickSize || '0.00001' }}</span>
            </div>
            <div class="orderbook-header">
              <span>价格 (USDT)</span>
              <span>数量 (USDT)</span>
              <span>合计 (USDT)</span>
            </div>
            <div class="orderbook-body">
              <div class="orderbook-list asks">
                <div v-for="(item, index) in depthAsks" :key="'ask-' + index" class="orderbook-row sell">
                  <span class="depth-bg" :style="{ width: (item.depthPct || 0) + '%' }" />
                  <span class="price">{{ formatNumber(item.price, pricePrecision) }}</span>
                  <span class="qty">{{ formatNumber(item.qty, qtyPrecision) }}</span>
                  <span class="total">{{ formatNumber(item.total, qtyPrecision) }}</span>
                </div>
              </div>
              <div class="orderbook-mid" :class="lastTradeSide">
                <span class="mid-price">{{ formatNumber(lastTradePrice, pricePrecision) }}</span>
                <span class="mid-ref">{{ formatNumber(lastTradePrice, pricePrecision) }}</span>
              </div>
              <div class="orderbook-list bids">
                <div v-for="(item, index) in depthBids" :key="'bid-' + index" class="orderbook-row buy">
                  <span class="depth-bg" :style="{ width: (item.depthPct || 0) + '%' }" />
                  <span class="price">{{ formatNumber(item.price, pricePrecision) }}</span>
                  <span class="qty">{{ formatNumber(item.qty, qtyPrecision) }}</span>
                  <span class="total">{{ formatNumber(item.total, qtyPrecision) }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="kline-trades">
        <div class="side-section trades">
          <div class="side-header">
            <div class="side-tabs">
              <button class="tab-btn" :class="{ active: tradeTab === 'trades' }" @click="tradeTab = 'trades'">
                最新成交
              </button>
              <button class="tab-btn" :class="{ active: tradeTab === 'flow' }" @click="tradeTab = 'flow'">
                市场异动
              </button>
            </div>
            <div class="side-actions">
              <i class="el-icon-rank" />
              <i class="el-icon-more" />
            </div>
          </div>
          <div class="trades-header">
            <span>价格 (USDT)</span>
            <span>数量 (USDT)</span>
            <span>时间</span>
          </div>
          <div class="trades-list" v-if="tradeTab === 'trades'">
            <div
              v-for="(item, index) in trades"
              :key="item.id || 'trade-' + index"
              class="trade-row"
              :class="[item.side, { flash: item.flash, big: item.big }]"
            >
              <span class="price">{{ formatNumber(item.price, pricePrecision) }}</span>
              <span class="qty">{{ formatNumber(item.qty, qtyPrecision) }}</span>
              <span class="time">{{ formatTradeTime(item.time) }}</span>
            </div>
          </div>
          <div v-else class="empty-tip">暂无数据</div>
        </div>
      </div>
      <div class="kline-tools">
        <div class="tool-group">
          <div class="tool-title">绘图</div>
          <button
            v-for="tool in drawTools"
            :key="tool.name"
            class="tool-icon-btn"
            :class="{ active: activeTool === tool.name }"
            :title="tool.label"
            @click="activateTool(tool.name)"
          >
            <i :class="tool.icon" />
          </button>
        </div>
        <div class="tool-group">
          <div class="tool-title">操作</div>
          <button class="tool-icon-btn" title="清除全部" @click="clearOverlays">
            <i class="el-icon-delete" />
          </button>
        </div>
      </div>
      <div class="kline-chart-wrap">
        <div class="kline-info-bar">
          <div class="info-left">
            <span class="info-item">时间: <strong>{{ infoBar.time }}</strong></span>
            <span class="info-item">开: <strong>{{ infoBar.open }}</strong></span>
            <span class="info-item">高: <strong>{{ infoBar.high }}</strong></span>
            <span class="info-item">低: <strong>{{ infoBar.low }}</strong></span>
            <span class="info-item">收: <strong>{{ infoBar.close }}</strong></span>
            <span class="info-item">成交量: <strong>{{ infoBar.volume }}</strong></span>
          </div>
          <div v-if="infoBar.maValues.length" class="info-right">
            <span
              v-for="item in infoBar.maValues"
              :key="item.label"
              class="info-ma"
              :style="{ color: item.color }"
            >
              {{ item.label }}: {{ item.value }}
            </span>
          </div>
        </div>
        <div
          v-show="crosshairLabel.visible"
          class="crosshair-time"
          :style="{ left: crosshairLabel.x + 'px' }"
        >
          {{ crosshairLabel.text }}
        </div>
        <div ref="chart" class="kline-chart" />
      </div>
    </div>
  </div>
</template>

<script>
import { init, dispose } from 'klinecharts'
import { getFuturesDepth, getFuturesKline, getFuturesTrades } from '@/api/trade'

export default {
  name: 'FuturesKline',
  data() {
    return {
      symbol: '',
      tickSize: '',
      stepSize: '',
      currentInterval: '1m',
      chart: null,
      ws: null,
      wsDepth: null,
      wsTrades: null,
      activeTool: '',
      timezone: 'Asia/Shanghai',
      settings: {
        showGrid: true,
        showCrosshair: true,
        showVolume: true,
        showMacd: true,
      },
      indicatorIds: {},
      activeIndicators: {},
      infoBar: {
        time: '--',
        open: '--',
        high: '--',
        low: '--',
        close: '--',
        volume: '--',
        maValues: [],
      },
      depthBids: [],
      depthAsks: [],
      trades: [],
      tradeTab: 'trades',
      depthLimit: 20,
      tradesLimit: 100,
      depthMaxTotal: 1,
      lastTradePrice: null,
      lastTradeSide: '',
      tradeFlashMs: 1000,
      largeTradeNotional: 50000,
      depthStatus: 'disconnected',
      tradesStatus: 'disconnected',
      depthReconnectDelay: 1000,
      tradesReconnectDelay: 1000,
      depthReconnectTimer: null,
      tradesReconnectTimer: null,
      depthPollTimer: null,
      tradesPollTimer: null,
      restPollInterval: 3000,
      marketActive: true,
      pricePrecision: 2,
      qtyPrecision: 3,
      crosshairLabel: {
        visible: false,
        x: 0,
        text: '',
      },
      mouseEventsBound: false,
      lastCrosshairIndex: null,
      isFullscreen: false,
      intervalOptions: [
        { label: '1m', value: '1m' },
        { label: '5m', value: '5m' },
        { label: '15m', value: '15m' },
        { label: '1H', value: '1h' },
        { label: '2H', value: '2h' },
        { label: '4H', value: '4h' },
        { label: 'D', value: '1d' },
        { label: 'W', value: '1w' },
        { label: 'M', value: '1M' },
        { label: 'Y', value: '1y', disabled: true },
      ],
      indicatorOptions: [
        { label: 'MA', value: 'MA' },
        { label: 'EMA', value: 'EMA' },
        { label: 'BOLL', value: 'BOLL' },
        { label: 'SAR', value: 'SAR' },
        { label: 'VOL', value: 'VOL' },
        { label: 'MACD', value: 'MACD' },
        { label: 'RSI', value: 'RSI' },
        { label: 'KDJ', value: 'KDJ' },
      ],
      timezoneOptions: [
        { label: 'Asia/Shanghai', value: 'Asia/Shanghai' },
        { label: 'UTC', value: 'UTC' },
        { label: 'America/New_York', value: 'America/New_York' },
      ],
      drawTools: [
        { label: '游标', name: 'cursor', icon: 'el-icon-view' },
        { label: '趋势线', name: 'straightLine', icon: 'el-icon-edit' },
        { label: '射线', name: 'rayLine', icon: 'el-icon-arrow-right' },
        { label: '线段', name: 'segment', icon: 'el-icon-minus' },
        { label: '水平线', name: 'horizontalStraightLine', icon: 'el-icon-more' },
        { label: '垂直线', name: 'verticalStraightLine', icon: 'el-icon-menu' },
        { label: '平行线', name: 'parallelStraightLine', icon: 'el-icon-rank' },
        { label: '价格通道', name: 'priceChannelLine', icon: 'el-icon-tickets' },
        { label: '斐波那契', name: 'fibonacciLine', icon: 'el-icon-share' },
        { label: '标注', name: 'simpleAnnotation', icon: 'el-icon-info' },
        { label: '标签', name: 'simpleTag', icon: 'el-icon-document' },
      ],
    }
  },
  mounted() {
    this.symbol = (this.$route.query.symbol || '').toUpperCase()
    this.tickSize = this.$route.query.tickSize || ''
    this.stepSize = this.$route.query.stepSize || ''
    const interval = this.$route.query.interval
    if (interval) {
      this.currentInterval = interval
    }
    this.initChart()
  },
  beforeDestroy() {
    this.marketActive = false
    this.closeWs()
    this.closeMarketWs()
    this.clearReconnectTimers()
    this.stopDepthPolling()
    this.stopTradesPolling()
    if (this.chart) {
      dispose(this.$refs.chart)
      this.chart = null
    }
  },
  methods: {
    goBack() {
      this.$router.back()
    },
    initChart() {
      if (!this.symbol) {
        this.$message.error('symbol不能为空')
        return
      }
      this.chart = init(this.$refs.chart, {
        locale: 'zh-CN',
        timezone: this.timezone,
        styles: 'dark',
        zoomAnchor: 'cursor_point',
      })
      this.applyFormatters()
      this.applyTheme()
      this.applySettings()
      this.setupDefaultIndicators()
      this.setupDataLoader()
      this.subscribeCrosshair()
      this.applySymbol()
      this.applyPeriod()
      this.initResizeObserver()
      this.$nextTick(() => this.initChartMouseEvents())
      this.connectMarketWs()
    },
    initChartMouseEvents() {
      if (this.mouseEventsBound) return
      const chartEl = this.$refs.chart
      if (!chartEl) return
      this.mouseEventsBound = true
      const yAxisDom = this.chart?.getDom?.('candle_pane', 'yAxis') || null
      const onLeave = () => {
        this.crosshairLabel.visible = false
        if (this.chart?.getChartStore) {
          this.chart.getChartStore().setCrosshair()
        }
      }
      let rafId = null
      let lastPoint = null
      const onMove = event => {
        lastPoint = { x: event.clientX, y: event.clientY }
        if (rafId) return
        rafId = window.requestAnimationFrame(() => {
          rafId = null
          if (!lastPoint || !this.chart) return
          const rect = chartEl.getBoundingClientRect()
          const x = lastPoint.x - rect.left
          const y = lastPoint.y - rect.top
          if (x < 0 || y < 0 || x > rect.width || y > rect.height) {
            this.crosshairLabel.visible = false
            this.chart.getChartStore().setCrosshair()
            return
          }
          let paneId = 'candle_pane'
          const panes = this.chart.getDrawPanes?.() || []
          for (const pane of panes) {
            const bound = pane.getBounding?.()
            if (bound && y >= bound.top && y <= bound.bottom) {
              paneId = pane.getId()
              break
            }
          }
          this.chart.getChartStore().setCrosshair({ x, y, paneId }, { forceInvalidate: true })
        })
      }
      chartEl.addEventListener('mouseleave', onLeave)
      window.addEventListener('mousemove', onMove)
      const onYAxisWheel = event => {
        if (!this.chart) return
        event.preventDefault()
        const pane = this.chart.getDrawPaneById?.('candle_pane')
        const yAxis = pane?.getAxisComponent?.()
        if (!yAxis) return
        const range = yAxis.getRange()
        const delta = Math.sign(event.deltaY || 0) * Math.min(1, Math.abs((event.deltaY || 0) / 100))
        if (!delta) return
        const scale = 1 + delta * 0.1
        const newRange = range.range * scale
        const difRange = (newRange - range.range) / 2
        const newFrom = range.from - difRange
        const newTo = range.to + difRange
        const newRealFrom = yAxis.valueToRealValue(newFrom, { range })
        const newRealTo = yAxis.valueToRealValue(newTo, { range })
        const newDisplayFrom = yAxis.realValueToDisplayValue(newRealFrom, { range })
        const newDisplayTo = yAxis.realValueToDisplayValue(newRealTo, { range })
        yAxis.setRange({
          from: newFrom,
          to: newTo,
          range: newRange,
          realFrom: newRealFrom,
          realTo: newRealTo,
          realRange: newRealTo - newRealFrom,
          displayFrom: newDisplayFrom,
          displayTo: newDisplayTo,
          displayRange: newDisplayTo - newDisplayFrom,
        })
        this.chart.layout({
          measureWidth: true,
          update: true,
          buildYAxisTick: true,
        })
      }
      if (yAxisDom) {
        yAxisDom.addEventListener('wheel', onYAxisWheel, { passive: false })
      }
      this.$once('hook:beforeDestroy', () => {
        chartEl.removeEventListener('mouseleave', onLeave)
        window.removeEventListener('mousemove', onMove)
        if (yAxisDom) {
          yAxisDom.removeEventListener('wheel', onYAxisWheel)
        }
        if (rafId) {
          window.cancelAnimationFrame(rafId)
          rafId = null
        }
        this.mouseEventsBound = false
      })
    },
    applyFormatters() {
      if (!this.chart) return
      this.chart.setFormatter({
        formatDate: ({ timestamp, type, dateTimeFormat }) => {
          if (type === 'crosshair') {
            return this.formatCrosshairTime(timestamp)
          }
          return dateTimeFormat.format(new Date(timestamp))
        },
      })
    },
    initResizeObserver() {
      const resize = () => {
        this.chart && this.chart.resize()
      }
      window.addEventListener('resize', resize)
      this.$once('hook:beforeDestroy', () => {
        window.removeEventListener('resize', resize)
      })
    },
    applyTheme() {
      if (!this.chart) return
      this.chart.setStyles({
        grid: {
          horizontal: { color: '#1f2937' },
          vertical: { color: '#1f2937' },
        },
        candle: {
          bar: {
            upColor: '#2ac37d',
            downColor: '#f05454',
            upBorderColor: '#2ac37d',
            downBorderColor: '#f05454',
            upWickColor: '#2ac37d',
            downWickColor: '#f05454',
          },
          priceMark: {
            last: {
              upColor: '#2ac37d',
              downColor: '#f05454',
            },
          },
          tooltip: {
            showRule: 'none',
          },
        },
        indicator: {
          tooltip: {
            showRule: 'none',
          },
        },
        xAxis: {
          axisLine: { color: '#1f2937' },
          tickLine: { color: '#1f2937' },
          tickText: { color: '#8a98ad' },
        },
        yAxis: {
          axisLine: { color: '#1f2937' },
          tickLine: { color: '#1f2937' },
          tickText: { color: '#8a98ad' },
        },
        separator: {
          color: '#1f2937',
        },
        crosshair: {
          show: true,
          horizontal: {
            show: true,
            line: { show: true, color: '#4b5563' },
            text: {
              show: true,
              color: '#e2e8f0',
              backgroundColor: '#111827',
              borderColor: '#111827',
            },
          },
          vertical: {
            show: true,
            line: { show: true, color: '#4b5563' },
            text: {
              show: true,
              color: '#e2e8f0',
              backgroundColor: '#111827',
              borderColor: '#111827',
            },
          },
        },
      })
    },
    setupDataLoader() {
      this.chart.setDataLoader({
        getBars: async ({ type, timestamp, period, symbol, callback }) => {
          const interval = this.periodToInterval(period)
          const limit = type === 'init' ? 500 : 300
          let endTime = 0
          if (type === 'forward' && timestamp) {
            // 向左加载更早的数据
            endTime = timestamp - 1
          }
          if (type === 'backward') {
            // 右侧不再加载，避免旧数据追加到末尾造成时间倒序
            callback([], { forward: false, backward: false })
            return
          }
          const resp = await getFuturesKline({
            symbol: symbol.ticker,
            interval,
            limit,
            endTime: endTime || undefined,
          })
          const data = Array.isArray(resp?.data) ? resp.data : []
          const hasMore = data.length >= limit
          callback(data, { forward: hasMore, backward: false })
          if (type === 'init') {
            this.$nextTick(() => this.updateInfoFromLast())
          }
        },
        subscribeBar: ({ symbol, period, callback }) => {
          const interval = this.periodToInterval(period)
          this.connectWs(symbol.ticker, interval, callback)
        },
        unsubscribeBar: () => {
          this.closeWs()
        },
      })
    },
    subscribeCrosshair() {
      this.chart.subscribeAction('onCrosshairChange', () => {
        const crosshair = this.chart.getChartStore().getCrosshair()
        if (crosshair?.kLineData) {
          const idx = crosshair.realDataIndex ?? crosshair.dataIndex
          if (idx !== this.lastCrosshairIndex) {
            this.lastCrosshairIndex = idx
            this.updateInfo(crosshair.kLineData, idx)
          }
          this.updateCrosshairLabel(crosshair)
          return
        }
        this.crosshairLabel.visible = false
        this.lastCrosshairIndex = null
      })
    },
    setupDefaultIndicators() {
      this.indicatorIds = {}
      this.activeIndicators = {}
      this.addIndicator('MA', { id: 'candle_pane' })
      this.addIndicator('VOL', { height: 100 })
      this.addIndicator('MACD', { height: 120 })
      this.settings.showVolume = true
      this.settings.showMacd = true
    },
    addIndicator(name, paneOptions) {
      const isOverlay = paneOptions && paneOptions.id === 'candle_pane'
      const id = this.chart.createIndicator(name, false, paneOptions)
      if (id) {
        this.indicatorIds[name] = id
        this.activeIndicators[name] = true
      }
      if (isOverlay) {
        this.activeIndicators[name] = true
      }
    },
    updateInfoFromLast() {
      if (!this.chart) return
      const dataList = this.chart.getDataList()
      if (!dataList || !dataList.length) return
      const lastIndex = dataList.length - 1
      const last = dataList[lastIndex]
      this.updateInfo(last, lastIndex)
    },
    updateInfo(kLineData, dataIndex) {
      const pricePrecision = this.getPrecision(this.tickSize)
      const volumePrecision = this.getPrecision(this.stepSize)
      this.infoBar = {
        time: this.formatTime(kLineData.timestamp),
        open: this.formatNumber(kLineData.open, pricePrecision),
        high: this.formatNumber(kLineData.high, pricePrecision),
        low: this.formatNumber(kLineData.low, pricePrecision),
        close: this.formatNumber(kLineData.close, pricePrecision),
        volume: this.formatNumber(kLineData.volume, volumePrecision),
        maValues: this.getMaValues(dataIndex, pricePrecision),
      }
    },
    updateCrosshairLabel(crosshair) {
      const chartEl = this.$refs.chart
      if (!chartEl) return
      const width = chartEl.clientWidth || 0
      const x = crosshair.realX ?? crosshair.x
      if (x === undefined || x === null || !crosshair.timestamp) {
        this.crosshairLabel.visible = false
        return
      }
      const clamp = (val, min, max) => Math.max(min, Math.min(max, val))
      const safeX = clamp(x, 40, Math.max(40, width - 40))
      this.crosshairLabel = {
        visible: true,
        x: safeX,
        text: this.formatCrosshairTime(crosshair.timestamp),
      }
    },
    getMaValues(dataIndex, precision) {
      if (dataIndex === undefined || dataIndex === null) {
        return []
      }
      if (!this.activeIndicators.MA) {
        return []
      }
      const list = this.chart.getIndicators({ name: 'MA' }) || []
      if (!list.length) return []
      const indicator = list[0]
      const result = indicator.result?.[dataIndex] || {}
      const colors = ['#f59e0b', '#60a5fa', '#ef4444', '#8b5cf6']
      return indicator.calcParams.map((p, idx) => {
        const key = `ma${idx + 1}`
        const value = result[key]
        return {
          label: `MA${p}`,
          value: value === undefined ? '--' : this.formatNumber(value, precision),
          color: colors[idx % colors.length],
        }
      })
    },
    removeIndicator(name) {
      const id = this.indicatorIds[name]
      if (id) {
        this.chart.removeIndicator({ id })
        this.indicatorIds[name] = null
      } else {
        this.chart.removeIndicator({ name })
      }
      this.activeIndicators[name] = false
    },
    toggleIndicator(name) {
      if (this.activeIndicators[name]) {
        this.removeIndicator(name)
        if (name === 'MA') {
          this.infoBar.maValues = []
        }
        this.updateInfoFromLast()
        return
      }
      if (['MA', 'EMA', 'BOLL', 'SAR'].includes(name)) {
        this.addIndicator(name, { id: 'candle_pane' })
        this.updateInfoFromLast()
        return
      }
      if (name === 'VOL') {
        this.addIndicator(name, { height: 100 })
        this.updateInfoFromLast()
        return
      }
      this.addIndicator(name, { height: 120 })
      this.updateInfoFromLast()
    },
    applySymbol() {
      const pricePrecision = this.getPrecision(this.tickSize)
      const volumePrecision = this.getPrecision(this.stepSize)
      this.pricePrecision = pricePrecision || 2
      this.qtyPrecision = volumePrecision || 0
      this.chart.setSymbol({
        ticker: this.symbol,
        pricePrecision: this.pricePrecision,
        volumePrecision: this.qtyPrecision,
      })
    },
    applyPeriod() {
      this.chart.setPeriod(this.intervalToPeriod(this.currentInterval))
    },
    changeInterval(interval) {
      const option = this.intervalOptions.find(item => item.value === interval)
      if (option?.disabled) {
        this.$message.warning('该周期暂不支持')
        return
      }
      if (interval === this.currentInterval) return
      this.currentInterval = interval
      this.applyPeriod()
      this.chart.resetData()
    },
    changeTimezone(value) {
      this.timezone = value
      this.chart.setTimezone(value)
    },
    applySettings() {
      if (!this.chart) return
      this.chart.setStyles({
        grid: {
          horizontal: { show: this.settings.showGrid },
          vertical: { show: this.settings.showGrid },
        },
        crosshair: {
          show: this.settings.showCrosshair,
        },
      })
      if (this.settings.showVolume && !this.activeIndicators.VOL) {
        this.addIndicator('VOL', { height: 100 })
      }
      if (!this.settings.showVolume && this.activeIndicators.VOL) {
        this.removeIndicator('VOL')
      }
      if (this.settings.showMacd && !this.activeIndicators.MACD) {
        this.addIndicator('MACD', { height: 120 })
      }
      if (!this.settings.showMacd && this.activeIndicators.MACD) {
        this.removeIndicator('MACD')
      }
    },
    activateTool(name) {
      if (name === 'cursor') {
        this.activeTool = ''
        return
      }
      this.activeTool = name
      this.chart.createOverlay({ name })
    },
    clearOverlays() {
      this.chart.removeOverlay()
      this.activeTool = ''
    },
    downloadSnapshot() {
      const url = this.chart.getConvertPictureUrl(true, 'png', '#0f1115')
      if (!url) return
      const link = document.createElement('a')
      link.href = url
      link.download = `${this.symbol}-${this.currentInterval}.png`
      link.click()
    },
    toggleFullscreen() {
      const target = this.$refs.fullscreenRoot
      if (!target) return
      if (!document.fullscreenElement) {
        target.requestFullscreen()
        this.isFullscreen = true
        return
      }
      document.exitFullscreen()
      this.isFullscreen = false
    },
    connectWs(symbol, interval, callback) {
      this.closeWs()
      const url = this.buildWsUrl(`/ws/futures/kline?symbol=${symbol}&interval=${interval}`)
      this.ws = new WebSocket(url)
      this.ws.onmessage = event => {
        try {
          const payload = JSON.parse(event.data || '{}')
          if (payload?.data) {
            callback(payload.data)
            this.$nextTick(() => this.updateInfoFromLast())
          }
        } catch (e) {
          // ignore invalid message
        }
      }
    },
    connectMarketWs() {
      if (!this.symbol) return
      this.marketActive = true
      this.connectDepthWs(this.symbol, this.depthLimit)
      this.connectTradesWs(this.symbol, this.tradesLimit)
    },
    statusLabel(status) {
      switch (status) {
        case 'connected':
          return 'WS已连接'
        case 'connecting':
          return '连接中'
        case 'reconnecting':
          return '重连中'
        case 'polling':
          return 'REST轮询'
        default:
          return '已断开'
      }
    },
    clearReconnectTimers() {
      if (this.depthReconnectTimer) {
        clearTimeout(this.depthReconnectTimer)
        this.depthReconnectTimer = null
      }
      if (this.tradesReconnectTimer) {
        clearTimeout(this.tradesReconnectTimer)
        this.tradesReconnectTimer = null
      }
    },
    startDepthPolling() {
      if (this.depthPollTimer) return
      this.depthStatus = this.depthStatus === 'connected' ? 'connected' : 'polling'
      this.fetchDepthSnapshot()
      this.depthPollTimer = setInterval(() => {
        this.fetchDepthSnapshot()
      }, this.restPollInterval)
    },
    scheduleDepthReconnect() {
      if (this.depthReconnectTimer || !this.marketActive) return
      const delay = this.depthReconnectDelay
      this.depthReconnectTimer = setTimeout(() => {
        this.depthReconnectTimer = null
        this.depthReconnectDelay = Math.min(delay * 2, 10000)
        if (this.marketActive && this.depthStatus !== 'connected') {
          this.connectDepthWs(this.symbol, this.depthLimit)
        }
      }, delay)
    },
    stopDepthPolling() {
      if (this.depthPollTimer) {
        clearInterval(this.depthPollTimer)
        this.depthPollTimer = null
      }
    },
    startTradesPolling() {
      if (this.tradesPollTimer) return
      this.tradesStatus = this.tradesStatus === 'connected' ? 'connected' : 'polling'
      this.fetchTradesSnapshot()
      this.tradesPollTimer = setInterval(() => {
        this.fetchTradesSnapshot()
      }, this.restPollInterval)
    },
    scheduleTradesReconnect() {
      if (this.tradesReconnectTimer || !this.marketActive) return
      const delay = this.tradesReconnectDelay
      this.tradesReconnectTimer = setTimeout(() => {
        this.tradesReconnectTimer = null
        this.tradesReconnectDelay = Math.min(delay * 2, 10000)
        if (this.marketActive && this.tradesStatus !== 'connected') {
          this.connectTradesWs(this.symbol, this.tradesLimit)
        }
      }, delay)
    },
    stopTradesPolling() {
      if (this.tradesPollTimer) {
        clearInterval(this.tradesPollTimer)
        this.tradesPollTimer = null
      }
    },
    async fetchDepthSnapshot() {
      if (!this.symbol) return
      try {
        const resp = await getFuturesDepth({ symbol: this.symbol, limit: this.depthLimit })
        const data = resp?.data || {}
        const bids = Array.isArray(data.bids) ? data.bids : []
        const asks = Array.isArray(data.asks) ? data.asks : []
        this.applyDepthSnapshot(bids, asks)
      } catch (e) {
        // ignore
      }
    },
    async fetchTradesSnapshot() {
      if (!this.symbol) return
      try {
        const resp = await getFuturesTrades({ symbol: this.symbol, limit: this.tradesLimit })
        const items = Array.isArray(resp?.data?.items) ? resp.data.items : []
        this.applyTradesSnapshot(items)
      } catch (e) {
        // ignore
      }
    },
    applyDepthSnapshot(bids, asks) {
      const buildDepth = list => {
        let total = 0
        return list.map(item => {
          const qty = Number(item.qty ?? item.Quantity ?? item[1])
          const price = Number(item.price ?? item.Price ?? item[0])
          total += qty
          return {
            price,
            qty,
            total,
          }
        })
      }
      const nextBids = buildDepth(bids)
      const nextAsks = buildDepth(asks)
      const maxTotal = Math.max(
        nextBids.length ? nextBids[nextBids.length - 1].total : 0,
        nextAsks.length ? nextAsks[nextAsks.length - 1].total : 0,
        1
      )
      const applyDepthPct = item => ({
        ...item,
        depthPct: Math.min((item.total / maxTotal) * 100, 100),
      })
      this.depthMaxTotal = maxTotal
      this.depthBids = nextBids.map(applyDepthPct)
      this.depthAsks = nextAsks.map(applyDepthPct)
    },
    applyTradesSnapshot(items) {
      const normalized = items
        .map(item => ({ ...this.buildTradeItem(item), flash: false }))
        .sort((a, b) => b.time - a.time)
      this.trades = normalized.slice(0, this.tradesLimit)
      if (this.trades.length) {
        this.lastTradePrice = this.trades[0].price
        this.lastTradeSide = this.trades[0].side
      }
    },
    connectDepthWs(symbol, limit) {
      this.closeDepthWs()
      if (!this.marketActive) return
      const url = this.buildWsUrl(`/ws/futures/depth?symbol=${symbol}&limit=${limit}`)
      this.depthStatus = 'connecting'
      const ws = new WebSocket(url)
      this.wsDepth = ws
      ws.onopen = () => {
        this.depthStatus = 'connected'
        this.depthReconnectDelay = 1000
        this.stopDepthPolling()
      }
      ws.onerror = () => {
        if (!this.marketActive) return
        if (!this.depthPollTimer) {
          this.depthStatus = 'reconnecting'
        }
        this.startDepthPolling()
        this.scheduleDepthReconnect()
      }
      ws.onclose = () => {
        if (!this.marketActive) return
        if (!this.depthPollTimer) {
          this.depthStatus = 'reconnecting'
        }
        this.startDepthPolling()
        this.scheduleDepthReconnect()
      }
      ws.onmessage = event => {
        try {
          const payload = JSON.parse(event.data || '{}')
          if (!payload?.data) return
          const bids = Array.isArray(payload.data.bids) ? payload.data.bids : []
          const asks = Array.isArray(payload.data.asks) ? payload.data.asks : []
          this.applyDepthSnapshot(bids, asks)
        } catch (e) {
          // ignore invalid message
        }
      }
    },
    connectTradesWs(symbol, limit) {
      this.closeTradesWs()
      const url = this.buildWsUrl(`/ws/futures/trades?symbol=${symbol}&limit=${limit}`)
      if (!this.marketActive) return
      this.tradesStatus = 'connecting'
      const ws = new WebSocket(url)
      this.wsTrades = ws
      ws.onopen = () => {
        this.tradesStatus = 'connected'
        this.tradesReconnectDelay = 1000
        this.stopTradesPolling()
      }
      ws.onerror = () => {
        if (!this.marketActive) return
        if (!this.tradesPollTimer) {
          this.tradesStatus = 'reconnecting'
        }
        this.startTradesPolling()
        this.scheduleTradesReconnect()
      }
      ws.onclose = () => {
        if (!this.marketActive) return
        if (!this.tradesPollTimer) {
          this.tradesStatus = 'reconnecting'
        }
        this.startTradesPolling()
        this.scheduleTradesReconnect()
      }
      ws.onmessage = event => {
        try {
          const payload = JSON.parse(event.data || '{}')
          if (payload?.type === 'snapshot' && Array.isArray(payload.data)) {
            this.applyTradesSnapshot(payload.data)
            return
          }
          if (payload?.type === 'update' && payload?.data) {
            const next = this.buildTradeItem(payload.data)
            this.trades.unshift(next)
            if (this.trades.length > this.tradesLimit) {
              this.trades = this.trades.slice(0, this.tradesLimit)
            }
            this.lastTradePrice = next.price
            this.lastTradeSide = next.side
            if (next.flash) {
              const id = next.id
              setTimeout(() => {
                const target = this.trades.find(item => item.id === id)
                if (target) target.flash = false
              }, this.tradeFlashMs)
            }
          }
        } catch (e) {
          // ignore invalid message
        }
      }
    },
    closeDepthWs() {
      if (this.wsDepth) {
        this.wsDepth.onclose = null
        this.wsDepth.onerror = null
        this.wsDepth.close()
        this.wsDepth = null
      }
    },
    closeTradesWs() {
      if (this.wsTrades) {
        this.wsTrades.onclose = null
        this.wsTrades.onerror = null
        this.wsTrades.close()
        this.wsTrades = null
      }
    },
    closeMarketWs() {
      this.closeDepthWs()
      this.closeTradesWs()
      this.depthStatus = 'disconnected'
      this.tradesStatus = 'disconnected'
    },
    closeWs() {
      if (this.ws) {
        this.ws.close()
        this.ws = null
      }
    },
    buildWsUrl(path) {
      const base = process.env.VUE_APP_BASE_API || ''
      if (base.startsWith('http')) {
        return base.replace(/^http/, 'ws') + path
      }
      const protocol = window.location.protocol === 'https:' ? 'wss' : 'ws'
      return `${protocol}://${window.location.host}${base}${path}`
    },
    formatTime(ts) {
      if (!ts) return '--'
      const date = new Date(ts)
      const pad = val => String(val).padStart(2, '0')
      return `${date.getFullYear()}-${pad(date.getMonth() + 1)}-${pad(date.getDate())} ${pad(date.getHours())}:${pad(date.getMinutes())}`
    },
    formatCrosshairTime(ts) {
      if (!ts) return '--'
      const date = new Date(ts)
      const pad = val => String(val).padStart(2, '0')
      return `${pad(date.getMonth() + 1)}-${pad(date.getDate())} ${pad(date.getHours())}:${pad(date.getMinutes())}`
    },
    formatNumber(value, precision = 2) {
      if (value === undefined || value === null || Number.isNaN(Number(value))) return '--'
      const num = Number(value)
      return num.toFixed(Math.min(Math.max(precision, 0), 8))
    },
    formatTradeTime(ts) {
      if (!ts) return '--'
      const date = new Date(Number(ts))
      const pad = val => String(val).padStart(2, '0')
      return `${pad(date.getHours())}:${pad(date.getMinutes())}`
    },
    buildTradeItem(raw) {
      const price = Number(raw.price)
      const qty = Number(raw.qty)
      const time = Number(raw.time)
      const side = raw.maker ? 'sell' : 'buy'
      const notional = price * qty
      return {
        id: `${time}-${price}-${qty}-${side}`,
        price,
        qty,
        time,
        side,
        notional,
        big: notional >= this.largeTradeNotional,
        flash: true,
      }
    },
    intervalToPeriod(interval) {
      const num = parseInt(interval, 10)
      if (interval.endsWith('m')) {
        return { type: 'minute', span: num }
      }
      if (interval.endsWith('h')) {
        return { type: 'hour', span: num }
      }
      if (interval.endsWith('d')) {
        return { type: 'day', span: num }
      }
      if (interval.endsWith('w')) {
        return { type: 'week', span: num }
      }
      if (interval.endsWith('M')) {
        return { type: 'month', span: num }
      }
      if (interval.endsWith('y')) {
        return { type: 'year', span: num }
      }
      return { type: 'minute', span: 1 }
    },
    periodToInterval(period) {
      const span = period.span || 1
      if (period.type === 'minute') return `${span}m`
      if (period.type === 'hour') return `${span}h`
      if (period.type === 'day') return `${span}d`
      if (period.type === 'week') return `${span}w`
      if (period.type === 'month') return `${span}M`
      if (period.type === 'year') return `${span}y`
      return '1m'
    },
    getPrecision(step) {
      if (!step) return 0
      const str = String(step)
      if (str.indexOf('1e-') !== -1) {
        const parts = str.split('1e-')
        return parseInt(parts[1], 10)
      }
      const dot = str.indexOf('.')
      if (dot === -1) return 0
      return str.length - dot - 1
    },
  },
}
</script>

<style lang="scss" scoped>
.kline-page {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background: #0f1115;
  color: #d5dae2;
}

.kline-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 16px;
  background: #0c1017;
  border-bottom: 1px solid #1b2230;
}

.kline-header-left {
  display: flex;
  align-items: center;
  gap: 10px;
}

.icon-btn {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  border: 1px solid #263041;
  background: #141a23;
  color: #9fb0c6;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s ease;
}

.icon-btn:hover {
  border-color: #365072;
  color: #d5dae2;
}

.symbol-pill {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 6px 12px;
  background: #131a24;
  border: 1px solid #263041;
  border-radius: 999px;
  font-weight: 600;
}

.symbol-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #f59e0b;
}

.symbol-text {
  font-size: 13px;
  letter-spacing: 0.3px;
}

.kline-header-center {
  display: flex;
  flex: 1;
  justify-content: center;
}

.intervals {
  display: flex;
  gap: 6px;
  flex-wrap: wrap;
}

.interval-btn {
  border: 1px solid transparent;
  background: transparent;
  color: #9fb0c6;
  padding: 4px 8px;
  border-radius: 6px;
  font-size: 12px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.interval-btn.active {
  background: #16253a;
  color: #4fd3a8;
  border-color: #2b7a5f;
}

.interval-btn:hover {
  color: #d5dae2;
  background: #141a23;
}

.interval-btn.disabled {
  color: #4b5563;
  cursor: not-allowed;
  background: transparent;
}

.kline-header-right {
  display: flex;
  align-items: center;
  gap: 8px;
}

.header-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 6px 10px;
  border-radius: 8px;
  background: #141a23;
  border: 1px solid #263041;
  color: #cbd5e1;
  font-size: 12px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.header-btn:hover {
  border-color: #365072;
  color: #e2e8f0;
}

.dropdown-label {
  margin-right: 6px;
}

.dropdown-check {
  color: #3fe3a4;
}

.kline-body {
  display: flex;
  flex: 1;
  overflow: hidden;
}

.kline-trades,
.kline-orderbook {
  width: 320px;
  background: #0b0f14;
  border-right: 1px solid #1b2230;
  display: flex;
  flex-direction: column;
  padding: 8px;
  overflow: hidden;
}

.kline-trades {
  width: 320px;
}

.kline-tools {
  width: 56px;
  background: #0d1117;
  border-right: 1px solid #1b2230;
  padding: 10px 8px;
  overflow-y: auto;
}

.kline-orderbook .side-section {
  flex: 1;
}

.side-section {
  display: flex;
  flex-direction: column;
  min-height: 0;
  background: #0f131a;
  border: 1px solid #1c232f;
  border-radius: 10px;
  padding: 8px;
}

.side-section.trades {
  flex: 1;
}

.side-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 8px;
}

.side-title-text {
  font-size: 12px;
  color: #e5e7eb;
  font-weight: 600;
}

.side-tabs {
  display: flex;
  gap: 12px;
}

.tab-btn {
  background: transparent;
  border: none;
  color: #8b95a7;
  font-size: 12px;
  padding: 0 0 6px;
  cursor: pointer;
  position: relative;
}

.tab-btn.active {
  color: #f0b90b;
}

.tab-btn.active::after {
  content: '';
  position: absolute;
  left: 0;
  right: 0;
  bottom: -2px;
  height: 2px;
  background: #f0b90b;
  border-radius: 2px;
}

.side-actions {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #6b7280;
}

.side-actions i {
  font-size: 14px;
  cursor: pointer;
}

.side-actions i:hover {
  color: #cbd5e1;
}

.status-pill {
  font-size: 10px;
  padding: 2px 6px;
  border-radius: 999px;
  background: rgba(251, 191, 36, 0.12);
  color: #fbbf24;
}

.status {
  font-size: 10px;
  padding: 2px 6px;
  border-radius: 999px;
  background: #1b2230;
  color: #9fb0c6;
}

.status.connected {
  color: #4ade80;
  background: rgba(74, 222, 128, 0.12);
}

.status.reconnecting,
.status.connecting,
.status.polling {
  color: #fbbf24;
  background: rgba(251, 191, 36, 0.12);
}

.status.disconnected {
  color: #f87171;
  background: rgba(248, 113, 113, 0.12);
}

.orderbook {
  display: flex;
  flex-direction: column;
  gap: 6px;
  flex: 1;
  min-height: 0;
}

.orderbook-top {
  display: flex;
  justify-content: flex-end;
}

.precision-pill {
  font-size: 11px;
  color: #9aa4b2;
  padding: 2px 8px;
  border-radius: 10px;
  background: #141922;
  border: 1px solid #252b36;
}

.orderbook-header {
  display: grid;
  grid-template-columns: minmax(90px, 1.2fr) minmax(80px, 1fr) minmax(90px, 1fr);
  column-gap: 8px;
  font-size: 11px;
  color: #6f7a8c;
  padding: 2px 4px;
}

.orderbook-header span:first-child,
.trades-header span:first-child {
  text-align: left;
}

.orderbook-header span:not(:first-child),
.trades-header span:not(:first-child) {
  text-align: right;
}

.orderbook-body {
  display: grid;
  grid-template-rows: minmax(0, 1fr) auto minmax(0, 1fr);
  flex: 1;
  min-height: 0;
  gap: 6px;
}

.orderbook-list {
  display: flex;
  flex-direction: column;
  gap: 4px;
  height: 100%;
  min-height: 0;
  overflow: auto;
  position: relative;
  padding-right: 2px;
}

.orderbook-list.asks::before,
.orderbook-list.bids::before {
  content: '';
  position: absolute;
  inset: 0;
  pointer-events: none;
  z-index: 0;
}

.orderbook-list.asks::before {
  background: linear-gradient(180deg, rgba(246, 70, 93, 0.12), rgba(246, 70, 93, 0));
}

.orderbook-list.bids::before {
  background: linear-gradient(0deg, rgba(14, 203, 129, 0.12), rgba(14, 203, 129, 0));
}

.orderbook-row {
  display: grid;
  grid-template-columns: minmax(90px, 1.2fr) minmax(80px, 1fr) minmax(90px, 1fr);
  column-gap: 8px;
  font-size: 12px;
  position: relative;
  padding: 1px 4px;
  border-radius: 4px;
  overflow: hidden;
  background: transparent;
  font-variant-numeric: tabular-nums;
  white-space: nowrap;
  align-items: center;
}

.orderbook-row .depth-bg {
  position: absolute;
  top: 0;
  bottom: 0;
  right: 0;
  background: rgba(246, 70, 93, 0.22);
  z-index: 0;
}

.orderbook-row.buy .depth-bg {
  left: 0;
  right: auto;
  background: rgba(14, 203, 129, 0.22);
}

.orderbook-row span {
  position: relative;
  z-index: 1;
  min-width: 0;
}

.orderbook-row.sell .price {
  color: #f6465d;
}

.orderbook-row.buy .price {
  color: #0ecb81;
}

.orderbook-row .price {
  text-align: left;
}

.orderbook-row .qty,
.orderbook-row .total {
  text-align: right;
}

.orderbook-row .qty {
  color: #cbd5e1;
}

.orderbook-row .total {
  color: #9fb0c6;
}

.orderbook-mid {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 6px 4px;
  border-radius: 0;
  background: transparent;
  border-top: 1px solid #1b2230;
  border-bottom: 1px solid #1b2230;
  margin: 2px 0;
}

.orderbook-mid .mid-price {
  font-size: 18px;
  font-weight: 700;
}

.orderbook-mid .mid-ref {
  font-size: 12px;
  color: #7c8798;
}

.orderbook-mid.buy .mid-price {
  color: #0ecb81;
}

.orderbook-mid.sell .mid-price {
  color: #f6465d;
}

.trades-header {
  display: grid;
  grid-template-columns: minmax(90px, 1.2fr) minmax(90px, 1fr) minmax(80px, 1fr);
  column-gap: 8px;
  font-size: 11px;
  color: #6f7a8c;
  margin-bottom: 6px;
  padding: 2px 4px;
}

.trades-list {
  display: flex;
  flex-direction: column;
  gap: 4px;
  overflow: auto;
  min-height: 0;
  padding-right: 2px;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

.trades-list::-webkit-scrollbar {
  width: 0;
  height: 0;
}

.trade-row {
  display: grid;
  grid-template-columns: minmax(90px, 1.2fr) minmax(90px, 1fr) minmax(80px, 1fr);
  column-gap: 8px;
  font-size: 12px;
  padding: 1px 4px;
  border-radius: 4px;
  font-variant-numeric: tabular-nums;
  white-space: nowrap;
  align-items: center;
  background: transparent;
}

.trade-row span {
  min-width: 0;
}

.empty-tip {
  color: #6b7280;
  font-size: 12px;
  padding: 8px 0;
}

.trade-row .price {
  text-align: left;
}

.trade-row .qty,
.trade-row .time {
  text-align: right;
}

.trade-row.buy .price {
  color: #0ecb81;
}

.trade-row.sell .price {
  color: #f6465d;
}

.trade-row .qty,
.trade-row .time {
  color: #cbd5e1;
}

.trade-row.flash {
  animation: trade-flash 1s ease-out;
}

.trade-row.big {
  background: rgba(255, 255, 255, 0.06);
}

@keyframes trade-flash {
  0% {
    background: rgba(59, 130, 246, 0.25);
  }
  100% {
    background: transparent;
  }
}

.tool-group {
  margin-bottom: 14px;
}

.tool-group + .tool-group {
  border-top: 1px solid #1b2230;
  padding-top: 10px;
}

.tool-title {
  font-size: 10px;
  color: #6d7a8d;
  margin-bottom: 6px;
  text-align: center;
}

.tool-icon-btn {
  width: 100%;
  height: 36px;
  border: 1px solid #263041;
  background: #131a24;
  color: #a6b2c2;
  border-radius: 8px;
  font-size: 16px;
  cursor: pointer;
  margin-bottom: 6px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
}

.tool-icon-btn:hover {
  border-color: #365072;
  color: #d5dae2;
}
.tool-icon-btn.active {
  background: #16253a;
  color: #4fd3a8;
  border-color: #2b7a5f;
}

.kline-chart-wrap {
  flex: 1;
  background: #0f1115;
  padding: 0;
  position: relative;
  min-width: 0;
}

.kline-chart {
  width: 100%;
  height: 100%;
}

.crosshair-time {
  position: absolute;
  bottom: 8px;
  transform: translateX(-50%);
  background: #1f2937;
  color: #e2e8f0;
  font-size: 12px;
  padding: 4px 8px;
  border-radius: 4px;
  border: 1px solid #1f2937;
  z-index: 3;
  pointer-events: none;
}

.kline-info-bar {
  position: absolute;
  top: 10px;
  left: 16px;
  right: 16px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 12px;
  color: #9fb0c6;
  gap: 16px;
  pointer-events: none;
  z-index: 2;
}

.info-left {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.info-item strong {
  color: #e2e8f0;
  font-weight: 600;
}

.info-right {
  display: flex;
  align-items: center;
  gap: 12px;
}

.info-ma {
  font-weight: 600;
}

.settings-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 10px;
  color: #d5dae2;
}
</style>
