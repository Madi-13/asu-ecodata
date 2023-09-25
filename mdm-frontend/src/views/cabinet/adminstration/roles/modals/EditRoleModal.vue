<template>
  <ModalBackdrop
    :name="modalName"
    title="Редактирование роли"
    v-if="selectedRole"
  >
    <BaseInput
      label="Название роля"
      placeholder="Введите название"
      :width="100"
      v-model="form.name"
      name="name"
    ></BaseInput>
    <BaseButton
      color="green"
      @click="onSubmit"
      class="mt--3"
      width="30%"
      :loading="saveStatusPending"
      height="60px"
      >Сохранить</BaseButton
    >
  </ModalBackdrop>
</template>

<script>
import ModalBackdrop from '@/components/common/ModalBackdrop.vue'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

import { updateRole } from '@/api/roleAPI'

export default {
  components: {
    ModalBackdrop
  },
  props: {
    selectedRole: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      saveStatus: IDLE,
      modalName: 'editRoleModal',
      form: {
        id: null,
        name: null
      }
    }
  },
  watch: {
    selectedRole: {
      handler(newValue) {
        this.form = newValue
      }
    }
  },
  computed: {
    ...apiStatusComputedFactory(['saveStatus'])
  },
  methods: {
    async onSubmit() {
      this.saveStatus = PENDING
      const { error } = await withAsync(updateRole, this.form)
      if (error) {
        this.saveStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.saveStatus = SUCCESS
      this.$toast.success('Роль успешно обновлен!')
      this.$modal.hide(this.modalName)
      this.$emit('onClose')
    }
  }
}
</script>

<style lang="scss" scoped></style>
