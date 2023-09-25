import Vue from 'vue'

Vue.component('paginate', () =>
  import(/* webpackChunkName: "paginate" */ 'vuejs-paginate')
)
