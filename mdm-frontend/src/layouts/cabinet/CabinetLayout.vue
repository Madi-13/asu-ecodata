<template>
  <div class="layout">
    <Sidebar class="sticky" />
    <main class="main">
      <Header />
      <router-view></router-view>
    </main>
  </div>
</template>

<script>
import { initObjectTypes } from '@/services/columnTypesService'

import Header from './components/Header.vue'
import Sidebar from './components/Sidebar.vue'
export default {
  components: {
    Header,
    Sidebar
  },
  created() {
    initObjectTypes()
  }
}
</script>

<style lang="scss" scoped>
.layout {
  display: flex;
  min-height: 100vh;
}
.main {
  padding: 4.3rem 3.9rem 0 2rem;
  width: 75%;
}
</style>
