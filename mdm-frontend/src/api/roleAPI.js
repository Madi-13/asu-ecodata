import api from './api'

const URLS = {
  listRoleUrl: '/api/v1/role/get-client-roles',
  listUserRoleUrl: '/api/v1/role/get-user-roles',
  listClientRolesUrl: '/api/v1/role/get-clients',
  createRoleUrl: '/api/v1/role/add-user-roles',
  deleteRoleUrl: '/api/v1/role/delete-user-roles'
}

export const getRealmClient = (configs = {}) => {
  return api.get(URLS.listClientRolesUrl, configs)
}

export const createRole = (payload, configs = {}) => {
  return api.post(URLS.createRoleUrl, payload, configs)
}

export const getUserRoles = (payload, configs = {}) =>
  api.get(URLS.listUserRoleUrl, payload, configs)

export const deleteRole = (configs = {}) => {
  return api.delete(URLS.deleteRoleUrl, configs)
}

export const getRoles = (configs = {}) => api.get(URLS.listRoleUrl, configs)
