<template>
  <ModalBackdrop :name="modalName" title="Добавление атрибута">
    <BaseInput
      label="Название атрибута"
      placeholder="ведите название"
      :width="100"
      v-model="newAttrValue"
      name="name"
      :has-validation="false"
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

import { createAttributeValues } from '@/api/attsAPI'

export default {
  components: {
    ModalBackdrop
  },
  props: {
    attributeId: {
      type: Number
    }
  },
  data() {
    return {
      saveStatus: IDLE,
      modalName: 'createAttrModal',
      newAttrValue: null
    }
  },
  computed: {
    ...apiStatusComputedFactory(['saveStatus']),
    getCurrentRefId() {
      return this.$route.query.table_id
    }
  },
  methods: {
    async onSubmit() {
      const { response, error } = await withAsync(createAttributeValues, {
        attrib_id: this.attributeId,
        value: this.newAttrValue
      })
      if (error) {
        this.saveStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.saveStatus = SUCCESS
      this.$emit('onClose', {
        key: response.data.out_key,
        value: this.newAttrValue
      })
      this.$modal.hide(this.modalName)
    }
  }
}
</script>

<style lang="scss" scoped></style>
