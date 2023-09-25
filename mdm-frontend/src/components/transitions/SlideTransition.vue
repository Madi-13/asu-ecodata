<template>
  <transition name="slide-fade" mode="out-in" v-bind="$attrs">
    <slot />
  </transition>
</template>

<script>
export default {
  inheritAttrs: false
}
</script>

<style scoped>
.slide-fade-enter-active {
  transition: all 0.3s ease;
}
.slide-fade-leave-active {
  transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}
.slide-fade-enter,
.slide-fade-leave-to {
  transform: translateX(10px);
  opacity: 0;
}
</style>
