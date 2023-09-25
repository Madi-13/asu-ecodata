<template>
  <div :class="dynamicClass('dropdown__wrapper')" v-click-outside="close">
    <slot :toggle="toggleDropdown" :selected="selected"  name="input">
      <div
        :class="[
        'dropdown',
        disabled ? 'dropdown--disabled' : '',
        isOpened ? 'dropdown--active' : ''
      ]"
      >
        <div
          @click="toggleDropdown"
          :class="[
          'dropdown__controller',
          isOpened ? 'dropdown__controller--active' : ''
        ]"
        >
          <div :class="dynamicClass('dropdown__value')">
            {{ selected ? selected[optionTitleField] : null }}
          </div>
          <!-- <inline-svg
            :src="require(`@/assets/icons/chevron_down.svg`)"
            :class="['dropdown__icon', isOpened ? 'dropdown__icon--active' : '']"
          /> -->
        </div>
      </div>
    </slot>
    <ul
      :class="dynamicClass('dropdown__list')"
      v-if="isOpened"
      :style="{
        left: `${listLeftPosition}px`,
        top: `${listTopPosition}px`,
        width: `${listWidth}px`
      }"
    >
      <li
        :class="[
          ...dynamicClass('dropdown__item'),
          selected && selected[optionValueField] == option[optionValueField]
            ? 'dropdown__item--selected'
            : ''
        ]"
        @click="onSelectValue(option)"
        v-for="option in options"
      >
        {{ optionTitleField ? option[optionTitleField] : option.name }}
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  props: {
    value: {
      type: [String, Number, Boolean, Object]
    },
    optionValueField: {
      type: String,
      default: 'id'
    },
    optionTitleField: {
      type: String,
      default: 'name'
    },
    width: {
      type: Number,
      default: 100
    },
    disabled: {
      type: Boolean,
      default: false
    },
    options: {
      type: Array,
      requried: true
    },
    error: {
      type: Object,
      required: false
    }
  },

  data() {
    return {
      isOpened: false,
      selected: null,
      listTopPosition: null,
      listLeftPosition: null,
      listWidth: null
    }
  },
  watch: {
    value: {
      handler(newValue) {
        if (!this.options || !newValue) {
          return
        }
        this.selected = this.options.find(
          (el) => el[this.optionValueField] == newValue[this.optionValueField]
        )
      },
      immediate: true
    },
    options: {
      handler(newValue) {
        this.selected = newValue.find(
          (el) => el[this.optionValueField] == this.value[this.optionValueField]
        )
      }
    }
  },
  computed: {
    restOptions() {
      return this.options.filter((option) =>
        !this.selected
          ? true
          : option[this.optionValueField] !=
            this.selected[this.optionValueField]
          ? true
          : false
      )
    },
    hasError() {
      return this.error && this.error.$invalid && this.error.$dirty
    }
  },
  methods: {
    toggleDropdown(event) {
      if (this.disabled) return
      this.isOpened = !this.isOpened
      let controllerPosition = event.target.getBoundingClientRect()
      this.listTopPosition = controllerPosition.y + controllerPosition.height
      this.listLeftPosition = controllerPosition.x
      this.listWidth = controllerPosition.width
    },
    close() {
      this.isOpened = false
    },
    dynamicClass(className) {
      return [
        className,
        this.selected ? `${className}--active` : '',
        this.hasError && `${className}--error`
      ]
    },
    onSelectValue(option) {
      this.isOpened = false
      this.selected = option
      this.$emit('input', option)
      this.$emit('selected', option)
    }
  }
}
</script>

<style lang="scss" scoped>
.dropdown {
  overflow: hidden;
  background: #ffffff;
  width: 100%;
  cursor: pointer;

  height: 100%;
  z-index: 3;
  &__wrapper {
    width: 100%;
    position: relative;
    &--error {
      border: 1px dashed red;
    }
  }

  &--error {
    border: 1px solid #eb5757;
  }

  &--disabled {
    background: #f2f2f2;
    border: none;
    cursor: not-allowed;
  }

  &__value {
    width: 100%;
    transition: all 0.3s;
    font-weight: 400;
    font-size: 16px;
    padding-right: 2rem;
    color: rgba(17, 17, 17, 0.5);
    &--active {
      font-size: 16px;
      color: #2b2b2b;
    }
  }
  &__label {
    font-weight: 400;
    font-size: 20px;
    color: #acacac;
    width: 100%;
    transition: all 0.3s;
    &--active {
      top: 0%;
    }
    &--small {
      display: none;
    }
  }
  &__item {
    color: #111111;
    cursor: pointer;
    display: flex;
    align-items: center;
    padding: 16px;

    &--selected {
      background: #acacac;
      color: white;
    }

    &:hover {
      background: var(--green-light);
      color: white;
    }
  }
  &__controller {
    width: 100%;
    min-height: 20px;
    position: relative;
    &::after {
      content: '';
      position: absolute;
      background-image: url(http://localhost:8080/img/chevron_down.06333bbe.svg);
      right: 0px;
      top: 50%;
      height: 25px;
      width: 25px;
      z-index: 123;
      transform: translateY(-50%);
      background-size: cover;
      background-repeat: no-repeat;
    }
    &--active {
      &::after {
        transform: translateY(-45%) rotate(180deg);
      }
    }
  }
  &__list {
    position: fixed;
    flex-direction: column;
    z-index: 23;
    overflow: hidden;
    max-height: 16rem;
    overflow-y: auto;
    min-height: 40px;
    border: 1px solid #2b2b2b;
    border-radius: 8px;
    background: white;
    display: flex;

    &--error {
      border-bottom: 2px solid #eb5757;
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
    border-radius: 8px;
  }
}
</style>
