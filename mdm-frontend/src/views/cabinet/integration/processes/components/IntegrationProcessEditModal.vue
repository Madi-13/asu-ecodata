<template>
  <ModalBackdrop
    :name="modalName"
    title="Редактирование процесса"
    v-if="selected"
  >
    <BaseInput
      label="Название процесса"
      placeholder="Введите название"
      :width="100"
      v-model="form.process.process_name"
      v-can:update.disable="{
        resource: 'integration',
        permission: 'integration_operations'
      }"
      name="name"
    ></BaseInput>
    <InputAutocomplete
      name="dictionary"
      placeholder="Введите значение"
      label="Справочник"
      :searchMethod="getSystemRefs"
      optionTitleField="table_name"
      v-can:update.disable="{
        resource: 'integration',
        permission: 'integration_operations'
      }"
      optionValueField="id"
      v-model="form.process.dictionary"
    />
    <InputAutocomplete
      name="scenario"
      placeholder="Введите значение"
      label="Сценарий"
      :searchMethod="getScenarios"
      v-can:update.disable="{
        resource: 'integration',
        permission: 'integration_operations'
      }"
      optionTitleField="scenario_name"
      optionValueField="id"
      v-model="form.process.scenario"
    />
    <InputAutocomplete
      name="property"
      placeholder="Введите значение"
      label="Маршрут"
      :searchMethod="getIntegrationRoutes"
      v-can:update.disable="{
        resource: 'integration',
        permission: 'integration_operations'
      }"
      optionTitleField="property_name"
      optionValueField="id"
      v-model="form.process.property"
    />
    <BaseButton
      color="green"
      @click="onSubmit"
      class="mt--3"
      width="30%"
      :loading="saveStatusPending"
      v-can:update.hide="{
        resource: 'integration',
        permission: 'integration_operations'
      }"
      height="60px"
      >Сохранить</BaseButton
    >
  </ModalBackdrop>
</template>

<script>
import ModalBackdrop from '@/components/common/ModalBackdrop.vue'
import InputAutocomplete from '@/components/common/InputAutocomplete.vue'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

import {
  getIntegrationRoutes,
  getScenarios,
  operateProcess
} from '@/api/integrationAPI'

import { getSystemRefs } from '@/api/objectsAPI'

import {
  processBuildForm,
  processDeserialize
} from '@/api/serializers/intergrationForm'

export default {
  components: {
    ModalBackdrop,
    InputAutocomplete
  },
  props: {
    selected: {
      type: Object
    }
  },
  data() {
    return {
      saveStatus: IDLE,
      modalName: 'editIntegrationProcessModal',
      activeComponentName: null,
      form: {
        operation_type: 'add',
        process: null
      }
    }
  },
  computed: {
    ...apiStatusComputedFactory(['saveStatus'])
  },
  watch: {
    selected: {
      handler(newValue) {
        if (newValue) {
          this.form.process = processBuildForm(newValue)
        }
      },
      immediate: true
    }
  },
  methods: {
    getSystemRefs,
    getScenarios,
    getIntegrationRoutes,
    onActivateComponent(name) {
      this.activeComponentName = name
    },
    async onSubmit() {
      this.saveStatus = PENDING
      const { error } = await withAsync(
        operateProcess,
        processDeserialize(this.form)
      )
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
