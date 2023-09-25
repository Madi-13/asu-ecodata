<template>
  <ModalBackdrop
    :name="modalName"
    title="Редактирование пользователя"
    v-if="selectedUser && this.form.roles"
  >
    <form class="form">
      <BaseInput
        label="Имя пользователя"
        placeholder="Введите имя"
        :width="100"
        :error="$v.form.firstname"
        @blur="$v.form.firstname.$touch()"
        v-model="form.firstname"
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
        disabled
        v-model="form.username"
        name="name"
      ></BaseInput>
      <BaseDropdown
        label="Роли"
        placeholder="Выберите из списка"
        v-model="form.roles"
        optionValueField="id"
        optionTitleField="name"
        name="roles"
        hasMultiSelection
        :options="roles"
        :width="100"
      />
      <BaseInput
        label="Пароль пользователя"
        placeholder="Введите пароль"
        :width="100"
        :error="$v.form.password"
        @blur="$v.form.password.$touch()"
        type="password"
        v-model="form.password"
        name="name"
      ></BaseInput>
      <BaseInput
        label="Повтор пароля пользователя"
        placeholder="Введите пароль повторно"
        type="password"
        :width="100"
        v-model="form.password_rec"
        :error="$v.form.password_rec"
        @blur="$v.form.password_rec.$touch()"
        name="name"
      ></BaseInput>
    </form>
    <BaseButton
      color="green"
      @click.prevent="onSubmit"
      class="mt--3"
      width="30%"
      :disabled="($v.form.$invalid && $v.form.$dirty) || saveStatusPending"
      :loading="saveStatusPending"
      v-can:update.hide="{
        resource: 'adminstration',
        permission: 'user_operations'
      }"
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

import { serializeUser, deserializeUser } from '@/api/serializers/users'
import { updateUser } from '@/api/userAPI'
import { getRoles, createRole, deleteRole, getUserRoles } from '@/api/roleAPI'
import { mapGetters } from 'vuex'

export default {
  components: {
    ModalBackdrop
  },
  props: {
    selectedUser: {
      type: Object
    }
  },
  data() {
    return {
      saveStatus: IDLE,
      modalName: 'editUserModal',
      roles: [],
      form: {
        firstname: null,
        lastname: null,
        username: null,
        password: null,
        roles: [],
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
  watch: {
    selectedUser: {
      async handler(newValue) {
        if (!newValue) return
        this.form = serializeUser(newValue)
        await this.fetchUserRoles(this.form.id)
      },
      immediate: true
    },
    'form.roles': {
      handler(newValue = [], oldValue = []) {
        let rolesToDelete = oldValue.filter(
          (o1) => newValue.map((o2) => o2.id).indexOf(o1.id) === -1
        )

        let rolesToSave = newValue.filter(
          (o1) => oldValue.map((o2) => o2.id).indexOf(o1.id) === -1
        )

        if (rolesToDelete.length > 0) {
          this.removeRole(rolesToDelete)
        }
        if (rolesToSave.length > 0) {
          this.saveRole()
        }
      }
    },
    immediate: false
  },
  created() {
    this.fetchRoles()
  },
  methods: {
    async saveRole() {
      const { error } = await withAsync(createRole, this.form.roles, {
        headers: {
          'user-id': this.form.id,
          'client-id': this.getRealmClientId
        }
      })
      if (error) {
        this.$toast.warning('Не удалось сохранить роль пользавателя')
      }
    },
    removeRole(roles) {
      withAsync(deleteRole, {
        headers: {
          'user-id': this.form.id,
          'client-id': this.getRealmClientId
        },
        data: roles
      })
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
    async fetchUserRoles(userId) {
      const { response, error } = await withAsync(getUserRoles, {
        headers: {
          'user-id': userId
        }
      })
      if (error) {
        return []
      }
      let userRoles = response ? response.data : []
      this.$set(this.form, 'roles', userRoles)
    },
    async onSubmit() {
      this.$v.form.$touch()
      if (!this.$v.form.$invalid) {
        this.saveStatus = PENDING
        const { error } = await withAsync(
          updateUser,
          deserializeUser(this.form)
        )
        if (error) {
          this.saveStatus = ERROR
          this.$toast.error(error.message)
          return
        }
        this.saveStatus = SUCCESS
        this.$toast.success('Пользователь успешно обновлен!')
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
  grid-template-columns: repeat(2, 1fr);
  gap: 10px;
}
</style>
