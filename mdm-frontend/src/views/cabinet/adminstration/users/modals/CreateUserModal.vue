<template>
  <ModalBackdrop :name="modalName" title="Создание пользователя">
    <form class="form">
      <BaseInput
        label="Имя пользователя"
        placeholder="Введите имя"
        :width="100"
        v-model="form.firstname"
        :error="$v.form.firstname"
        @blur="$v.form.firstname.$touch()"
        name="name"
      ></BaseInput>
      <BaseInput
        label="Фамилия пользователя"
        placeholder="Введите фамилию"
        :width="100"
        v-model="form.lastname"
        name="name"
      ></BaseInput>
      <BaseInput
        label="Логин пользователя"
        placeholder="Введите логин"
        :width="100"
        :error="$v.form.username"
        @blur="$v.form.username.$touch()"
        v-model="form.username"
        name="name"
      ></BaseInput>
      <BaseDropdown
        label="Роли"
        placeholder="Выберите из списка"
        v-model="form.roles"
        optionValueField="id"
        optionTitleField="name"
        hasMultiSelection
        :options="roles"
        :width="100"
      />
      <BaseInput
        label="Пароль пользователя"
        placeholder="Введите пароль"
        :width="100"
        type="password"
        v-model="form.password"
        :error="$v.form.password"
        @blur="$v.form.password.$touch()"
        name="name"
      ></BaseInput>
      <BaseInput
        label="Повтор пароля пользователя"
        placeholder="Введите пароль повторно"
        type="password"
        :error="$v.form.password_rec"
        @blur="$v.form.password_rec.$touch()"
        :width="100"
        v-model="form.password_rec"
        name="name"
      ></BaseInput>
    </form>
    <BaseButton
      color="green"
      @click.prevent="onSubmit"
      :disabled="($v.form.$invalid && $v.form.$dirty) || saveStatusPending"
      class="mt--3"
      width="30%"
      :loading="saveStatusPending"
      height="60px"
      >Сохранить</BaseButton
    >
  </ModalBackdrop>
</template>

<script>
import { required, minLength, sameAs } from 'vuelidate/lib/validators'

import ModalBackdrop from '@/components/common/ModalBackdrop.vue'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

import { getRoles, createRole } from '@/api/roleAPI'

import { createUser } from '@/api/userAPI'
import { operationsClassifierTable } from '@/fixtures/options'
import { deserializeUser } from '@/api/serializers/users'
import { mapGetters } from 'vuex'

export default {
  components: {
    ModalBackdrop
  },
  data() {
    return {
      operationsClassifierTable,
      saveStatus: IDLE,
      modalName: 'createUserModal',
      roles: null,
      form: {
        firstname: null,
        lastname: null,
        roles: [],
        username: null,
        password: null,
        password_rec: null
      }
    }
  },
  validations: {
    form: {
      firstname: {
        required
      },
      username: {
        required
      },
      password: {
        required
      },
      password_rec: {
        required,
        sameAsPassword: sameAs(function () {
          return this.form.password
        }),
        minLength: minLength(8)
      }
    }
  },
  computed: {
    ...mapGetters(['getRealmClientId']),
    ...apiStatusComputedFactory(['saveStatus'])
  },
  created() {
    this.fetchRoles()
  },
  methods: {
    async saveRole(userId) {
      const { error } = await withAsync(createRole, this.form.roles, {
        headers: { 'user-id': userId }
      })
      if (error) {
        this.$toast.warning('Не удалось сохранить роль пользавателя')
      }
    },

    async fetchRoles() {
      const { response, error } = await withAsync(getRoles, {
        headers: {
          'client-id': this.getRealmClientId
        }
      })
      if (error) {
        return
      }
      this.roles = response.data
    },
    async onSubmit() {
      this.$v.form.$touch()
      if (!this.$v.form.$invalid) {
        this.saveStatus = PENDING
        const { response, error } = await withAsync(
          createUser,
          deserializeUser(this.form)
        )
        if (error) {
          this.saveStatus = ERROR
          this.$toast.error(error.message)
          return
        }
        await this.saveRole(response.data['user-id'])
        this.saveStatus = SUCCESS
        this.$toast.success('Пользователь успешно создан!')
        this.$modal.hide(this.modalName)
        this.$emit('onClose')
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.form {
  display: grid;
  grid-template-columns: repeat(2, 50%);
  gap: 10px;
}
</style>
