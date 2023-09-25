<template>
  <div class="chip" @click.prevent.self="$emit('clickOnChip')">
    {{ title }}
    <img
      src="../../assets/icons/close.svg"
      alt=""
      class="ml--2 chip__icon"
      @click.prevent="$emit('remove')"
    />
  </div>
</template>

<script>
export default {
  props: {
    title: {
      type: String
    }
  }
}
</script>

<style lang="scss" scoped>
.chip {
  background: #0a7937;
  padding: 8px 10px;
  border-radius: 10px;
  outline: none;
  color: white;
  font-size: 18px;
  display: flex;
  align-items: center;
  width: fit-content;
  max-width: 100%;
  &__icon {
    cursor: pointer;
  }
}
</style>
