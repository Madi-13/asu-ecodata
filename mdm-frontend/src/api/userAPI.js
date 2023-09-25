import api from './api'

const URLS = {
  usersListURL: '/api/v1/user/get-list',
  userByIdURL: '/api/v1/user/get-list',
  createUserURL: '/api/v1/user/create',
  updateUserURL: '/api/v1/user/update',
  deleteUserURL: '/api/v1/user/delete'
}

export const createUser = (payload) => api.post(URLS.createUserURL, payload)

export const updateUser = (payload) => {
  return api.post(URLS.updateUserURL, payload)
}

export const deleteUser = (configs = {}) => {
  return api.delete(URLS.deleteUserURL, null, configs)
}

export const getUsers = (configs = {}) => api.get(URLS.usersListURL, configs)

export const getUserById = (configs = {}) => api.get(URLS.userByIdURL, configs)
