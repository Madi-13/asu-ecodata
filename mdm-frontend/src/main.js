import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import { registerBaseComponents } from '@/helpers/registerBaseComponents'

import '@/directives/clickOutside'
import '@/directives/permission'
import '@/directives/vue-the-mask'
import '@/directives/scrollToBottom'

import '@/plugins/modal'
import '@/plugins/vuelidate'
import '@/plugins/elements-ui'
import '@/plugins/vue-toastification'
import '@/plugins/inline-svg'
import '@/plugins/vuejs-pagination'
import '@/plugins/vue-loading-skeleton'

import '@/api/interceptor/authInterceptor'

import '@/filters/dateFilters'

registerBaseComponents(Vue)
Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: (h) => h(App)
}).$mount('#app')
