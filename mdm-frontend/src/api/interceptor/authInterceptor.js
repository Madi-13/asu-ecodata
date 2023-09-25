import { axiosInstance } from '@/api/api'
import { getDataFromLocalStorage } from '@/helpers/permanentStorageHelper'
import store from '@/store'
import router from '@/router'

function getJwtTokens() {
  return {
    accessToken:
      store.getters.getTokens.access_token ||
      getDataFromLocalStorage('mdmAccessToken'),
    refreshToken:
      store.getters.getTokens.refresh_token ||
      getDataFromLocalStorage('mdmRefreshToken')
  }
}

axiosInstance.interceptors.request.use(
  async (config) => {
    let tokens = getJwtTokens()
    if (tokens.accessToken) {
      config.headers['Authorization'] = `Bearer ${tokens.accessToken}`
      config.headers['Content-Type'] = 'application/json'
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)
axiosInstance.interceptors.response.use(
  (response) => {
    return response
  },
  async function (error) {
    const originalRequest = error.config
    if (
      error.response &&
      error.response.status === 403 &&
      !originalRequest._retry
    ) {
      try {
        const response = await store.dispatch('refreshToken', {
          refresh_token: getJwtTokens().refreshToken
        })
        if (store.getters.getRefreshTokenIsSuccess) {
          axiosInstance.defaults.headers.common[
            'Authorization'
          ] = `Bearer ${store.getters.getTokens.access_token}`
          originalRequest.headers.Authorization = `Bearer ${store.getters.getTokens.access_token}`
          originalRequest.headers.ContentType = 'application/json'
          originalRequest._retry = true
        } else {
          router.push('/auth')
          return Promise.reject(error)
        }
        return axiosInstance.request(originalRequest)
      } catch (error) {
        router.push('/auth')
        return Promise.reject(error)
      }
    }
    return Promise.reject(error)
  }
)
