import axios from 'axios'

const axiosParams = {
  baseUrl:
    process.env.NODE_ENV === 'development' ? 'http://localhost:8080' : '/',
  headers: { 'Content-Type': 'application/json' }
}

export const axiosInstance = axios.create(axiosParams)

axiosInstance.defaults.headers.post['Content-Type'] = 'application/json'

const didAbort = (error) => axios.isCancel(error)

const getCancelSource = () => axios.CancelToken.source()

const withAbort =
  (fn) =>
  async (...args) => {
    const originalConfig = args[args.length - 1]
    let { abort, ...config } = originalConfig

    if (typeof abort === 'function') {
      const { cancel, token } = getCancelSource()
      config.cancelToken = token
      abort(cancel)
    }

    try {
      return await fn(...args.slice(0, args.length - 1), config)
    } catch (error) {
      didAbort(error) && (error.aborted = true)
      throw error
    }
  }

const api = (axios) => {
  return {
    get: (url, config = {}) => withAbort(axios.get)(url, config),
    post: (url, payload, config = {}) => {
      return axios.post(url, payload, config)
    },
    put: (url, payload, config = {}) => {
      return axios.put(url, payload, config)
    },
    delete: (url, config = {}) => {
      return axios.delete(url, config)
    }
  }
}

export default api(axiosInstance)
