import { apiStatus } from '@/api/constants/apiStatus'
const { PENDING, SUCCESS, ERROR, IDLE } = apiStatus
import {
  setDataToLocalStorage,
  getDataFromLocalStorage,
  clearDataLocalStorage
} from '@/helpers/permanentStorageHelper'

export default {
  SET_USER_LOGIN_STATE(state, payload) {
    switch (payload.status) {
      case PENDING:
        state.loginError && (state.loginError = null)
        state.loginErrorMessage && (state.loginErrorMessage = null)
        break

      case SUCCESS:
        state.isLoggedIn = true
        state.token.access_token = payload.data.access_token
        state.token.refresh_token = payload.data.refresh_token
        setDataToLocalStorage('mdmAccessToken', payload.data.access_token)
        setDataToLocalStorage('mdmRefreshToken', payload.data.refresh_token)
        break

      case ERROR:
        state.loginError = payload.error
        state.loginErrorMessage = payload.errorMessage || ''
        state.isLoggedIn = false
        break
    }
    state.loginStatus = payload.status
  },
  SET_REFRESH_TOKEN_STATE(state, payload) {
    switch (payload.status) {
      case PENDING:
        state.refreshTokenError && (state.refreshTokenError = null)
        state.refreshTokenErrorMessage &&
          (state.refreshTokenErrorMessage = null)
        break

      case SUCCESS:
        state.token.access_token = payload.data.access_token
        state.token.refresh_token = payload.data.refresh_token
        setDataToLocalStorage('mdmAccessToken', payload.data.access_token)
        setDataToLocalStorage('mdmRefreshToken', payload.data.refresh_token)
        break

      case ERROR:
        state.refreshTokenError = payload.error
        state.refreshTokenErrorMessage = payload.errorMessage || ''
        break
    }
    state.refreshTokenStatus = payload.status
  },
  SET_CHECK_TOKEN_STATE(state, payload) {
    switch (payload.status) {
      case PENDING:
        state.checkTokenError && (state.checkTokenError = null)
        state.checkTokenErrorMessage && (state.checkTokenErrorMessage = null)
        break

      case SUCCESS:
        state.token.access_token = getDataFromLocalStorage('mdmAccessToken')
        state.token.refresh_token = getDataFromLocalStorage('mdmRefreshToken')
        state.username = payload.data.name || payload.data.preferred_username
        state.isLoggedIn = true
        break

      case ERROR:
        state.checkTokenError = payload.error
        state.checkTokenErrorMessage = payload.errorMessage || ''
        break
    }
    state.checkTokenStatus = payload.status
  },
  SET_PERMISSIONS(state, payload) {
    switch (payload.status) {
      case PENDING:
        state.permissionErrorMessage && (state.permissionErrorMessage = null)
        break

      case SUCCESS:
        state.permissions = payload.data
        break

      case ERROR:
        state.permissionErrorMessage = payload.errorMessage || ''
        break
    }
    state.permissionsStatus = payload.status
  },
  SET_LOGOUT(state) {
    state.token = {
      access_token: null,
      refresh_token: null
    }
    state.isLoggedIn = false
    state.loginStatus = IDLE
    state.checkTokenStatus = IDLE
    clearDataLocalStorage()
  },
  SET_REALM_CLIENT_ID(state, payload) {
    switch (payload.status) {
      case PENDING:
        state.realClientIdErrorMessage &&
          (state.realClientIdErrorMessage = null)
        break

      case SUCCESS:
        state.realmClientId = payload.data
        break

      case ERROR:
        state.realClientIdErrorMessage = payload.errorMessage || ''
        break
    }
    state.realmClientId = payload.data
    state.realmClientStatus = payload.status
  }
}
