import Vue from 'vue'
import dayjs from 'dayjs'
var utc = require('dayjs/plugin/utc')

dayjs.extend(utc)

Vue.filter('fullDateTime', (value) => {
  if (!value) return 'Не установлен'
  return dayjs(value).utc().format('DD.MM.YYYY, в HH:mm')
})

Vue.filter('normalizeDate', (value) => {
  if (!value) return 'Не установлен'
  return dayjs(value).format('DD.MM.YYYY')
})
