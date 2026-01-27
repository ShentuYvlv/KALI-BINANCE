<template>
  <el-drawer
    :visible.sync="drawerVisible"
    :with-header="false"
    size="1400px"
    direction="rtl"
    custom-class="insight-drawer"
  >
    <div class="insight-wrapper" v-loading="loading">
      <div class="insight-toolbar">
        <div class="insight-title">
          <span class="insight-symbol">{{ symbol }}</span>
          <span class="insight-sub">{{ $t('trade.insight') }}</span>
        </div>
        <div class="insight-controls">
          <el-select v-model="interval" size="mini" class="insight-select" @change="fetchData">
            <el-option v-for="item in intervals" :key="item" :label="item" :value="item" />
          </el-select>
          <el-button size="mini" type="primary" icon="el-icon-refresh" @click="fetchData">{{ $t('table.refresh') }}</el-button>
        </div>
      </div>

      <el-row :gutter="12" class="insight-grid">
        <el-col v-for="card in cards" :key="card.key" :span="8">
          <div class="insight-card">
            <div class="insight-card__title">{{ card.title }}</div>
            <div v-if="card.empty" class="insight-empty">{{ card.emptyText }}</div>
            <insight-chart v-else :options="card.options" height="100%" />
          </div>
        </el-col>
      </el-row>
    </div>
  </el-drawer>
</template>

<script>
import { getFuturesInsight } from '@/api/trade'
import { parseTime } from '@/utils'
import InsightChart from './InsightChart'

