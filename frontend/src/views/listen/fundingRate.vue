<template>
  <div class="app-container">
    <div style="margin-bottom: 10px; display: flex;justify-content: space-between;align-items: center;">
      <div style="display: flex;flex-flow: row wrap;gap: 10px; width:75%">
        <symbol-suggest-input
          v-model="search.symbol"
          :placeholder="$t('trade.coin') "
          style="width: 150px;"
          class="filter-item"
          size="small"
        />
        <el-button
          type="success"
          size="mini"
          @click="fetchData"
        >
          {{ $t('table.search') }}
        </el-button>
      </div>
      <div style="width:25%;text-align:right;">
        {{ $t('table.totalCount') }}: {{ list.length }}
      </div>
    </div>
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      size="mini"
      :row-key="rowKey"
      :expand-row-keys="expandKeys"
      highlight-current-row
      @sort-change="sortChange"
      @expand-change="expandChange"
    >
      <!-- <el-table-column align="center" label="ID" width="65">
        <template slot-scope="scope">
          {{ scope.$index + 1 }}
        </template>
      </el-table-column> -->

      <el-table-column
        :label="$t('trade.coin')"
        align="center"
        show-overflow-tooltip
      >
        <template slot-scope="scope">
          {{ scope.row.symbol }}
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('trade.nowFundRate')"
        align="center"
        show-overflow-tooltip
        sortable="custom"
      >
        <template slot-scope="scope">
          <span :style="{color: Math.abs(round(scope.row.now_funding_rate * 100, 4)) > 1 ? 'red' : 'green' }">
            {{ round(scope.row.now_funding_rate * 100, 4) }}
          </span>
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('trade.time')"
        align="center"
        show-overflow-tooltip
      >
        <template slot-scope="scope">
          {{ parseTime(scope.row.now_funding_time) }}
        </template>
      </el-table-column>
      <el-table-column
        label="下次结算时间"
        align="center"
        show-overflow-tooltip
      >
        <template slot-scope="scope">
          {{ scope.row.next_funding_time ? parseTime(scope.row.next_funding_time) : '-' }}
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('trade.nowPrice')"
        align="center"
        show-overflow-tooltip
      >
        <template slot-scope="scope">
          {{ scope.row.now_price }}
        </template>
      </el-table-column>
      <!-- <el-table-column
        :label="$t('trade.lastNotice')"
        align="center"
        width="200"
      >
        <template slot-scope="scope">
          {{ scope.row.last_notice_funding_time !== 0 ? parseTime(scope.row.last_notice_time) : '' }}
        </template>
      </el-table-column> -->
      <el-table-column :label="$t('trade.autoTrade')" align="center" width="100">
        <template slot-scope="scope">
          <el-select v-model="scope.row.auto_order" size="small" @change="edit(scope.row)">
            <el-option :label="$t('trade.yes')" :value="1" />
            <el-option :label="$t('trade.no')" :value="0" />
          </el-select>
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('table.actions')"
        align="center"
        width="80"
        class-name="small-padding fixed-width"
      >
        <template slot-scope="{row}">
          <el-button
            type="primary"
            size="mini"
            @click="getHistory(row)"
          >{{ $t('trade.history') }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog :title="dialogTitle" :visible.sync="dialogHistoryVisible" width="80%">
      <div style="height: 520px;">
        <apexchart type="line" height="520" :options="historyOptions" :series="historySeries" />
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getFundingRates, getFundingRateHistory, editFundingRates } from '@/api/listenCoin'
import SymbolSuggestInput from '@/components/SymbolSuggestInput'
import VueApexCharts from 'vue-apexcharts'
import { parseTime } from '@/utils'
import { round } from 'mathjs'

export default {
  components: {
    SymbolSuggestInput,
    apexchart: VueApexCharts,
  },
  data() {
    return {
      list: [],
      history: [],
      historyMeta: [],
      historySeries: [
        { name: this.$t('trade.fundRate'), data: [] },
      ],
      historyOptions: {
        chart: {
          type: 'line',
          toolbar: { show: true },
          zoom: { enabled: true },
          foreColor: '#B3C2D4',
        },
        stroke: {
          width: 2,
          curve: 'smooth',
        },
        grid: {
          borderColor: 'rgba(149, 176, 203, 0.12)',
        },
        xaxis: {
          type: 'datetime',
          labels: {
            datetimeUTC: false,
            style: { colors: '#8EA2B8' },
          },
        },
        yaxis: {
          labels: {
            formatter: (val) => `${Number(val).toFixed(4)}%`,
            style: { colors: '#8EA2B8' },
          },
        },
        tooltip: {
          custom: null,
        },
      },
      sort: '+',
      listLoading: false,
      search: {
        symbol: '',
      },
      rowKey(row) {
        return row.symbol
      },
      expandKeys: [],
      dialogTitle: '资金费率历史',
      dialogHistoryVisible: false,
    }
  },
  async created() {
    await this.fetchData()
    this.timeId = setInterval(() => this.fetchData(), 30 * 1000)
    this.historyOptions.tooltip.custom = this.buildHistoryTooltip
  },
  beforeDestroy() {
    clearInterval(this.timeId)
  },
  methods: {
    parseTime,
    round(data, num = 2) {
      return round(data, num)
    },
    expandChange(row, expandedRows) {
      this.expandKeys = expandedRows.map(item => item.symbol)
    },
    sortChange(data) {
      const { order } = data
      this.sort = order === 'ascending' ? '+' : '-'
      this.fetchData()
    },
    async fetchData() {
      // this.listLoading = true
      const search = this.search
      const { data } = await getFundingRates({ sort: this.sort, ...search })
      this.list = data
      // this.listLoading = false
    },
    async getHistory(row) {
      this.dialogTitle = `${row.symbol} ` + this.$t('trade.fundRateHistory')
      this.dialogHistoryVisible = true
      const { data } = await getFundingRateHistory({
        symbol: row.symbol,
      })
      const history = data.reverse() // 数据从最新到最旧
      this.history = history
      this.historyMeta = history
      this.historySeries = [
        {
          name: this.$t('trade.fundRate'),
          data: history.map(item => ({
            x: item.fundingTime,
            y: round(item.fundingRate * 100, 4),
          })),
        },
      ]
    },
    buildHistoryTooltip({ dataPointIndex }) {
      const item = this.historyMeta[dataPointIndex]
      if (!item) {
        return ''
      }
      const rate = round(item.fundingRate * 100, 4)
      const time = this.parseTime(item.fundingTime)
      const price = item.markPrice
      return `
        <div style="padding:8px 10px;background:#0F1720;border:1px solid rgba(149,176,203,0.18);border-radius:8px;">
          <div style="font-size:12px;color:#8EA2B8;margin-bottom:4px;">${time}</div>
          <div style="font-size:14px;color:#E6EEF7;">资金费率：${rate}%</div>
          <div style="font-size:12px;color:#B3C2D4;margin-top:2px;">标记价格：${price}</div>
        </div>
      `
    },
    async edit(row) {
      const { id, enable, ...data } = row
      try {
        await editFundingRates(id, { ...data, enable: enable ? 1 : 0 })
        this.$message({ message: this.$t('table.actionSuccess'), type: 'success' })
        await this.fetchData()
      } catch (e) {
        this.$message({ message: this.$t('table.actionFail'), type: 'success' })
      }
    },
  },
}
</script>
