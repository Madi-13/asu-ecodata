<template>
  <div class="input__wrapper" :style="computedWith" v-click-outside="close">
    <div :class="dynamicClass('input-controller__wrapper')">
      <label :for="name" :class="dynamicClass('input__label')">{{
        label
      }}</label>
      <div class="input">
        <input
          v-on="inputListeners"
          v-bind="{ ...$attrs, disabled }"
          type="text"
          :value="inputValue"
          autocomplete="off"
          :placeholder="placeholder"
          :id="name"
          @focus="onInputFocus"
          :class="dynamicClass('input-controller')"
        />
        <img
          src="../../assets/icons/chevron_down.svg"
          alt=""
          :class="['input__icon', isResultVisible && 'input__icon--reverse']"
        />
      </div>
    </div>
    <div
      :class="dynamicClass('input-result__list')"
      v-if="isResultVisible"
      v-scroll-to-bottom="handleScroll"
    >
      <div
        :class="[
          'input-result__item',
          !!(value && item[optionValueField] == value.id) &&
            'input-result__item--selected'
        ]"
        v-for="(item, index) in result"
        :key="index"
        @click.stop="optionSelected(item)"
      >
        {{ item[optionTitleField] }}
      </div>
      <div
        class="input-result__item input-result__static"
        v-if="getResultsStatusPending"
      >
        <BaseSpiner :size="3" color="#2CCBCB"></BaseSpiner>
      </div>
      <div
        class="input-result__item input-result__static"
        v-else-if="getResultsStatusSuccess && !hasResult && offset == 1"
      >
        Записей не найдено
      </div>
    </div>
    <InputErrors :errorProp="error"></InputErrors>
  </div>
</template>

<script>
import InputErrors from '@/components/common/InputErrors.vue'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
import debounce from 'lodash-es/debounce'

const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

export default {
  inheritAttrs: false,
  components: {
    InputErrors
  },
  props: {
    optionValueField: {
      type: String,
      default: 'id'
    },
    optionTitleField: {
      type: String,
      default: 'value'
    },
    searchMethod: {
      type: Function,
      required: true
    },
    searchParams: {
      type: Object,
      default: () => {}
    },
    name: {
      type: String,
      required: true
    },
    placeholder: {
      type: String,
      required: true
    },
    width: {
      type: Number,
      default: 100
    },
    value: {
      type: Object
    },
    label: {
      type: String,
      required: true
    },
    disabled: {
      type: Boolean,
      default: false
    },
    error: {
      type: Object,
      required: false
    }
  },
  data() {
    return {
      inputValue: '',
      isResultVisible: false,
      isLabelVisible: false,
      result: [],
      getResultsStatus: IDLE,
      abortRequests: true,
      limit: 10,
      offset: 1
    }
  },
  watch: {
    value: {
      handler(newValue) {
        if (newValue != null) {
          this.inputValue = newValue[this.optionTitleField]
        }
      },
      immediate: true
    }
  },
  created() {
    this.debounceInputChanged = debounce((value) => {
      return this.inputChanged(value)
    }, 0)
    this.getResult()
  },
  computed: {
    inputListeners() {
      var vm = this
      return Object.assign({}, this.$listeners, {
        input: (event) => {
          event.target.value != ''
            ? (this.isLabelVisible = true)
            : (this.isLabelVisible = false)
          this.debounceInputChanged(event.target.value)
        }
      })
    },
    hasResult() {
      return this.result && this.result.length > 0
    },
    computedWith() {
      return { width: `${this.width}%` }
    },
    ...apiStatusComputedFactory(['getResultsStatus']),
    hasError() {
      return this.error && this.error.$invalid && this.error.$dirty
    }
  },

  methods: {
    handleScroll() {
      this.offset += 1
      this.getResult()
    },
    onInputFocus() {
      this.isResultVisible = true
    },
    dynamicClass(className) {
      return [
        className,
        this.isLabelVisible ? `${className}--active` : '',
        this.hasError ? `${className}--error` : '',
        this.isResultOpen ? `${className}--open` : ''
      ]
    },
    optionSelected(item) {
      this.isResultVisible = false
      this.$emit('input', item)
    },
    open() {
      this.isResultVisible = true
    },
    close() {
      this.isResultVisible = false
    },
    async inputChanged(value) {
      if (value != this.inputValue) {
        this.setInitForm()
      }
      this.inputValue = value
      if (!!value) {
        this.open()

        await this.getResult(value)
      } else {
        this.setInitForm()
        this.close()
      }
    },
    async getResult(inputValue) {
      this.$options.abort?.()
      this.getResultsStatus = PENDING

      const { response, error } = await withAsync(
        this.searchMethod,
        {
          ...this.searchParams,
          search: inputValue,
          search_popup: inputValue,
          limit: this.limit,
          offset: this.offset
        },
        {
          abort: (abort) => (this.$options.abort = abort)
        }
      )
      if (error) {
        if (error.aborted) {
          this.getResultsStatus = PENDING
        } else {
          this.getResultsStatus = PENDING
        }
        return
      }

      this.getResultsStatus = SUCCESS
      let mergedResults = [
        ...this.result,
        ...(response.data.records || response.data.result || [])
      ]
      this.result = mergedResults
    },
    setInitForm() {
      this.result = []
    }
  }
}
</script>

