import Vue from 'vue'

const Modal = {
  install(Vue) {
    this.event = new Vue()

    Vue.prototype.$modal = {
      show(modalName, params = {}) {
        Modal.event.$emit('show', modalName, params)
      },
      hide(modalName) {
        Modal.event.$emit('hide', modalName)
      },
      $event: this.event
    }
  }
}

Vue.use(Modal)
