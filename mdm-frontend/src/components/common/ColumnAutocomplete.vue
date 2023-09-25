<template>
  <div
    :class="['input__wrapper', hasError && 'input__wrapper--error']"
    :style="computedWith"
    v-click-outside="close"
  >
    <div
      :class="dynamicClass('input-controller__wrapper')"
      @click="setCoordinates"
    >
      <input
        v-on="inputListeners"
        v-bind="{ ...$attrs, disabled }"
        type="text"
        :value="inputValue"
        autocomplete="off"
        :id="name"
        :class="dynamicClass('input-controller')"
      />
    </div>
    <div
      :class="dynamicClass('input-result__list')"
      v-if="isResultVisible"
      :style="{
        left: `${listLeftPosition}px`,
        top: `${listTopPosition}px`,
        width: `${listWidth}px`
      }"
    >
      <div
        class="input-result__item"
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
        v-else-if="getResultsStatusSuccess && !hasResult"
      >
        Записей не найдено
      </div>
    </div>
  </div>
</template>

<script>
import { getSystemRefs } from '@/api/objectsAPI'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
import debounce from 'lodash-es/debounce'

const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

export default {
  inheritAttrs: false,
  props: {
    name: {
      type: String,
      required: true
    },
    controllerWidth: {
      type: Number,
      default: 100
    },
    optionValueField: {
      type: String,
      default: 'name'
    },
    optionTitleField: {
      type: String,
      default: 'name_ru'
    },
    width: {
      type: Number,
      default: 100
    },
    value: {
      type: Object
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
      isOpen: false,
      abortRequests: true,
      listTopPosition: null,
      listLeftPosition: null,
      listWidth: null
    }
  },
  watch: {
    value: {
      handler(newValue) {
        this.inputValue = !newValue ? newValue : newValue[this.optionTitleField]
      },
      immediate: true
    }
  },
  created() {
    this.debounceInputChanged = debounce((value) => {
      return this.inputChanged(value)
    }, 0)
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
    hasError() {
      return this.error && this.error.$invalid && this.error.$dirty
    },
    hasResult() {
      return this.result.length > 0
    },
    computedWith() {
      return { width: `${this.width}%` }
    },
    ...apiStatusComputedFactory(['getResultsStatus'])
  },

  methods: {
    dynamicClass(className) {
      return [
        className,
        this.isLabelVisible ? `${className}--active` : '',
        this.isResultOpen ? `${className}--open` : ''
      ]
    },
    optionSelected(item) {
      this.isResultVisible = false
      this.$emit('input', {
        [this.optionTitleField]: item[[this.optionTitleField]],
        [this.optionValueField]: item[[this.optionValueField]]
      })
    },
    open() {
      this.isResultVisible = true
    },
    setCoordinates(event) {
      let controllerPosition = event.target.getBoundingClientRect()
      this.listTopPosition = controllerPosition.y + controllerPosition.height
      this.listLeftPosition = controllerPosition.x
      this.listWidth = controllerPosition.width
    },
    close() {
      this.isResultVisible = false
    },
    async inputChanged(value) {
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
      if (this.value != inputValue) {
        this.setInitForm()
      }

      const { response, error } = await withAsync(
        getSystemRefs,
        {
          search: inputValue,
          limit: 10,
          offset: 1
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
      this.result = response.data.result
    },
    setInitForm() {
      this.result = []
    }
  }
}
</script>

<style lang="scss" scoped>
.input {
  &__wrapper {
    position: relative;
    display: flex;
    align-items: center;
    &--error {
      border: 1px dashed red;
    }
  }
  &-controller {
    background: #ffffff;
    width: 100%;
    transition: all 0.3s;
    border: none;
    color: #111111;
    overflow: hidden;
    font-size: 16px;
    width: 100%;
    display: block;
    &__wrapper {
      overflow: hidden;
      position: relative;
      width: 100%;
      &--error {
        border: 1px solid #eb5757;
        color: #eb5757;
      }
    }
    &:focus,
    &:hover {
      outline: none;
    }
    &:disabled {
      border: none;
      cursor: not-allowed;
    }
    &::placeholder {
      font-weight: 600;
      font-size: 16px;
      color: rgba(17, 17, 17, 0.5);
    }
  }
  &-result {
    &__list {
      position: fixed;
      width: 100%;
      background: var(--light);
      border: 1px solid rgba(17, 17, 17, 0.32);
      border-radius: 0 0 8px 8px;
      overflow: hidden;
      overflow-y: auto;
      max-height: 250px;
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
      &:hover {
        background: var(--green);
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
