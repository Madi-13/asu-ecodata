<template>
  <label
    :for="`id_${name}`"
    class="display--flex align--center checkbox__wrapper"
    :style="[width !== 'auto' ? { width: `${width}%` } : { width: width }]"
    @click.prevent="checkboxClicked"
  >
    <input
      class="checkbox"
      type="checkbox"
      :id="`id_${name}`"
      :name="name"
      :checked="value"
    />
    <span :class="dynamicClass"></span>
    <span
      :class="['checkbox__label', `checkbox__label--${color}`]"
      v-if="hasDefaultSlot"
      ><slot></slot
    ></span>
  </label>
</template>

<script>
export default {
  inheritAttrs: false,
  name: 'BaseCheckbox',
  props: {
    name: {
      type: String,
      default: 'check'
    },
    width: {
      type: String,
      default: 'auto'
    },
    size: {
      type: String,
      default: 'normal'
    },
    color: {
      type: String,
      default: 'green'
    },
    value: {
      type: Boolean
    },
    disabled: {
      type: Boolean,
      default: false
    },
    hasError: { type: Boolean, default: false },
    error: { type: Object, required: false }
  },
  computed: {
    dynamicClass() {
      return [
        'checkmark',
        `checkmark--${this.color}`,
        `checkmark--${this.size}`,
        this.disabled ? `checkmark--disabled` : ''
      ]
    },
    hasDefaultSlot() {
      return !!this.$slots.default
    }
  },
  methods: {
    checkboxClicked() {
      if (!this.disabled) {
        this.$emit('input', !this.value)
      }
    }
  }
}
</script>

<style scoped lang="scss">
.checkbox {
  opacity: 0;
  width: 0;
  height: 0;
  &__wrapper {
    display: inline-flex;
    display: flex;
    align-items: center;
    // padding-bottom: 20px;
  }
  &:checked {
    & ~ .checkmark {
      &--dark {
        background: #2d4059 url('../../assets/icons/checkbox.svg') no-repeat
          center;
      }
      &--green {
        background: #0a7937 url('../../assets/icons/checkbox.svg') no-repeat
          center;
      }
    }
  }
  &__label {
    cursor: pointer;
    display: block;
    font-weight: 400;
    font-size: 16px;
    color: #111111;
    margin-left: 12px;
    &--dark {
      color: #69778f;
    }
  }

  &__error {
    &-wrapper {
      margin-top: 5px;
      margin-left: 28px;
    }
    color: red;
    font-size: 12px;
  }
  &--hasError {
    border: 1px solid red;
  }
}
.checkmark {
  background: var(--light);
  border: 1px solid #bdbdbd;
  border-radius: 4px;
  width: 20px;
  height: 20px;
  display: block;
  cursor: pointer;
  &--normal {
    width: 2rem;
    height: 2rem;
  }
  &--large {
    width: 3rem;
    height: 3rem;
  }
  &--disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }
}
</style>
