import { apiStatus } from '@/api/constants/apiStatus'
const { SUCCESS, IDLE } = apiStatus

export default {
  getLoginStatus: (state) => state.loginStatus,
  getTokens: (state) => state.token,
  getUsername: (state) => state.username,
  getPermissions: (state) => state.permissions,
  getRealmClientId: (state) => state.realmClientId,
  getCheckTokenSuccessful: (state) => state.checkTokenStatus == SUCCESS,
  isCheckTokenAlreadyCalled: (state) => state.checkTokenStatus != IDLE,
  getRefreshTokenStatus: (state) => state.refreshTokenStatus,
  getRefreshTokenIsSuccess: (state) => state.refreshTokenStatus == SUCCESS
}
