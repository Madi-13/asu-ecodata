export const withAsync = async (fn, ...args) => {
  try {
    const response = await fn(...args)
    if (!response?.data) {
      return { response: null, error: null }
    }
    const { db_code, db_error_desc } = response.data
    if (!!db_code && db_code != 0) {
      return {
        response: null,
        error: {
          message: db_error_desc,
          data: db_code
        }
      }
    }
    return { response, error: null }
  } catch (error) {
    return {
      response: null,
      error: {
        message: error?.response?.data?.error || 'Запрос не удался',
        data: error?.response?.data
      }
    }
  }
}
