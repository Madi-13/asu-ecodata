<template>
  <div class="button__wrapper" :style="{ width, height }">
    <button
      :class="buttonClass"
      :disabled="disabled || isLoading"
      v-on="$listeners"
    >
      <BaseSpiner v-if="isLoading" :size="4"></BaseSpiner>
      <template v-else>
        <slot></slot>
      </template>
    </button>
  </div>
</template>

<script>
export default {
  props: {
    color: {
      type: String,
      default: 'green'
    },
    disabled: {
      type: Boolean,
      default: false
    },
    isLoading: {
      type: Boolean,
      default: false
    },
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '100%'
    }
  },
  computed: {
    buttonClass() {
      return [
        'button__controller',
        `button__controller--${this.color}`,
        `button__controller--${this.size}`
      ]
    }
  }
}
</script>

<style lang="scss" scoped>
.button {
  &__controller {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 15px 41px;
    cursor: pointer;
    border: none;
    border-radius: 10px;
    font-weight: 400;
    font-size: 16px;
    color: #ffffff;
    height: 100%;
    width: 100%;
    transition: all 0.3s;
    &:hover,
    &:active,
    &:focus {
      outline: none;
    }
    &:disabled {
      opacity: 0.5;
      cursor: not-allowed;
    }

    &--green {
      background: #0a7937;
      &:active {
        background: #0a9d9d;
      }
      &:hover {
        background: #025825;
      }
    }
    &--red {
      background: #ff2525;
      &:hover {
        background: #f01515;
      }
      &:active {
        background: #e90202;
      }
    }
    &--blue {
      background: #2f80ed;
      &:hover {
        background: #136ee6;
      }
      &:active {
        background: #015ace;
      }
    }
    &--outline {
      &-grey {
        background: white;
        border: 2px solid #acacac;
        border-radius: 10px;
        font-size: 16px;
        color: #acacac;
      }
      &-dark {
        background: white;
        border: 2px solid #2b2b2b;
        border-radius: 10px;
        font-size: 16px;
        color: #2b2b2b;
      }
    }
  }
}
</style>
