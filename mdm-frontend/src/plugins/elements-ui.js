import lang from 'element-ui/lib/locale/lang/en'
import locale from 'element-ui/lib/locale'

import Vue from 'vue'
import 'element-ui/lib/theme-chalk/table.css'
import 'element-ui/lib/theme-chalk/table-column.css'
import 'element-ui/lib/theme-chalk/select.css'
import 'element-ui/lib/theme-chalk/option.css'
import 'element-ui/lib/theme-chalk/button.css'
import 'element-ui/lib/theme-chalk/icon.css'
// configure language
locale.use(lang)

Vue.component('el-table', () =>
  import(/* webpackChunkName: "el-table" */ 'element-ui/lib/table')
)

Vue.component('el-table-column', () =>
  import(
    /* webpackChunkName: "el-table-column" */ 'element-ui/lib/table-column'
  )
)

Vue.component('el-button', () =>
  import(/* webpackChunkName: "el-button" */ 'element-ui/lib/button')
)

Vue.component('el-select', () =>
  import(/* webpackChunkName: "el-table-column" */ 'element-ui/lib/select')
)

Vue.component('el-option', () =>
  import(/* webpackChunkName: "el-table-column" */ 'element-ui/lib/option')
)
