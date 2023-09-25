import {
  login,
  checkTokenValidity,
  refreshToken,
  getPermissions
} from '@/api/authAPI'

import { getRealmClient } from '@/api/roleAPI'

import { withAsync } from '@/api/helpers/withAsync'
import {
  apiPendingFactory,
  apiSuccessFactory,
  apiErrorFactory
} from '@/api/helpers/apiStateFactory'

import { serializePermissions } from '@/api/serializers/permissions'

export default {
  async userLogin(context, { username, password }) {
    context.commit('SET_USER_LOGIN_STATE', apiPendingFactory())
    const { response, error } = await withAsync(login, undefined, {
      auth: {
        username,
        password
      }
    })

    if (error) {
      context.commit(
        'SET_USER_LOGIN_STATE',
        apiErrorFactory(error, error.message)
      )
      return
    }

    context.commit('SET_USER_LOGIN_STATE', apiSuccessFactory(response.data))
  },
  async checkToken(context, accessToken) {
    context.commit('SET_CHECK_TOKEN_STATE', apiPendingFactory())
    const { response, error } = await withAsync(
      checkTokenValidity,
      {},
      {
        headers: {
          Authorization: `Bearer ${accessToken}`
        }
      }
    )

    if (error) {
      context.commit(
        'SET_CHECK_TOKEN_STATE',
        apiErrorFactory(error, error.message)
      )
      return
    }

    context.commit('SET_CHECK_TOKEN_STATE', apiSuccessFactory(response.data))
  },
  async refreshToken(context, payload) {
    context.commit('SET_REFRESH_TOKEN_STATE', apiPendingFactory())

    const { response, error } = await withAsync(refreshToken, payload)

    if (error) {
      context.commit(
        'SET_REFRESH_TOKEN_STATE',
        apiErrorFactory(error, error.message)
      )
      return
    }

    context.commit('SET_REFRESH_TOKEN_STATE', apiSuccessFactory(response.data))
  },
  async getPermissions(context, payload) {
    context.commit('SET_PERMISSIONS', apiPendingFactory())

    const { response, error } = await withAsync(getPermissions, payload)

    if (error) {
      context.commit('SET_PERMISSIONS', apiErrorFactory(error, error.message))
      return
    }

    context.commit(
      'SET_PERMISSIONS',
      apiSuccessFactory(serializePermissions(response.data))
    )
  },
  async getRealmClientId(context) {
    context.commit('SET_REALM_CLIENT_ID', apiPendingFactory())

    const { response, error } = await withAsync(getRealmClient)

    if (error) {
      context.commit(
        'SET_REALM_CLIENT_ID',
        apiErrorFactory(error, error.message)
      )
      return
    }
    let client = response.data.find((el) => el.clientId == 'mdm-cli')
    context.commit('SET_REALM_CLIENT_ID', apiSuccessFactory(client.id))
  }
}
