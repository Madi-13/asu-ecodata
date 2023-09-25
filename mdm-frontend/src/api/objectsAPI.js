import api from './api'

const URLS = {
  columnTypesUrl: '/references/objects/types',
  systemRefsUrl: '/references/objects/tables',
  systemRefUrl: '/references/objects/table',
  systemRefEditUrl: '/references/objects/table/alter',
  systemRefDeleteUrl: '/references/objects/table/drop',
  systemRefsDelUrl: '/references/objects/table/drop'
}

export const getColumnTypes = (config = {}) => {
  return api.get(URLS.columnTypesUrl, config)
}

export const getSystemRefs = (payload, configs = {}) => {
  return api.post(URLS.systemRefsUrl, payload, configs)
}

export const getSystemRefById = (table_id, configs = {}) => {
  return api.post(
    URLS.systemRefsUrl,
    {
      limit: 1,
      offset: 1,
      table_id: +table_id
    },
    configs
  )
}

export const createSystemRef = (payload, configs = {}) => {
  return api.post(URLS.systemRefUrl, payload, configs)
}

export const editSystemRef = (payload, configs = {}) => {
  return api.post(URLS.systemRefEditUrl, payload, configs)
}

export const deleteSystemRef = (payload) => {
  return api.post(URLS.systemRefsDelUrl, payload)
}

export const uploadFiles = (payload) => {
  return api.post(`${URLS.systemRefsUrl}/file_upload`, payload)
}

export const getFiles = (payload) => {
  return api.post(`${URLS.systemRefsUrl}/file_list`, payload)
}

export const removeFile = (payload) => {
  return api.post(`${URLS.systemRefsUrl}/delete_file`, payload)
}
