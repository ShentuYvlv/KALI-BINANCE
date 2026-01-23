<template>
  <div class="app-container">
    <div style="margin-bottom: 10px">
      <el-button
        type="success"
        size="mini"
        @click="openDialog()"
      >
        {{ $t('table.add') }}
      </el-button>
      <el-button
        type="primary"
        size="mini"
        :loading="listLoading"
        @click="fetchData()"
      >
        {{ $t('table.refresh') }}
      </el-button>
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
        :label="$t('trade.listenType')"
        align="center"
        width="130"
      >
        <template slot-scope="scope">
          <el-select v-model="scope.row.listen_type" size="small" @change="edit(scope.row)">
            <el-option :label="$t('trade.klineBase')" value="kline_base" />
            <el-option label="价格通知" value="price_notice" />
          </el-select>
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('trade.klineInterval')"
        align="center"
        width="110"
      >
        <template slot-scope="scope">
          <el-select v-if="scope.row.listen_type !== 'price_notice'" v-model="scope.row.kline_interval" size="small" @change="edit(scope.row)">
            <el-option label="1m" value="1m" />
            <el-option label="3m" value="3m" />
            <el-option label="5m" value="5m" />
            <el-option label="15m" value="15m" />
            <el-option label="30m" value="30m" />
            <el-option label="1h" value="1h" />
            <el-option label="2h" value="2h" />
            <el-option label="4h" value="4h" />
            <el-option label="6h" value="6h" />
            <el-option label="8h" value="8h" />
            <el-option label="12h" value="12h" />
            <el-option label="1d" value="1d" />
            <el-option label="3d" value="3d" />
            <el-option label="1w" value="1w" />
            <el-option label="1M" value="1M" />
          </el-select>
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('trade.changePercent')"
        align="center"
        width="140"
      >
        <template slot-scope="scope">
          <el-input
            v-if="scope.row.listen_type !== 'price_notice'"
            v-model="scope.row.change_percent"
            class="edit-input"
            size="small"
            @blur="edit(scope.row)"
          />
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('trade.noticeLimitMin')"
        align="center"
        width="140"
      >
        <template slot-scope="scope">
          <el-input
            v-if="scope.row.listen_type !== 'price_notice'"
            v-model="scope.row.notice_limit_min"
            class="edit-input"
            size="small"
            @blur="edit(scope.row)"
          />
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('trade.positionSide')"
        align="center"
        width="110"
      >
        <template slot-scope="scope">
          <el-select v-if="scope.row.listen_type === 'price_notice'" v-model="scope.row.side" size="small" @change="edit(scope.row)">
            <el-option :label="$t('trade.buy')" value="buy" />
            <el-option :label="$t('trade.sell')" value="sell" />
          </el-select>
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('trade.noticePrice')"
        align="center"
        width="130"
      >
        <template slot-scope="scope">
          <el-input
            v-if="scope.row.listen_type === 'price_notice'"
            v-model="scope.row.notice_price"
            class="edit-input"
            size="small"
            @blur="edit(scope.row)"
          />
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('trade.hasNotice')"
        align="center"
        width="100"
      >
        <template slot-scope="scope">
          <span v-if="scope.row.listen_type === 'price_notice'">
            {{ scope.row.has_notice == '1' ? $t('trade.yes') : $t('trade.no') }}
          </span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('trade.autoTrade')" align="center" width="100">
        <template slot-scope="scope">
          <el-select v-if="scope.row.listen_type === 'price_notice'" v-model="scope.row.auto_order" size="small" @change="edit(scope.row)">
            <el-option :label="$t('trade.yes')" :value="1" />
            <el-option :label="$t('trade.no')" :value="0" />
          </el-select>
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('trade.usdt')"
        align="center"
        width="100"
      >
        <template slot-scope="scope">
          <el-input
            v-if="scope.row.listen_type === 'price_notice'"
            v-model="scope.row.usdt"
            class="edit-input"
            size="small"
            @blur="edit(scope.row)"
          />
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('trade.amount')"
        align="center"
        width="100"
      >
        <template slot-scope="scope">
          <el-input
            v-if="scope.row.listen_type === 'price_notice'"
            v-model="scope.row.quantity"
            class="edit-input"
            size="small"
            @blur="edit(scope.row)"
          />
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('trade.lastNotice')"
        align="center"
        width="200"
      >
        <template slot-scope="scope">
          {{ scope.row.last_notice_time !== 0 ? parseTime(scope.row.last_notice_time) : '' }}
          {{ typeText(scope.row.last_notice_type) }}
        </template>
      </el-table-column>
      <el-table-column :label="$t('trade.enable')" align="center" width="80">
        <template slot-scope="{ row }">
          <el-switch
            v-model="row.enable"
            active-color="#13ce66"
            inactive-color="#dcdfe6"
            @change="isChangeBuy($event, row)"
          />
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
            type="danger"
            size="mini"
            @click="del(row)"
          >{{ $t('table.delete') }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form
        ref="dataForm"
        :model="info"
        label-position="left"
        label-width="100px"
        style="width: 400px; margin-left:50px;"
      >
        <el-form-item :label="$t('trade.coin')" prop="symbol">
          <symbol-suggest-input v-model="info.symbol" />
        </el-form-item>
        <el-form-item :label="$t('trade.listenType')" prop="listen_type">
          <el-select v-model="info.listen_type" size="small">
            <el-option :label="$t('trade.klineBase')" value="kline_base" />
            <el-option label="价格通知" value="price_notice" />
          </el-select>
        </el-form-item>
        <el-form-item v-if="info.listen_type === 'price_notice'" :label="$t('trade.noticePrice')" prop="notice_price">
          <el-input v-model="info.notice_price" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">{{ $t('table.cancel') }}</el-button>
        <el-button type="primary" :loading="dialogLoading" @click="addCoin(info)">{{ $t('table.confirm') }}</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getListenCoins, setListenCoin, addListenCoin, delListenCoin } from '@/api/listenCoin'
