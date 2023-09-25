<template>
  <form class="auth__form">
    <BaseInput
      label="Логин"
      placeholder="Введите логин"
      :width="100"
      v-model="form.username"
      name="username"
      class=""
    ></BaseInput>
    <BaseInput
      label="Пароль"
      placeholder="Введите пароль"
      :width="100"
      v-model="form.password"
      name="password"
      type="password"
    ></BaseInput>
    <BaseButton
      color="green"
      class="mt--5"
      @click.prevent="onSubmit"
      :disabled="$v.form.$invalid || isLoading"
      :loading="isLoading"
      >Вход</BaseButton
    >
  </form>
</template>

<script>
import { required } from 'vuelidate/lib/validators'
import { mapActions, mapGetters } from 'vuex'
import { apiStatus } from '@/api/constants/apiStatus'
const { PENDING, SUCCESS, ERROR } = apiStatus

export default {
  data() {
    return {
      form: {
        username: null,
        password: null
      },
      isDisabled: false,
      isLoading: false
    }
  },
  computed: {
    ...mapGetters(['getLoginStatus', 'getTokens'])
  },
  validations: {
    form: {
      username: {
        required
      },
      password: {
        required
      }
    }
  },
  methods: {
    ...mapActions({ loginAction: 'userLogin' }),
    async onSubmit() {
      this.$v.form.$touch()
      if (!this.$v.form.invalid) {
        this.isLoading = true
        await this.loginAction(this.form)
        if (this.getLoginStatus === SUCCESS) {
          this.isLoading = false
          this.$toast.success('Вы успешно вошли в систему')
          this.$router.push({ name: 'Analytics' })
          return
        }
        this.isLoading = false
        this.$toast.error('Неправильный логин или пароль')
      }
    }
  }
}
</script>

<style lang="scss" scoped></style>
