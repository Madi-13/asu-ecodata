<template>
  <InfoCard :style="{ height: '300px' }">
    <template v-slot:title>
      <p :style="{ width: '100%' }">Разбивка по статусам</p>
    </template>

    <Doughnut
      v-if="chartData"
      :data="chartData"
      :chart-id="chartId"
      :dataset-id-key="datasetIdKey"
      :options="chartOptions"
    />
    <p v-else>Нет данных</p>
  </InfoCard>
</template>

<script>
import InfoCard from './InfoCard.vue'
import { Doughnut } from 'vue-chartjs'

import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  ArcElement,
  CategoryScale
} from 'chart.js'

ChartJS.register(Title, Tooltip, Legend, ArcElement, CategoryScale)

export default {
  components: {
    Doughnut,
    InfoCard
  },
  props: {
    chartId: {
      type: String,
      default: 'doughnut-chart'
    },
    datasetIdKey: {
      type: String,
      default: 'label'
    },
    data: {
      type: Array
    }
  },
  data() {
    return {
      chartData: null,
      chartOptions: {
        responsive: true,
        maintainAspectRatio: false
      }
    }
  },
  watch: {
    data(newValue) {
      if (!newValue) return
      this.chartData = {
        labels: ['Активно', 'Блокированно'],
        datasets: [
          {
            backgroundColor: ['#41B883', '#E46651'],
            data: [newValue[0]?.rows_count || 0, newValue[1]?.rows_count || 0]
          }
        ]
      }
    }
  }
}
</script>

<style lang="scss" scoped></style>
