<template>
  <transition name="from-side" mode="in-out" v-bind="$attrs">
    <slot />
  </transition>
</template>

<script>
export default {
  inheritAttrs: false
}
</script>

<style scoped>
.from-side-enter-active {
  transition: all 0.5s ease;
}
.from-side-leave-active {
  transition: all 0.5s cubic-bezier(1, 0.5, 0.8, 1);
}
.from-side-enter,
.from-side-leave-to {
  transform: translateX(100vw);
  opacity: 0;
}
</style>
