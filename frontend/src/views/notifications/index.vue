<template>
  <div class="app-container">
    <div style="margin-bottom: 10px">
      <el-input
        v-model="listQuery.keyword"
        size="mini"
        class="filter-item"
        style="width: 200px; margin-right: 8px;"
        placeholder="关键字"
      />
      <el-select
        v-model="listQuery.module"
        size="mini"
        clearable
        class="filter-item"
        style="width: 160px; margin-right: 8px;"
        placeholder="模块"
      >
        <el-option
          v-for="item in moduleOptions"
          :key="item.module"
          :label="`${item.module} (${item.count})`"
          :value="item.module"
        />
      </el-select>
      <el-select
        v-model="listQuery.channel"
        size="mini"
        clearable
        class="filter-item"
        style="width: 120px; margin-right: 8px;"
        placeholder="渠道"
      >
        <el-option label="dingding" value="dingding" />
        <el-option label="slack" value="slack" />
      </el-select>
      <el-select
        v-model="listQuery.level"
        size="mini"
        clearable
        class="filter-item"
        style="width: 110px; margin-right: 8px;"
        placeholder="等级"
      >
        <el-option label="info" value="info" />
        <el-option label="warning" value="warning" />
        <el-option label="error" value="error" />
      </el-select>
      <el-button
        size="mini"
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        :loading="listLoading"
        @click="handleFilter"
      >{{ $t('table.search') }}</el-button>
      <el-button
        type="danger"
        size="mini"
        :loading="listLoading"
        @click="clearAll"
      >
        {{ $t('table.deleteAll') }}
      </el-button>
    </div>

    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      size="mini"
      highlight-current-row
    >
      <el-table-column label="时间" align="center" width="160">
        <template slot-scope="scope">
          {{ parseTime(scope.row.createTime) }}
        </template>
      </el-table-column>
      <el-table-column label="模块" align="center" width="120" show-overflow-tooltip>
        <template slot-scope="scope">
          {{ scope.row.module }}
        </template>
      </el-table-column>
      <el-table-column label="等级" align="center" width="90">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.level === 'error'" type="danger" size="mini">error</el-tag>
          <el-tag v-else-if="scope.row.level === 'warning'" type="warning" size="mini">warning</el-tag>
          <el-tag v-else type="success" size="mini">info</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="渠道" align="center" width="100">
        <template slot-scope="scope">
          {{ scope.row.channel }}
        </template>
      </el-table-column>
      <el-table-column label="标题" align="center" show-overflow-tooltip>
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column label="内容" align="center" width="90">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="openContent(scope.row)">查看</el-button>
        </template>
      </el-table-column>
      <el-table-column :label="$t('table.actions')" align="center" width="90">
        <template slot-scope="scope">
          <el-button type="danger" size="mini" @click="deleteRow(scope.row)">{{ $t('table.delete') }}</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination
      v-show="total > 0"
      :total="total"
      :page.sync="listQuery.page"
      :limit.sync="listQuery.limit"
      @pagination="getList"
    />

    <el-dialog :title="dialogTitle" :visible.sync="dialogVisible" width="60%">
      <pre class="notify-content">{{ dialogContent }}</pre>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">{{ $t('table.cancel') }}</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getNotifications, deleteNotification, clearNotifications, getNotificationModules } from '@/api/notification'
import Pagination from '@/components/Pagination'
import { parseTime } from '@/utils'

export default {
  components: { Pagination },
  data() {
    return {
      list: [],
      total: 0,
      listLoading: false,
      listQuery: {
        page: 1,
        limit: 50,
        module: '',
        channel: '',
        level: '',
        keyword: '',
      },
      moduleOptions: [],
      dialogVisible: false,
      dialogTitle: '',
      dialogContent: '',
    }
  },
  created() {
    this.getList()
    this.getModules()
  },
  methods: {
    parseTime,
    async getList() {
      this.listLoading = true
      try {
        const { data: { list, total }} = await getNotifications(this.listQuery)
        this.list = list
        this.total = total
      } finally {
        this.listLoading = false
      }
    },
    async getModules() {
      const { data } = await getNotificationModules()
      this.moduleOptions = data || []
    },
    handleFilter() {
      this.listQuery.page = 1
      this.getList()
    },
    openContent(row) {
      this.dialogTitle = row.title || '通知内容'
      this.dialogContent = row.content || ''
      this.dialogVisible = true
    },
    deleteRow(row) {
      this.$confirm(`${this.$t('table.deleteConfirm')} ${row.title || ''}?`)
        .then(async() => {
          await deleteNotification(row.id)
          this.$message({ message: this.$t('table.deleteSuccess'), type: 'success' })
          this.getList()
        })
        .catch(() => {})
    },
    clearAll() {
      this.$confirm(`${this.$t('table.deleteAll')}?`)
        .then(async() => {
          await clearNotifications()
          this.$message({ message: this.$t('table.deleteSuccess'), type: 'success' })
          this.getList()
        })
        .catch(() => {})
    },
  },
}
</script>

<style scoped>
.notify-content {
  white-space: pre-wrap;
  word-break: break-word;
  background: #0f172a;
  color: #e2e8f0;
  padding: 12px;
  border-radius: 8px;
  max-height: 60vh;
  overflow: auto;
}
</style>
