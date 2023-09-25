import Vue from 'vue'
import store from '@/store'

const permissionAcl = {
  bind: canImplementation,
  update: canImplementation
}

function canImplementation(el, { value, arg, modifiers }, vnode) {
  const behaviour = modifiers.disable ? 'disable' : 'hide'
  const scopeName = arg
  const resourceName = value.resource
  const permissionName = value.permission
  const resourcePermissions = store.getters.getPermissions[resourceName]
  const permissionScopes = store.getters.getPermissions?.[resourceName]?.[
    permissionName
  ]?.scopes || ['read']

  if (!resourcePermissions) {
    if (behaviour == 'hide') {
      el.style.display = 'none'
    } else {
      el.disabled = true
    }
  } else {
    el.style.display = 'block'
    el.disabled = false
  }

  if (!permissionScopes.includes(scopeName)) {
    if (behaviour == 'hide') {
      el.style.display = 'none'
    } else {
      if (el.children[0].children[1]) {
        el.children[0].children[1].disabled = true
      } else {
        el.children[0].disabled = true
        el.style.opacity = 0.5
        el.style.cursor = 'not-allowed'
      }
    }
  }
}

Vue.directive('can', permissionAcl)
