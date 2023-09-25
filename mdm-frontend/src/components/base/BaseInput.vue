<template>
  <div>
    <div
      :class="['input__wrapper', !hasValidation && 'input__wrapper--errorless']"
      :style="computedStyle"
      v-if="appereance == 'normal'"
    >
      <label :for="name" :class="dynamicClass('input__label')">{{
        label
      }}</label>
      <div class="input__placeholder" v-if="isLoading"></div>
      <textarea
        @input="onInput($event)"
        @focus="$emit('focus', $event)"
        :id="name"
        v-bind="{ ...$attrs, value, disabled }"
        cols="30"
        :class="dynamicClass('input__controller')"
        :rows="computedRows"
        v-else-if="type == 'textarea'"
      ></textarea>
      <input
        @input="onInput($event)"
        v-else
        v-bind="{ ...$attrs, value, disabled }"
        :type="type"
        autocomplete="off"
        @focus="$emit('focus', $event)"
        v-mask="mask"
        :id="name"
        :class="dynamicClass('input__controller')"
      />
      <InputErrors :errorProp="error"></InputErrors>
    </div>
    <template v-if="appereance == 'transparent'">
      <textarea
        @input="onInput($event)"
        @focus="$emit('focus', $event)"
        :id="name"
        v-bind="{ ...$attrs, value, disabled }"
        cols="30"
        :class="['input--transparent', hasError && 'input--transparent--error']"
        :rows="computedRows"
        v-if="type == 'textarea'"
      ></textarea>
      <input
        @input="onInput($event)"
        v-bind="{ ...$attrs, value, disabled }"
        :type="type"
        v-else
        @focus="$emit('focus', $event)"
        autocomplete="off"
        v-mask="mask"
        :id="name"
        :class="['input--transparent', hasError && 'input--transparent--error']"
      />
    </template>
  </div>
</template>

<script>
import InputErrors from '@/components/common/InputErrors.vue'
export default {
  inheritAttrs: false,
  components: {
    InputErrors
  },
  props: {
    appereance: {
      type: String,
      default: 'normal'
    },
    name: {
      type: String,
      required: false
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
    label: {
      type: String
      // required: true
    },
    disabled: {
      type: Boolean,
      default: false
    },
    mask: {
      type: String,
      required: false
    },
    hasValidation: {
      type: Boolean,
      default: true
    },
    isLoading: {
      type: Boolean,
      default: false
    },
    error: {
      type: Object,
      required: false
    }
  },
  computed: {
    computedStyle() {
      return { width: `${this.width}%`, height: `${this.height}px` }
    },
    hasError() {
      return this.error && this.error.$invalid && this.error.$dirty
    },
    computedRows() {
      return this.value && this.value.split(/\n/).length > 1
        ? this.value.split(/\n/).length
        : 1
    }
  },
  data() {
    return {
      isLabelVisible: false
    }
  },
  watch: {
    value: {
      handler(value) {
        value ? (this.isLabelVisible = true) : (this.isLabelVisible = false)
      },
      immediate: true
    }
  },
  methods: {
    onInput(event) {
      this.$emit(
        'input',
        this.type == 'number' ? +event.target.value : event.target.value
      )
    },
    dynamicClass(className) {
      return [
        className,
        this.disabled ? `${className}--disabled` : '',
        this.hasError ? `${className}--error` : ''
      ]
    }
  }
}
</script>

<style lang="scss" scoped>
.input {
  &__wrapper {
    position: relative;
    overflow: hidden;
    padding-bottom: 20px;
    &--errorless {
      padding-bottom: 0 !important;
    }
    &--small {
      padding-bottom: 0;
    }
  }
  &__label {
    font-weight: 400;
    font-size: 16px;
    color: #2b2b2b;
    width: 100%;
    height: 13px;
    padding: 6px 0px;
    top: -100%;
    transition: all 0.3s;

    &--disabled {
      cursor: not-allowed;
    }
  }
  &__controller {
    background: #ffffff;
    border: none;
    border-bottom: 2px solid #2b2b2b;
    padding: 20px 0 10px;
    box-sizing: border-box;
    outline: none;
    width: 100%;
    transition: all 0.3s;
    font-weight: 400;
    font-size: 16px;
    color: #2b2b2b;
    &:active {
      outline: none;
      border-bottom: 2.7113px solid #2b2b2b;
      padding: 20px 0 10px;
    }
    &:focus,
    &:hover {
      outline: none;
      padding: 20px 0 10px;
    }
    &:disabled {
      color: #acacac;
      border-bottom: 2.7113px solid #acacac;
      cursor: not-allowed;
    }
    &::placeholder {
      font-size: 16px;
      color: rgba(17, 17, 17, 0.5);
    }
    &--error {
      border-bottom: 2.7113px solid #eb5757;
      color: #eb5757;
    }
  }
  &__placeholder {
    background-image: linear-gradient(
      to right,
      #f6f7f8 0%,
      #edeef1 10%,
      #f6f7f8 20%,
      #f6f7f8 100%
    );
    background-size: 200% 100%;
    height: 56px;
    animation: bgPos 1s linear infinite;
    width: 100%;
    font-size: 16px;
    border-radius: 8px;
  }

  &--transparent {
    border: none;
    outline: none;
    font-weight: 400;
    width: 100%;
    font-size: 16px;
    color: #2b2b2b;
    &--error {
      border: 1px dashed red;
    }
  }
}
</style>
