<template>
  <textarea
    name="textarea"
    :rows="computedRows"
    cols="50"
    :class="['textarea', theme == 'light' ? 'textarea--light' : '']"
    @input="$emit('input', $event.target.value)"
    :value="value"
  ></textarea>
</template>

<script>
export default {
  props: {
    value: {
      type: String,
      default: ''
    },
    theme: {
      type: String,
      default: 'dark'
    }
  },
  computed: {
    computedRows() {
      return this.value && this.value.split(/\n/).length > 5
        ? this.value.split(/\n/).length
        : 5
    }
  }
}
</script>

<style lang="scss" scoped>
.textarea {
  background: #f2f2f2;
  border-radius: 8px;
  padding: 16px;
  border: none;
  font-family: Inter;
  font-style: normal;
  font-weight: 600;
  font-size: 16px;
  color: #111111;
  width: 100%;
  resize: none;
  &--light {
    background: #fff;
    border: 1px solid rgba(17, 17, 17, 0.32);
  }
  &:hover,
  &:focus {
    outline: none;
  }
}
</style>
