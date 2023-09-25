import api from './api'

const URLS = {
  importUrl: '/references/import',
  importTemplateUrl: '/references/import/tamplate',
  exportUrl: '/references/export/values',
  monitoringUrl: '/references/import/monitoring'
}

export const uploadFile = (payload) => api.post(URLS.importUrl, payload)

export const downloadFileTemplate = (payload) =>
  api.post(URLS.importTemplateUrl, payload)

export const downloadFile = (payload) => api.post(URLS.exportUrl, payload)

export const getMonitoring = (payload) => api.post(URLS.monitoringUrl, payload)

export const cancelMonitoring = (payload) =>
  api.post(`${URLS.monitoringUrl}/ends`, payload)
