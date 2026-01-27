<template>
  <div class="insight-chart" :style="{ height }" />
</template>

<script>
import echarts from 'echarts'
import resize from '@/views/dashboard/admin/components/mixins/resize'

export default {
  name: 'InsightChart',
  mixins: [resize],
  props: {
    options: {
      type: Object,
      default: () => ({})
    },
    height: {
      type: String,
      default: '180px'
    }
  },
  data() {
    return {
      chart: null
    }
  },
  watch: {
    options: {
      deep: true,
      handler(val) {
        if (!this.chart) {
          return
        }
        this.chart.setOption(val, true)
      }
    }
  },
  mounted() {
    this.$nextTick(() => {
      this.chart = echarts.init(this.$el)
      this.chart.setOption(this.options, true)
    })
  },
  beforeDestroy() {
    if (this.chart) {
      this.chart.dispose()
      this.chart = null
    }
  }
}
</script>

<style scoped>
.insight-chart {
  width: 100%;
  flex: 1;
}
</style>