import SymbolSuggestInput from '@/components/SymbolSuggestInput'
import { parseTime } from '@/utils'

export default {
  components: {
    SymbolSuggestInput,
  },
  data() {
    return {
      list: [],
      tickets: {},
      sort: '+',
      listLoading: false,
      enableLoading: false,
      dialogFormVisible: false,
      dialogLoading: false,
      dialogTitle: '新增币种信息',
      info: {},
      rowKey(row) {
        return row.symbol
      },
      expandKeys: []
    }
  },
  async created() {
    await this.fetchData()
    // this.timeId = setInterval(() => this.fetchData(), 5 * 1000)
  },
  beforeDestroy() {
    // clearInterval(this.timeId)
  },
  methods: {
    typeText(type) {
      return {
        'up': this.$t('trade.up'),
        'down': this.$t('trade.down')
      }[type] || ''
    },
    parseTime,
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
      const { data } = await getListenCoins({ sort: this.sort, type: 1 })
      this.list = data.map(item => ({ ...item, enable: item.enable > 0 }))
      // this.listLoading = false
    },
    async edit(row) {
      const { id, enable, notice_limit_min, ...other } = row
      try {
        const limit = row.listen_type === 'price_notice' ? 0 : Number(notice_limit_min)
        await setListenCoin(id, { notice_limit_min: limit, enable: enable ? 1 : 0, ...other })
        this.$message({ message: this.$t('table.actionSuccess'), type: 'success' })
        await this.fetchData()
      } catch (e) {
        this.$message({ message: this.$t('table.actionFail'), type: 'success' })
      }
    },
    del(row) {
      this.$confirm(this.$t('table.deleteConfirm') + ` ${row.symbol} ？`)
        .then(async() => {
          try {
            await delListenCoin(row.id)
            this.$message({ message: this.$t('table.deleteSuccess'), type: 'success' })
            await this.fetchData()
          } catch (e) {
            this.$message({ message: this.$t('table.deleteFail'), type: 'success' })
          }
        })
        .catch(() => {})
    },
    async isChangeBuy(event, row) {
      await this.edit(row)
    },
    openDialog() {
      this.dialogTitle = this.$t('table.add')
      this.info = {
        symbol: '',
        listen_type: 'kline_base',
        notice_price: '',
        side: 'buy',
        auto_order: 1,
        usdt: '10',
        quantity: '0',
      }
      this.dialogFormVisible = true
    },
    async addCoin(row) {
      const data = {
        ...row,
        'type': 1, // 现货交易
        'createTime': +new Date(),
        'updateTime': +new Date()
      }
      await addListenCoin(data)
      await this.fetchData()
      this.dialogFormVisible = false
    }
  }
}
</script>
