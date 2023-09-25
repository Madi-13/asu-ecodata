import getters from './authGetters'
import mutations from './authMutations'
import actions from './authActions'

import { apiStatus } from '@/api/constants/apiStatus'
const { IDLE } = apiStatus

const state = {
  token: {},
  username: null,
  isLoggedIn: false,
  realmClientId: null,
  realmClientStatus: IDLE,
  realClientIdErrorMessage: null,
  loginStatus: IDLE,
  loginError: null,
  loginErrorMessage: '',
  checkTokenStatus: IDLE,
  checkTokenError: null,
  checkTokenErrorMessage: '',
  refreshTokenStatus: IDLE,
  refreshTokenError: null,
  refreshTokenErrorMessage: '',
  permissions: null,
  permissionsStatus: IDLE,
  permissionErrorMessage: ''
}

export default {
  state,
  getters,
  actions,
  mutations
}
