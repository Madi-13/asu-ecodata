import { apiStatus } from '@/api/constants/apiStatus'
const { PENDING, SUCCESS, ERROR } = apiStatus

export const apiPendingFactory = () => {
  return {
    status: PENDING
  }
}

export const apiSuccessFactory = (data) => {
  return {
    status: SUCCESS,
    data
  }
}

export const apiErrorFactory = (error, errorMessages = []) => {
  return {
    status: ERROR,
    error,
    errorMessages
  }
}
