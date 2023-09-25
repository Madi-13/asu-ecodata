import api from './api'

const URLS = {
  referenceRecordsURL: '/references/records'
}

export const getReferenceRecords = (payload, configs = {}) => {
  return api.post(URLS.referenceRecordsURL, payload)
}

export const getReferenceRecordsTree = (payload, configs = {}) => {
  return api.post(`${URLS.referenceRecordsURL}/tree`, payload)
}

export const getReferenceFormForEdit = (payload, configs = {}) => {
  return api.post(`${URLS.referenceRecordsURL}/for_edit`, payload)
}

export const getReferenceFormForCreate = (payload, configs = {}) => {
  return api.post(`${URLS.referenceRecordsURL}/for_create`, payload)
}

export const getReferenceRecordsByName = (payload, configs = {}) => {
  return api.post(`${URLS.referenceRecordsURL}/popup`, payload, configs)
}

export const getReferenceRecordHistory = (payload, configs = {}) =>
  api.post(`${URLS.referenceRecordsURL}/history`, payload, configs)

export const createNewRecord = (payload, configs = {}) => {
  return api.post(`${URLS.referenceRecordsURL}/create`, payload, configs)
}

export const updateRecord = (payload, configs = {}) => {
  return api.post(`${URLS.referenceRecordsURL}/update`, payload, configs)
}

export const deleteRecord = (payload, configs = {}) => {
  return api.post(`${URLS.referenceRecordsURL}/delete`, payload, configs)
}

export const setStatusRef = (payload) => {
  return api.post(`${URLS.referenceRecordsURL}/set_status`, payload)
}
