<template>
  <InfoCard>
    <Bar
      :options="chartOptions"
      :data="chartData"
      :chart-id="chartId"
      :dataset-id-key="datasetIdKey"
      v-if="chartData"
    />
    <p v-else>Нет данных</p>
  </InfoCard>
</template>

<script>
import { Bar } from 'vue-chartjs'
import InfoCard from './InfoCard.vue'

import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale
} from 'chart.js'

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)
export default {
  name: 'BarChart',
  components: {
    Bar,
    InfoCard
  },
  props: {
    chartId: {
      type: String,
      default: 'bar-chart'
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
        maintainAspectRatio: false,
        indexAxis: 'y'
      }
    }
  },
  watch: {
    data(newValue) {
      if (!newValue) {
        this.chartData = null
        return
      }
      this.chartData = {
        labels: newValue.map((el) => el.user_name || 'Неизвестный'),
        datasets: [
          {
            label: 'Количество по авторам',
            backgroundColor: '#41B883',
            data: newValue.map((el) => el.rows_count)
          }
        ]
      }
    }
  }
}
</script>

<style lang="scss" scoped></style>
