import api from './api'

const URLS = {
  getScenarioURL: '/references/export/scenario',
  integrationRouteURL: '/references/export/property',
  logsURL: '/references/export/journal',
  processURL: '/references/export/process'
}

export const getScenarios = (payload, configs = {}) =>
  api.post(URLS.getScenarioURL, payload, configs)

export const operateScenario = (payload, configs = {}) =>
  api.post(`${URLS.getScenarioURL}/dml`, payload, configs)

export const getIntegrationRoutes = (payload, configs = {}) =>
  api.post(URLS.integrationRouteURL, payload, configs)

export const operateIntegrationRoute = (payload, configs = {}) =>
  api.post(`${URLS.integrationRouteURL}/dml`, payload, configs)

export const getIntegrationLogs = (payload, configs = {}) =>
  api.post(URLS.logsURL, payload, configs)

export const getProcesses = (payload, configs = {}) =>
  api.post(URLS.processURL, payload, configs)

export const operateProcess = (payload, configs = {}) =>
  api.post(`${URLS.processURL}/dml`, payload, configs)
