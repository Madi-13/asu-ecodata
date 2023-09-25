<template>
  <ModalBackdrop :name="modalName" title="Создание сценария">
    <BaseInput
      label="Название сценария"
      placeholder="Введите название"
      :width="100"
      v-model="form.scenario.scenario_name"
      name="name"
      class=""
      :has-validation="false"
    ></BaseInput>
    <BaseButton
      color="green"
      @click="onSubmit"
      class="mt--3"
      width="30%"
      v-can:write.hide="{
        resource: 'integration',
        permission: 'integration_operations'
      }"
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

import { operateScenario } from '@/api/integrationAPI'

export default {
  components: {
    ModalBackdrop
  },
  data() {
    return {
      saveStatus: IDLE,
      modalName: 'createScenarioModal',
      form: {
        operation_type: 'add',
        scenario: {
          scenario_name: null,
          status: 'ACTIVE'
        }
      }
    }
  },
  computed: {
    ...apiStatusComputedFactory(['saveStatus'])
  },
  methods: {
    async onSubmit() {
      this.saveStatus = PENDING
      const { error } = await withAsync(operateScenario, this.form)
      if (error) {
        this.saveStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.saveStatus = SUCCESS
      this.$modal.hide(this.modalName)
      this.$emit('onClose')
    }
  }
}
</script>

<style lang="scss" scoped></style>
