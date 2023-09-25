<template>
  <div>
    <RecordHistory :event="event" v-for="event in history" />
  </div>
</template>

<script>
import RecordHistory from './components/RecordHistory.vue'

import { getReferenceRecordHistory } from '@/api/referenceRecordsAPI'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

export default {
  components: {
    RecordHistory
  },
  data() {
    return {
      history: null,
      fetchStatus: IDLE,
      searchParams: {
        dictionary_id: null,
        row_id: null,
        date_from: null,
        date_to: null
      }
    }
  },
  computed: {
    ...apiStatusComputedFactory(['fetchStatus'])
  },
  created() {
    this.fetchList()
  },
  watch: {
    $route: {
      handler(newValue) {
        this.searchParams = {
          dictionary_id: newValue.query && +newValue.query.reference_id,
          row_id: newValue.params && +newValue.params.row_id
        }
      },
      immediate: true
    }
  },
  methods: {
    async fetchList() {
      this.fetchStatus = PENDING
      const { response, error } = await withAsync(
        getReferenceRecordHistory,
        this.searchParams
      )
      if (error) {
        this.$toast.error(error.message)
        this.fetchStatus = ERROR
        return
      }
      this.fetchStatus = SUCCESS
      this.history = response.data.result
    }
  }
}
</script>

<style lang="scss" scoped></style>
