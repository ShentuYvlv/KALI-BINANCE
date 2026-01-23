<template>
  <el-autocomplete
    v-model="innerValue"
    :fetch-suggestions="querySearch"
    :placeholder="placeholder"
    :debounce="200"
    :trigger-on-focus="triggerOnFocus"
    :clearable="clearable"
    :disabled="disabled"
    :size="size"
    :popper-class="popperClass"
    @input="handleInput"
  />
</template>

<script>
import { getFeaturesOptions } from '@/api/trade'

let cachedSymbols = null
let cachedPromise = null

async function loadSymbols() {
  if (Array.isArray(cachedSymbols)) {
    return cachedSymbols
  }
  if (cachedPromise) {
    return cachedPromise
  }
  cachedPromise = getFeaturesOptions()
    .then(({ data }) => {
      cachedSymbols = Array.isArray(data) ? data : []
      cachedPromise = null
      return cachedSymbols
    })
    .catch(() => {
      cachedPromise = null
      return []
    })
  return cachedPromise
}

export default {
  name: 'SymbolSuggestInput',
  props: {
    value: {
      type: String,
      default: ''
    },
    placeholder: {
      type: String,
      default: ''
    },
    size: {
      type: String,
      default: 'mini'
    },
    disabled: {
      type: Boolean,
      default: false
    },
    uppercase: {
      type: Boolean,
      default: true
    },
    clearable: {
      type: Boolean,
      default: true
    },
    triggerOnFocus: {
      type: Boolean,
      default: true
    },
    suggestLimit: {
      type: Number,
      default: 50
    },
    popperClass: {
      type: String,
      default: 'symbol-suggest-popper'
    }
  },
  data() {
    return {
      innerValue: this.value,
      symbols: []
    }
  },
  watch: {
    value(val) {
      this.innerValue = val
    }
  },
  created() {
    this.ensureSymbols()
  },
  methods: {
    async ensureSymbols() {
      this.symbols = await loadSymbols()
    },
    querySearch(queryString, cb) {
      const raw = queryString || ''
      const query = this.uppercase ? raw.toUpperCase() : raw
      const list = this.symbols.length ? this.symbols : []
      let result = list
      if (query) {
        result = list.filter((item) => item.startsWith(query))
      }
      const mapped = result.slice(0, this.suggestLimit).map(item => ({ value: item }))
      cb(mapped)
    },
    handleInput(val) {
      const nextVal = this.uppercase ? (val || '').toUpperCase() : val
      this.innerValue = nextVal
      this.$emit('input', nextVal)
    }
  }
}
</script>
