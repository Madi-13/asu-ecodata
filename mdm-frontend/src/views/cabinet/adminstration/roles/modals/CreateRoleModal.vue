<template>
  <ModalBackdrop :name="modalName" title="Создание роли">
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

import { createRole } from '@/api/roleAPI'

export default {
  components: {
    ModalBackdrop
  },
  data() {
    return {
      saveStatus: IDLE,
      modalName: 'createRoleModal',
      form: {
        name: null
      }
    }
  },
  computed: {
    ...apiStatusComputedFactory(['saveStatus'])
  },
  methods: {
    async onSubmit() {
      this.saveStatus = PENDING
      const { error } = await withAsync(createRole, this.form)
      if (error) {
        this.saveStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.saveStatus = SUCCESS
      this.$toast.success('Роль успешно создан!')
      this.$modal.hide(this.modalName)
      this.$emit('onClose')
    }
  }
}
</script>

<style lang="scss" scoped></style>
