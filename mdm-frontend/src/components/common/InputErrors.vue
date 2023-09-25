<template>
  <div class="error__list" v-if="hasError">
    <p
      class="error__item"
      v-if="isContainsError('required') && !error.required && error.$dirty"
    >
      Поле обязательное
    </p>
    <p
      class="error__item"
      v-else-if="isContainsError('sameAsPassword') && !error.sameAsPassword"
    >
      Пароли не совпадают
    </p>
    <p class="error__item" v-else-if="isContainsError('email') && !error.email">
      Невалидный почтовый адрес
    </p>
    <p
      class="error__item"
      v-else-if="isContainsError('minLength') && !error.minLength"
    >
      Вы ввели недостаточное количество символов
    </p>
  </div>
</template>

<script>
export default {
  props: {
    errorProp: {
      type: Object,
      default: () => {}
    }
  },
  computed: {
    hasError() {
      return this.error && this.error.$invalid && this.error.$dirty
    }
  },
  data() {
    return {
      error: null
    }
  },
  watch: {
    errorProp: {
      handler(newValue) {
        this.error = newValue
      },
      immediate: true
    }
  },
  methods: {
    isContainsError(errName) {
      const result = errName in this.error
      return result
    }
  }
}
</script>

<style lang="scss" scoped>
.error {
  &__list {
    position: absolute;
    top: 76%;
    left: 0;
    height: 16px;
    transition: all 0.3s;
    width: 100%;

    &--active {
      top: 35%;
    }
  }
  &__item {
    font-size: 12px;
    color: #db524e;
    padding-top: 7px;
  }
}
</style>
