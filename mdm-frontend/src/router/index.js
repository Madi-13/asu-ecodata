import Vue from 'vue'
import VueRouter from 'vue-router'
import store from '@/store'

import authRoutes from '@/views/auth/authRoutes'
import cabinetRoutes from '@/views/cabinet/cabinetRoutes'

import { getDataFromLocalStorage } from '@/helpers/permanentStorageHelper'

Vue.use(VueRouter)

const routes = [
  ...authRoutes,
  ...cabinetRoutes,
  { path: '/', redirect: '/cabinet' }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
  scrollBehavior() {
    return { x: 0, y: 0, behavior: 'smooth' }
  }
})

router.beforeEach(async (to, from, next) => {
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    const accessToken =
      store.getters.getTokens.access_token ||
      getDataFromLocalStorage('mdmAccessToken')
    if (accessToken) {
      if (!store.getters.isCheckTokenAlreadyCalled) {
        await store.dispatch('checkToken', accessToken)
        await store.dispatch('getPermissions')
        await store.dispatch('getRealmClientId')
        if (store.getters.getCheckTokenSuccessful) {
          if (store.getters.getPermissions[to.meta.resourceName]) {
            next()
          } else {
            next('/cabinet/analytics')
          }
        } else {
          next('/auth/login')
        }
      } else {
        next()
      }
    } else {
      next('/auth/login')
    }
  } else {
    next()
  }
})

const DEFAULT_TITLE = 'MDM'

router.afterEach((to, from) => {
  Vue.nextTick(() => {
    document.title = to.meta.title
      ? `${DEFAULT_TITLE} | ${to.meta.title}`
      : DEFAULT_TITLE
  })
})

export default router
