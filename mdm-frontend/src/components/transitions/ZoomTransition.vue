<template>
  <transition name="zoom" mode="out-in" v-bind="$attrs">
    <slot />
  </transition>
</template>

<script>
export default {
  inheritAttrs: false
}
</script>

<style scoped>
.zoom-enter-active,
.zoom-leave-active {
  animation-duration: 0.5s;
  animation-fill-mode: both;
  animation-name: zoom;
}

.zoom-leave-active {
  animation-direction: reverse;
}
</style>
