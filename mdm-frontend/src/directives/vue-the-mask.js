import Vue from 'vue'
import { mask } from 'vue-the-mask'

const vueTheMask = {
  bind(el, binding, vnode, oldVnode) {
    if (binding.value) {
      mask(el, binding, vnode, oldVnode)
    }
  }
}

Vue.directive('mask', vueTheMask)
