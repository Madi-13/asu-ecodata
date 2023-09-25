import api from './api'

const URLS = {
  filterUrl: '/references/info/filter',
  mainPanelUrl: '/references/info/main',
  detailUrl: '/references/info/date'
}

export const getFilters = (payload, configs = {}) =>
  api.post(URLS.filterUrl, payload, configs)

export const getMainPanelInfo = (payload, configs = {}) =>
  api.post(URLS.mainPanelUrl, payload, configs)

export const getDetailInfo = (payload, configs = {}) =>
  api.post(URLS.detailUrl, payload, configs)
