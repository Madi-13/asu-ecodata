import Vue from 'vue'

const scrollToBottom = {
  inserted: function (el, binding) {
    binding.event = function (event) {
      if (
        event.target.scrollHeight - event.target.scrollTop ===
        event.target.clientHeight
      ) {
        if (binding.value instanceof Function) {
          binding.value(event)
        }
      }
    }
    el.addEventListener('scroll', binding.event)
  },
  unbind: function (el, binding) {
    el.removeEventListener('scroll', binding.event)
  }
}

Vue.directive('scroll-to-bottom', scrollToBottom)