export default {
  name: 'InsightDrawer',
  components: { InsightChart },
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    symbol: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      loading: false,
      interval: '5m',
      intervals: ['5m', '15m', '30m', '1h', '4h'],
      limit: 120,
      raw: {
        oi_hist: [],
        top_account_ratio: [],
        top_position_ratio: [],
        global_account_ratio: [],
        taker_volume: [],
        basis: [],
        funding_rate: [],
        oi_marketcap_ratio: [],
      }
    }
  },
  computed: {
    drawerVisible: {
      get() {
        return this.visible
      },
      set(val) {
        this.$emit('update:visible', val)
      }
    },
    cards() {
      return [
        {
          key: 'oi',
          title: '合约持仓量',
          options: this.buildOIOptions(),
          empty: this.raw.oi_hist.length === 0,
          emptyText: '暂无数据',
        },
        {
          key: 'topAccount',
          title: '大户账户数多空比',
          options: this.buildRatioOptions(this.raw.top_account_ratio, '多空账户比值'),
          empty: this.raw.top_account_ratio.length === 0,
          emptyText: '暂无数据',
        },
        {
          key: 'topPosition',
          title: '大户持仓量多空比',
          options: this.buildRatioOptions(this.raw.top_position_ratio, '多空持仓比值'),
          empty: this.raw.top_position_ratio.length === 0,
          emptyText: '暂无数据',
        },
        {
          key: 'globalAccount',
          title: '多空账户数比',
          options: this.buildRatioOptions(this.raw.global_account_ratio, '多空账户比值'),
          empty: this.raw.global_account_ratio.length === 0,
          emptyText: '暂无数据',
        },
        {
          key: 'taker',
          title: '合约主动买卖量',
          options: this.buildTakerOptions(),
          empty: this.raw.taker_volume.length === 0,
          emptyText: '暂无数据',
        },
        {
          key: 'basis',
          title: '基差',
          options: this.buildBasisOptions(),
          empty: this.raw.basis.length === 0,
          emptyText: '暂无数据',
        },
        {
          key: 'funding',
          title: '资金费率',
          options: this.buildFundingOptions(),
          empty: this.raw.funding_rate.length === 0,
          emptyText: '暂无数据',
        },
        {
          key: 'oiMarketCap',
          title: '未平仓量与市值比率',
          options: this.buildRatioOptions(this.raw.oi_marketcap_ratio, 'OI/市值'),
          empty: true,
          emptyText: '未配置市值源',
        }
      ]
    }
  },
  watch: {
    visible(val) {
      if (val) {
        this.fetchData()
      }
    },
    symbol(val) {
      if (this.visible && val) {
        this.fetchData()
      }
    }
  },
  methods: {
    async fetchData() {
      if (!this.symbol) {
        return
      }
      this.loading = true
      try {
        const { data } = await getFuturesInsight({
          symbol: this.symbol,
          interval: this.interval,
          limit: this.limit,
        })
        this.raw = data || this.raw
        if (this.raw && Array.isArray(this.raw.funding_rate)) {
          this.raw.funding_rate = this.raw.funding_rate.slice(-40)
        }
      } catch (e) {
        this.$message.error('获取指标失败')
      } finally {
        this.loading = false
      }
    },
    formatTime(ts) {
      return parseTime(ts, '{m}-{d} {h}:{i}')
    },
    formatTimeShort(ts) {
      return parseTime(ts, '{h}:{i}')
    },

    formatCompact(value) {
      const num = Number(value)
      if (Number.isNaN(num)) {
        return value
      }
      const abs = Math.abs(num)
      if (abs >= 1e12) return `${(num / 1e12).toFixed(2)}T`
      if (abs >= 1e9) return `${(num / 1e9).toFixed(2)}B`
      if (abs >= 1e6) return `${(num / 1e6).toFixed(2)}M`
      if (abs >= 1e3) return `${(num / 1e3).toFixed(2)}K`
      return num.toFixed(2)
    },
    formatDecimal(value, digits = 2) {
      const num = Number(value)
      if (Number.isNaN(num)) {
        return value
      }
      return num.toFixed(digits)
    },
    calcAxisInterval(count) {
      if (!count || count <= 6) {
        return 0
      }
      const step = Math.ceil(count / 6) - 1
      return Math.max(1, step)
    },

    formatTooltip(params) {
      if (!params || !params.length) {
        return ''
      }
      const title = params[0].axisValue
      const lines = params.map(item => {
        const value = Array.isArray(item.value) ? item.value[1] : item.value
        let formatted = value
        if (item.seriesName.includes('USDT') || item.seriesName.includes('买入') || item.seriesName.includes('卖出') || item.seriesName.includes('持仓量') || item.seriesName.includes('价格')) {
          formatted = this.formatCompact(value)
        } else if (item.seriesName.includes('基差')) {
          formatted = this.formatDecimal(value, 2)
        } else if (item.seriesName.includes('费率')) {
          formatted = `${this.formatDecimal(value, 3)}%`
        } else {
          formatted = this.formatDecimal(value, 2)
        }
        return `${item.marker}${item.seriesName}: ${formatted}`
      })
      return `${title}<br>${lines.join('<br>')}`
    },


    buildVolumeYAxis() {
      return {
        type: 'value',
        min: 0,
        max: (value) => {
          if (value.max <= 0) return 1
          return value.max * 1.1
        },
        axisLabel: { color: '#94a3b8', formatter: (v) => this.formatCompact(v) },
        splitNumber: 5,
        axisLine: { lineStyle: { color: '#2f3a45' } },
        splitLine: { lineStyle: { color: 'rgba(148, 163, 184, 0.08)' } },
      }
    },
    buildScaledYAxis(formatter) {
      return {
        type: 'value',
        scale: true,
        min: (value) => {
          if (value.min === value.max) {
            const pad = Math.max(Math.abs(value.min) * 0.01, 0.000001)
            return value.min - pad
          }
          return value.min - (value.max - value.min) * 0.08
        },
        max: (value) => {
          if (value.min === value.max) {
            const pad = Math.max(Math.abs(value.max) * 0.01, 0.000001)
            return value.max + pad
          }
          return value.max + (value.max - value.min) * 0.08
        },
        axisLabel: { color: '#94a3b8', formatter },
        splitNumber: 5,
        axisLine: { lineStyle: { color: '#2f3a45' } },
        axisTick: { show: false },
        splitLine: { lineStyle: { color: 'rgba(148, 163, 184, 0.08)' } },
      }
    },

    buildBaseOptions(xAxisData, yAxis, series, xAxisInterval = 0, showLegend = true) {
      const gridBottom = showLegend ? 30 : 16
      return {
        backgroundColor: 'transparent',
        grid: {
          left: 8,
          right: 8,
          top: 24,
          bottom: gridBottom,
          containLabel: true,
        },
        tooltip: {
          trigger: 'axis',
          axisPointer: { type: 'cross', lineStyle: { color: 'rgba(148, 163, 184, 0.4)' } },
          formatter: (params) => this.formatTooltip(params),
        },
        legend: showLegend ? {
          textStyle: { color: '#94a3b8', fontSize: 10 },
          bottom: 2,
          itemWidth: 8,
          itemHeight: 4
        } : undefined,
        xAxis: {
          type: 'category',
          data: xAxisData,
          boundaryGap: false,
          axisLine: { lineStyle: { color: '#2f3a45' } },
          axisTick: { show: false },
          axisLabel: { color: '#94a3b8', fontSize: 10, interval: xAxisInterval, hideOverlap: true },
        },
        yAxis,
        series,
      }
    },
    buildOIOptions() {
      const points = this.raw.oi_hist || []
      const xData = points.map(item => this.formatTimeShort(item.time))
      const oiData = points.map(item => item.open_interest)
      const oiValueData = points.map(item => item.open_interest_value)
      return {
        ...this.buildBaseOptions(
          xData,
          [
            this.buildScaledYAxis(value => this.formatCompact(value)),
            this.buildScaledYAxis(value => this.formatCompact(value))
          ],
          [
            {
              name: '持仓量',
              type: 'bar',
              data: oiData,
              itemStyle: { color: '#f5c84b' },
              barMaxWidth: 6,
            },
            {
              name: '持仓价值(USDT)',
              type: 'line',
              yAxisIndex: 1,
              smooth: true,
              data: oiValueData,
              itemStyle: { color: '#cbd5e1' },
              lineStyle: { width: 2 },
              symbol: 'none'
            }
          ],
          this.calcAxisInterval(xData.length)
        )
      }
    },
    buildRatioOptions(points, name) {
      const xData = (points || []).map(item => this.formatTimeShort(item.time))
      const yData = (points || []).map(item => item.ratio)
      return this.buildBaseOptions(
        xData,
        this.buildScaledYAxis(value => this.formatDecimal(value, 2)),
        [
          {
            name,
            type: 'line',
            smooth: true,
            showSymbol: false,
            data: yData,
            itemStyle: { color: '#f5c84b' },
            lineStyle: { width: 1.4 },
            symbol: 'none'
          }
        ],
        this.calcAxisInterval(xData.length),
        false
      )
    },
    buildTakerOptions() {
      const points = this.raw.taker_volume || []
      const xData = points.map(item => this.formatTimeShort(item.time))
      const buyData = points.map(item => item.buy_vol)
      const sellData = points.map(item => item.sell_vol)
      return this.buildBaseOptions(
        xData,
        this.buildVolumeYAxis(),
        [
          {
            name: '主动买入量',
            type: 'bar',
            data: buyData,
            itemStyle: { color: '#22c55e' },
            barGap: '20%',
            barCategoryGap: '40%',
          },
          {
            name: '主动卖出量',
            type: 'bar',
            data: sellData,
            itemStyle: { color: '#ef4444' },
            barGap: '20%',
            barCategoryGap: '40%',
          }
        ],
        this.calcAxisInterval(xData.length),
        true
      )
    },
    buildBasisOptions() {
      const points = this.raw.basis || []
      const xData = points.map(item => this.formatTimeShort(item.time))
      const basisData = points.map(item => item.basis)
      const markData = points.map(item => item.mark_price)
      const spotData = points.map(item => item.spot_price)
      return {
        ...this.buildBaseOptions(
          xData,
          [
            this.buildScaledYAxis(value => this.formatDecimal(value, 2)),
            this.buildScaledYAxis(value => this.formatDecimal(value, 2))
          ],
          [
            {
              name: '基差',
              type: 'line',
              smooth: true,
              showSymbol: false,
              data: basisData,
              itemStyle: { color: '#f5c84b' },
              lineStyle: { width: 1.4 },
              symbol: 'none',
              areaStyle: { color: 'rgba(245, 200, 75, 0.14)' }
            },
            {
              name: '标记价格',
              type: 'line',
              smooth: true,
              showSymbol: false,
              yAxisIndex: 1,
              data: markData,
              itemStyle: { color: '#94a3b8' },
              lineStyle: { width: 1 },
              symbol: 'none'
            },
            {
              name: '现货价格',
              type: 'line',
              smooth: true,
              showSymbol: false,
              yAxisIndex: 1,
              data: spotData,
              itemStyle: { color: '#38bdf8' },
              lineStyle: { width: 1 },
              symbol: 'none'
            }
          ],
          this.calcAxisInterval(xData.length),
          true
        )
      }
    },
    buildFundingOptions() {
      const points = this.raw.funding_rate || []
      const xData = points.map(item => parseTime(item.time, '{m}-{d}'))
      const yData = points.map(item => item.rate * 100)
      return this.buildBaseOptions(
        xData,
        this.buildScaledYAxis(value => `${this.formatDecimal(value, 3)}%`),
        [
          {
            name: '资金费率',
            type: 'line',
            smooth: true,
            showSymbol: false,
            data: yData,
            itemStyle: { color: '#f5c84b' },
            lineStyle: { width: 1.4 },
            symbol: 'none'
          }
        ],
        this.calcAxisInterval(xData.length),
        false
      )
    }
  }
}
</script>

