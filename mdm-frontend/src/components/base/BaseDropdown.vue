<template>
  <div
    :class="dynamicClass('dropdown__wrapper')"
    v-click-outside="close"
    :style="computedStyle"
  >
    <div class="dropdown__placeholder" v-if="isLoading"></div>
    <div
      :class="[
        'dropdown',
        isOpened ? 'dropdown--active' : '',
        hasError ? 'dropdown--error' : ''
      ]"
      v-else
    >
      <div @click="toggleDropdown" class="dropdown__controller">
        <p :class="dynamicClass('dropdown__label')">
          {{ label }}
        </p>
        <p
          :class="[
            ...dynamicClass('dropdown__value'),
            disabled && 'dropdown__value--disabled'
          ]"
        >
          <template v-if="!hasMultiSelection && !isSearchable">
            {{ selected ? selected[optionTitleField] : placeholder }}
          </template>
          <template v-else-if="isSearchable">
            <input
              type="text"
              class="dropdown__input"
              v-model="searchValue"
              :disabled="disabled"
              :placeholder="placeholder"
            />
          </template>
          <template v-else>
            <template v-if="selected && selected.length">
              <span v-for="(sl, idx) in selected" :key="sl[optionValueField]"
                >{{ sl[optionTitleField] }}
                {{ idx + 1 != selected.length ? ', ' : '' }}
              </span>
            </template>
            <span v-else>{{ placeholder }}</span>
          </template>
          <inline-svg
            :src="require(`@/assets/icons/chevron_down.svg`)"
            :class="[
              'dropdown__icon',
              isOpened ? 'dropdown__icon--active' : ''
            ]"
          />
        </p>
      </div>
      <InputErrors :errorProp="error"></InputErrors>
    </div>
    <ul :class="dynamicClass('dropdown__list')" v-if="isOpened">
      <template v-if="optionsToShow.length > 0">
        <li
          :class="[
            ...dynamicClass('dropdown__item'),
            isSelected(option) ? 'dropdown__item--selected' : ''
          ]"
          @click="onSelectValue(option)"
          v-for="option in optionsToShow"
        >
          {{ option[optionTitleField] }}
        </li>
      </template>
      <li class="dropdown__item" v-else>Нет данных</li>
    </ul>
  </div>
</template>

<script>
import InputErrors from '@/components/common/InputErrors.vue'

export default {
  components: {
    InputErrors
  },
  props: {
    value: {
      type: [String, Number, Boolean, Object, Array]
    },
    placeholder: {
      type: String,
      required: true
    },
    optionValueField: {
      type: String,
      default: 'id'
    },
    optionTitleField: {
      type: String,
      default: 'title'
    },
    hasMultiSelection: {
      type: Boolean,
      default: false
    },
    width: {
      type: Number,
      default: 100
    },
    label: {
      type: String,
      required: true
    },
    disabled: {
      type: Boolean,
      default: false
    },
    isLoading: {
      type: Boolean,
      default: false
    },
    error: {
      type: Object,
      required: false
    },
    options: {
      type: Array,
      requried: true
    },
    isSearchable: {
      type: Boolean,
      default: false
    }
  },

  data() {
    return {
      isOpened: false,
      selected: null,
      searchValue: null
    }
  },
  watch: {
    value: {
      handler(newValue) {
        if (newValue == null || !this.options) return
        if (this.hasMultiSelection) {
          this.selected = this.options.filter((opt) =>
            newValue.some(
              (nv) => nv[this.optionValueField] == opt[this.optionValueField]
            )
          )
        } else {
          this.selected = this.options.find(
            (el) => el[this.optionValueField] == newValue[this.optionValueField]
          )
        }
      },
      immediate: true
    },
    options: {
      handler(newValue) {
        this.options = newValue
        if (!newValue || !this.value) {
          this.selected = null
          return
        }
        if (this.hasMultiSelection) {
          this.selected = this.options.filter((opt) =>
            this.value.some(
              (nv) => nv[this.optionValueField] == opt[this.optionValueField]
            )
          )
        } else {
          this.selected = this.options.find(
            (el) => el[this.optionValueField] == newValue[this.optionValueField]
          )
        }
      },
      immediate: true
    }
  },
  computed: {
    hasError() {
      return this.error && this.error.$invalid && this.error.$dirty
    },
    computedStyle() {
      return { width: `${this.width}%`, height: `${this.height}px` }
    },
    optionsToShow() {
      return this.isSearchable
        ? this.options.filter((opt) =>
            opt[this.optionTitleField]
              .toLowerCase()
              .includes(this.searchValue.toLowerCase())
          )
        : this.options
    }
  },
  methods: {
    isSelected(option) {
      if (!this.hasMultiSelection) {
        return (
          this.selected &&
          this.selected[this.optionValueField] == option[this.optionValueField]
        )
      } else {
        return (
          this.selected &&
          this.selected.some(
            (sl) => sl[this.optionValueField] == option[this.optionValueField]
          )
        )
      }
    },
    toggleDropdown() {
      if (this.disabled) return
      this.isOpened = !this.isOpened
    },
    close() {
      this.isOpened = false
    },
    dynamicClass(className) {
      return [
        className,
        this.selected ? `${className}--active` : '',
        this.hasError ? `${className}--error` : ''
      ]
    },
    onSelectValue(option) {
      if (!this.hasMultiSelection) {
        this.selected = option
        this.searchValue = option[this.optionTitleField]
        this.$emit('input', option)
      } else {
        let optIdx =
          this.selected &&
          this.selected.findIndex(
            (sl) => sl[this.optionValueField] == option[this.optionValueField]
          )
        if (optIdx == null || optIdx == -1) {
          this.selected = this.selected ? [...this.selected, option] : [option]
        } else {
          this.selected.splice(optIdx, 1)
          this.selected.length == 0 && this.close()
        }
        this.$emit('input', this.selected)
      }
      this.close()
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
    padding-bottom: 20px;
    position: relative;
  }

  &__value {
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    width: 100%;
    transition: all 0.3s;
    position: relative;
    border-bottom: 2.7113px solid #2b2b2b;
    font-weight: 400;
    padding: 20px 15px 10px 0;
    font-size: 16px;
    color: rgba(17, 17, 17, 0.5);
    &--active {
      color: #2b2b2b;
    }
    &--error {
      border-bottom: 2px solid #eb5757;
    }
    &--disabled {
      border-bottom: 2.7113px solid rgba(78, 69, 69, 0.5);
      color: rgba(78, 69, 69, 0.5);
      cursor: not-allowed;
    }
  }
  &__label {
    font-weight: 400;
    font-size: 16px;
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
    font-family: Inter;
    font-weight: 600;
    color: #111111;
    cursor: pointer;
    display: flex;
    align-items: center;
    width: 100%;
    padding: 16px;
    font-weight: 600;
    font-size: 16px;

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
    height: 100%;
  }
  &__icon {
    position: absolute;
    top: 58%;
    right: 0;
    transition: all 0.3s;
    transform: translateY(-50%);
    &--active {
      transform: translateY(-50%) rotate(180deg);
    }
  }
  &__list {
    position: absolute;
    flex-direction: column;
    top: 100%;
    z-index: 23;
    overflow: hidden;
    max-height: 16rem;
    overflow-y: auto;
    width: 100%;
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
  &__input {
    width: 100%;
    height: 100%;
    border: none;
    outline: none;
    font-size: 16px;

    &:focus,
    &:hover {
      outline: none;
    }
    &::placeholder {
      font-size: 16px;
      color: rgba(17, 17, 17, 0.5);
    }
  }
}
</style>