<style lang="scss" scoped>
.input {
  position: relative;
  &__icon {
    position: absolute;
    top: 40%;
    right: 0;
    transition: all 0.4s;

    &--reverse {
      transform: rotate(180deg);
    }
  }
  &__wrapper {
    position: relative;
    display: flex;
    align-items: center;
    padding-bottom: 20px;
  }
  &__label {
    font-weight: 400;
    font-size: 16px;
    color: #2b2b2b;
    width: 100%;
    transition: all 0.3s;
  }
  &-controller {
    background: #ffffff;
    width: 100%;
    transition: all 0.3s;
    border: none;
    overflow: hidden;
    width: 100%;
    padding: 20px 20px 10px 0;
    font-weight: 400;
    font-size: 16px;
    color: #2b2b2b;
    display: block;

    &__wrapper {
      overflow: hidden;
      position: relative;
      width: 100%;
      border: none;
      border-bottom: 2px solid #2b2b2b;
      &--error {
        border: 2px solid #eb5757;
        color: #eb5757;
      }
    }
    // &--active {
    //   padding: 25px 16px 8px;
    // }

    &:focus,
    &:hover {
      outline: none;
    }
    &:disabled {
      background: #f2f2f2;
      border: none;
      border-radius: 8px;
    }
    &::placeholder {
      font-weight: 400;
      font-size: 16px;
      color: rgba(17, 17, 17, 0.5);
    }
  }
  &-result {
    &__list {
      position: absolute;
      top: 80px;
      width: 100%;
      background: var(--light);
      border-radius: 8px;
      overflow: hidden;
      overflow-y: auto;
      max-height: 250px;
      border: 1px solid #2b2b2b;
      z-index: 15;

      &::-webkit-scrollbar {
        width: 1px;
      }
      &::-webkit-scrollbar-track {
        box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.3);
      }
      &::-webkit-scrollbar-thumb {
        border-radius: 6px;
        background-color: darkgrey;
      }
      &--error {
        border: 1px solid #eb5757;
      }
    }
    &__item {
      padding: 16px;
      cursor: pointer;
      transition: all 0.3s;
      min-height: 40px;
      font-weight: 600;
      font-size: 14px;
      color: var(--dark);

      &--selected {
        background: var(--green);
        color: var(--light);
      }

      &:hover {
        background: var(--green-light);
        color: var(--light);
      }
    }
    &__static {
      display: flex;
      align-items: center;
      justify-content: center;
      cursor: initial;
      &:hover {
        color: var(--dark);
        background: var(--light);
      }
    }
  }
}
</style>
