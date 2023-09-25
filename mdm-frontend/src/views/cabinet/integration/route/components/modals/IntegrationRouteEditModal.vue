<template>
  <ModalBackdrop :name="modalName" title="Создание маршрута" v-if="selected">
    <BaseRadio
      :radioValues="routeTypes"
      v-model="form.property.property_type"
      v-can:update.disable="{
        resource: 'integration',
        permission: 'integration_operations'
      }"
      label="Тип маршрута"
    ></BaseRadio>
    <BaseInput
      label="Название сценария"
      v-can:update.disable="{
        resource: 'integration',
        permission: 'integration_operations'
      }"
      placeholder="Введите название"
      :width="100"
      v-model="form.property.property_name"
      name="name"
    ></BaseInput>
    <BaseInput
      label="Адрес"
      placeholder="Введите адрес"
      v-can:update.disable="{
        resource: 'integration',
        permission: 'integration_operations'
      }"
      :width="100"
      v-model="form.property.path"
      name="name"
    ></BaseInput>

    <BaseInput
      :label="
        form.property.property_type == 'KAFKA'
          ? 'Топик брокера сообщении'
          : 'Content-Type'
      "
      :placeholder="
        form.property.property_type == 'KAFKA'
          ? 'Введите название топика'
          : 'Введите Content-Type'
      "
      :width="100"
      v-can:update.disable="{
        resource: 'integration',
        permission: 'integration_operations'
      }"
      v-model="form.property.rq_token"
      name="name"
    ></BaseInput>

    <BaseInput
      label="Группа брокера сообщении"
      v-if="form.property.property_type == 'KAFKA'"
      placeholder="Введите адрес"
      :width="100"
      v-model="form.property.rq_group"
      v-can:update.disable="{
        resource: 'integration',
        permission: 'integration_operations'
      }"
      name="name"
    ></BaseInput>

    <BaseButton
      color="green"
      @click="onSubmit"
      class="mt--3"
      v-can:update.hide="{
        resource: 'integration',
        permission: 'integration_operations'
      }"
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

import { operateIntegrationRoute } from '@/api/integrationAPI'

export default {
  components: {
    ModalBackdrop
  },
  props: {
    selected: {
      type: Object
    }
  },
  data() {
    return {
      saveStatus: IDLE,
      modalName: 'editIntegrationRouteModal',
      activeComponentName: null,
      form: {
        operation_type: 'alter',
        property: {
          property_name: null,
          property_type: 'HTTP',
          path: null,
          rq_group: null,
          rq_token: null
        }
      },
      routeTypes: [
        {
          name: 'HTTP',
          value: 'HTTP',
          group: 'routeTypes'
        },
        {
          name: 'Брокер сообщении',
          value: 'KAFKA',
          group: 'routeTypes'
        }
      ]
    }
  },
  computed: {
    ...apiStatusComputedFactory(['saveStatus'])
  },
  watch: {
    selected: {
      handler(newValue) {
        this.form = {
          operation_type: 'alter',
          property: {
            id: newValue.id,
            property_name: newValue.property_name,
            property_type: newValue.property_type,
            path: newValue.path,
            rq_group: newValue.rq_group,
            rq_token: newValue.rq_token
          }
        }
      },
      immediate: true
    }
  },
  methods: {
    onActivateComponent(name) {
      this.activeComponentName = name
    },
    async onSubmit() {
      this.saveStatus = PENDING
      const { error } = await withAsync(operateIntegrationRoute, this.form)
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
