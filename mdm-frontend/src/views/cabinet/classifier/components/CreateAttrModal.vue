<template>
  <ModalBackdrop name="createAttrModal" title="Создание аттрибутов">
    <BaseInput
      label="Название аттрибута"
      placeholder="Введите название"
      :width="100"
      v-model="form.name"
      name="name"
      class=""
      :has-validation="false"
    ></BaseInput>
    <div class="row mt--2">
      <BaseInput
        label="Значение"
        placeholder="Введите значение"
        :has-validation="false"
        :width="100"
        v-model="newAttr.value"
        name="value"
        class=""
      ></BaseInput>
      <BaseButton
        color="green"
        class="operations__btn"
        height="45px"
        :disabled="$v.newAttr.$invalid"
        @click="addNewValue"
        width="100%"
        >Добавить</BaseButton
      >
    </div>
    <div class="chips">
      <Chips
        v-for="(chip, idx) in form.values"
        :key="chip.title"
        @remove="onRemove(idx)"
        :title="chip.value"
      />
    </div>
    <BaseButton
      color="green"
      @click="onsubmit"
      class="mt--3"
      width="30%"
      :loading="onSubmitStatusPending"
      height="60px"
      >Сохранить</BaseButton
    >
  </ModalBackdrop>
</template>

<script>
import ModalBackdrop from '@/components/common/ModalBackdrop.vue'
import Chips from '@/components/common/Chips.vue'

import { required } from 'vuelidate/lib/validators'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus
import { operateAttrs } from '@/api/attsAPI'

export default {
  components: {
    ModalBackdrop,
    Chips
  },
  props: {
    classifierId: {
      type: Number
    }
  },
  data() {
    return {
      onSubmitStatus: IDLE,
      newAttr: {
        key: null,
        value: null,
        status: 'active'
      },
      form: {
        operation_type: 'add',
        attribute_id: null,
        name: null,
        values: []
      }
    }
  },
  validations: {
    newAttr: {
      value: {
        required
      }
    }
  },
  computed: {
    ...apiStatusComputedFactory(['onSubmitStatus']),
    getCurrentTableId() {
      return this.$route.query.table_id
    }
  },
  methods: {
    addNewValue() {
      this.form.values.push(this.newAttr)
      this.newAttr = {
        key: null,
        value: null,
        status: 'active'
      }
    },
    onRemove(idx) {
      this.form.values.splice(idx, 1)
    },
    async onsubmit() {
      this.onSubmitStatus = PENDING
      this.form['attribute_dict_id'] = +this.getCurrentTableId
      this.form['dictionary_row_id'] = +this.classifierId
      const { error } = await withAsync(operateAttrs, this.form)
      if (error) {
        this.onSubmitStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.onSubmitStatus = SUCCESS
      this.$toast.success('Аттрибуты успешно сохранены')
      this.form = {
        operation_type: 'add',
        attribute_id: null,
        name: null,
        values: []
      }
      this.$modal.hide('createAttrModal')
    }
  }
}
</script>

<style lang="scss" scoped>
.row {
  display: grid;
  grid-template-columns: 3fr 1fr;
  align-items: flex-end;
  gap: 2rem;
}
.chips {
  margin-top: 2rem;
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
}
</style>
