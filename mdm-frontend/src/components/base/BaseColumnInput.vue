<template>
  <input
    v-on="inputListeners"
    v-bind="{ ...$attrs, value, disabled }"
    :type="type"
    autocomplete="off"
    v-mask="mask"
    :id="name"
    class="input"
  />
</template>

<script>
export default {
  props: {
    name: {
      type: String,
      required: true
    },
    width: {
      type: Number,
      default: 100
    },
    value: {
      type: [String, Number],
      default: null
    },
    type: {
      type: String,
      default: 'text'
    },
    disabled: {
      type: Boolean,
      default: false
    },
    mask: {
      type: String,
      required: false
    }
  }
}
</script>

<style lang="scss" scoped>
.input {
  border: none;
  outline: none;
  font-weight: 400;
  font-size: 16px;
  color: #2b2b2b;
}
</style>
