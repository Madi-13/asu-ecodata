import Vue from 'vue'

const clickOutside = {
  bind: function (el, binding) {
    binding.event = function (event) {
      event.stopPropagation()
      // event.preventDefault()
      if (!(el === event.target || el.contains(event.target))) {
        if (binding.value instanceof Function) {
          binding.value(event)
        }
      }
    }
    document.body.addEventListener('click', binding.event)
  },
  unbind: function (el, binding) {
    document.body.removeEventListener('click', binding.event)
  }
}

Vue.directive('click-outside', clickOutside)
