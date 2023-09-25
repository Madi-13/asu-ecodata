<template>
  <ModalBackdrop :name="modalName" title="Смена статуса" v-if="selected">
    <BaseInput
      label="Причина смены статуса"
      :placeholder="`Введите причину ${
        isBlocking ? 'разблокировки' : 'блокировки'
      }`"
      :width="100"
      v-model="reasonValue"
      name="name"
      type="textarea"
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

import { setStatusRef } from '@/api/referenceRecordsAPI'

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
      modalName: 'blockReasonModal',
      reasonValue: null,
      form: null,
      selectedRow: null
    }
  },
  computed: {
    ...apiStatusComputedFactory(['saveStatus']),
    getCurrentRefId() {
      return this.$route.query.table_id
    },
    isBlocking() {
      return this.selected.row_status_code == 'BLOCK'
    }
  },
  watch: {
    selected: {
      handler(newValue) {
        this.selectedRow = newValue
      },
      immediate: true
    }
  },
  methods: {
    async onSubmit() {
      this.saveStatus = PENDING
      const { error } = await withAsync(setStatusRef, {
        dictionary_id: this.getCurrentRefId,
        row_id: this.selectedRow.id,
        new_status:
          this.selectedRow.row_status_code == 'ACTIVE' ? 'BLOCKED' : 'ACTIVE',
        change_reason: this.reasonValue
      })
      if (error) {
        this.saveStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.saveStatus = SUCCESS
      this.$toast.success(
        `Запись с идентификатором ${this.selectedRow.id} успешно ${
          this.isBlocking ? 'разблокирован' : 'заблокирован'
        }!`
      )
      this.reasonValue = null
      this.$emit('onClose')
    }
  }
}
</script>

<style lang="scss" scoped></style>
