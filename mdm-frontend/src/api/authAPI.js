import api from './api'
import axios from 'axios'

const localAxiosInstance = axios.create()

const URLS = {
  loginUrl: '/api/v1/auth/login/',
  refreshTokenUrl: '/api/v1/auth/refresh/',
  checkTokenValidityUrl: '/api/v1/auth/validate/',
  permissionURL: '/api/v1/auth/permission'
}

export const login = (payload, config = {}) => {
  return api.post(URLS.loginUrl, payload, config)
}

export const refreshToken = (payload, config = {}) => {
  return api.post(URLS.refreshTokenUrl, payload, config)
}

export const checkTokenValidity = (payload, config = {}) => {
  return api.post(URLS.checkTokenValidityUrl, payload, config)
}

export const getPermissions = (configs = {}) => {
  return api.get(URLS.permissionURL, configs)
}