<style scoped>
::v-deep .insight-drawer .el-drawer__body {
  background: #0b1220;
  overflow-y: auto;
}
::v-deep .insight-drawer {
  background: #0b1220;
  color: #cbd5e1;
}
::v-deep .el-drawer__wrapper {
  background: rgba(2, 6, 23, 0.65);
}
::v-deep .insight-drawer__wrapper {
  background: rgba(2, 6, 23, 0.65);
}

.insight-wrapper {
  padding: 18px;
  min-height: 100%;
  background: #0b1220;
  height: 100%;
  overflow-y: auto;
}
.insight-toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 12px;
}
.insight-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 500;
  color: #cbd5e1;
}
.insight-symbol {
  font-size: 18px;
}
.insight-sub {
  font-size: 13px;
  color: #8b95a5;
}
.insight-controls {
  display: flex;
  align-items: center;
  gap: 8px;
}
.insight-select ::v-deep .el-input__inner {
  background: #111827;
  border-color: #1f2937;
  color: #cbd5e1;
}
.insight-grid {
  margin-top: 8px;
}
.insight-card {
  background: #0e1623;
  border: 1px solid rgba(148, 163, 184, 0.12);
  box-shadow: inset 0 0 0 1px rgba(15, 23, 42, 0.2);
  border-radius: 10px;
  padding: 6px;
  margin-bottom: 12px;
  min-height: 380px;
  display: flex;
  flex-direction: column;
}
.insight-card__title {
  font-size: 13px;
  color: #cbd5e1;
  margin-bottom: 4px;
}
.insight-empty {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #8b95a5;
  font-size: 13px;
}
</style>
