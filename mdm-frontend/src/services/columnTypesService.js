import Vue from 'vue'
import { getColumnTypes } from '@/api/objectsAPI'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus


const state = Vue.observable({
  objectTypes: null,
  objectTypesStatus: IDLE
})

export const objectTypesComputed = {
  ...apiStatusComputedFactory(['objectTypesStatus']),
  objectTypes: {
    get() {
      return state.objectTypes
    }
  }
}

export const initObjectTypes = async () => {
    state.objectTypesStatus = PENDING
    const {response, error} = await withAsync(getColumnTypes)
    if(error){
        Vue.$toast.error("Не удалось подгрузить список типов", {
            timeout: 2000
        });
        state.objectTypesStatus = ERROR;
        return
    }
    state.objectTypesStatus = SUCCESS;
    state.objectTypes = response.data;
}
