import api from './api'

const URLS = {
  menuUrl: '/references/objects/tables_menu'
}

export const getMenuContent = (payload = {}, configs = {}) =>
  api.post(URLS.menuUrl, payload)
