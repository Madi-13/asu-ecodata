<template>
  <div class="analytic" v-if="info">
    <div class="analytic__block">
      <div class="analytic__row">
        <InfoCard>
          <p class="info-card__score">{{ info.dictionaries_count || 0 }}</p>
          <p class="info-card__descr text--center">
            Количество справочниов/классификатров
          </p>
        </InfoCard>
        <InfoCard>
          <p class="info-card__score">{{ info.rows_count || 0 }}</p>
          <p class="info-card__descr text--center">
            Количество записей в справочнике
          </p>
        </InfoCard>
      </div>
      <DoughnutInfo
        :data="info.rows_by_status"
        class="mt--1"
        v-if="filter.table"
      />
      <BaseDropdown
        label="Фильтр по дате"
        placeholder="Выберите из списка"
        v-model="filter.date"
        optionValueField="id"
        optionTitleField="filter_name"
        :options="filters"
        v-if="filters"
        class="mt--2"
        :width="100"
      />
      <BarByDateInfo
        :data="infoByDate && infoByDate.rows_by_date"
        class="mt--1"
        v-if="filter.table && infoByDate"
      />
    </div>
    <div class="analytic__block">
      <InputAutocomplete
        name="refe"
        placeholder="Введите значение"
        label="Справочник"
        :searchMethod="getSystemRefs"
        optionTitleField="table_name"
        optionValueField="id"
        v-model="filter.table"
      />
      <BarUserInfo
        :data="info.rows_by_user"
        class="mt--1"
        v-if="filter.table"
      />
    </div>
  </div>
</template>

<script>
import InputAutocomplete from '@/components/common/InputAutocomplete.vue'
import InfoCard from './components/InfoCard.vue'
import DoughnutInfo from './components/DoughnutInfo.vue'
import BarUserInfo from './components/BarUserInfo.vue'
import BarByDateInfo from './components/BarByDateInfo.vue'

import { getSystemRefs } from '@/api/objectsAPI'
import { getMainPanelInfo, getFilters, getDetailInfo } from '@/api/analyticsAPI'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

export default {
  components: {
    InfoCard,
    DoughnutInfo,
    BarUserInfo,
    BarByDateInfo,
    InputAutocomplete
  },
  data() {
    return {
      info: null,
      filters: null,
      infoByDate: null,
      fetchInfoStatus: IDLE,
      fetchFilterStatus: IDLE,
      fetchByDateStatus: IDLE,
      filter: {
        table: null,
        date: 2
      }
    }
  },
  watch: {
    'filter.table': {
      async handler() {
        this.getAnalyticsInfo()
        await this.getInfoFilters()
        await this.getInfoByDate()
      }
    },
    'filter.date': {
      async handler() {
        await this.getInfoByDate()
      }
    }
  },
  computed: {
    ...apiStatusComputedFactory([
      'fetchStatus',
      'fetchFilterStatus',
      'fetchByDateStatus'
    ])
  },
  created() {
    this.getAnalyticsInfo()
  },
  methods: {
    getSystemRefs,
    async getAnalyticsInfo() {
      this.fetchInfoStatus = PENDING
      const { response, error } = await withAsync(getMainPanelInfo, {
        dictionary_id: this.filter?.table?.id || 0,
        filter_id: this.filter?.date?.id || this.filter?.date
      })
      if (error) {
        this.fetchInfoStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.fetchInfoStatus = SUCCESS
      this.info = response.data
    },
    async getInfoFilters() {
      this.fetchFilterStatus = PENDING
      const { response, error } = await withAsync(getFilters, {
        table_id: this.filter?.table?.id
      })
      if (error) {
        this.fetchFilterStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.fetchFilterStatus = SUCCESS
      this.filters = response.data
    },
    async getInfoByDate() {
      this.fetchByDateStatus = PENDING
      const { response, error } = await withAsync(getDetailInfo, {
        dictionary_id: this.filter?.table?.id || 0,
        filter_id: this.filter?.date?.id || this.filter?.date
      })
      if (error) {
        this.fetchByDateStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.fetchByDateStatus = SUCCESS
      this.infoByDate = response.data
    }
  }
}
</script>

<style lang="scss" scoped>
.analytic {
  display: grid;
  grid-template-columns: repeat(2, 50%);
  gap: 30px;

  &__row {
    display: flex;
    justify-content: space-between;
    gap: 15px;
  }
}
</style>
