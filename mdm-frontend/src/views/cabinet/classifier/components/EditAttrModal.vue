<template>
  <ModalBackdrop
    name="editAttrModal"
    title="Редактирование аттрибутов"
    @onOpen="fetchAttr"
  >
    <tempalte v-if="form">
      <div
        class="attr__form"
        v-for="(formItem, formIdx) in form"
        :key="formItem.name"
      >
        <inline-svg
          :src="require(`@/assets/icons/close.svg`)"
          @click="removeAttr(formIdx)"
          class="attr__form-icon"
        />
        <BaseInput
          label="Название аттрибута"
          placeholder="Введите название"
          :width="100"
          v-model="formItem.name"
          :name="formItem.name"
          :has-validation="false"
        ></BaseInput>
        <div class="row mt--2">
          <BaseInput
            label="Значение"
            placeholder="Введите значение"
            :has-validation="false"
            :width="100"
            v-model="formItem.value"
            :name="`${formItem.name}-value`"
          ></BaseInput>
          <BaseButton
            color="green"
            class="operations__btn"
            height="45px"
            @click="addNewValue(formIdx)"
            :disabled="$v.form.$each[formIdx].value.$invalid"
            width="100%"
            >Добавить</BaseButton
          >
        </div>
        <div class="chips">
          <Chips
            v-for="(chip, idx) in formItem.values"
            :key="chip.title"
            @remove="onRemoveOption(formItem, idx)"
            :title="chip.value"
          />
        </div>
      </div>

      <div class="btns">
        <BaseButton
          color="green"
          @click="onsubmit"
          class="mt--3"
          width="30%"
          :loading="onSubmitStatusPending"
          height="50px"
          >Сохранить</BaseButton
        >
      </div>
    </tempalte>
    <template v-else>
      <p class="paragraph--lg">Нет созданных аттрибутов</p>
      <BaseButton
        color="green"
        @click="openCreateModal"
        class="mt--3"
        width="50%"
        height="50px"
        >Создать аттрибут</BaseButton
      >
    </template>
  </ModalBackdrop>
</template>

<script>
import ModalBackdrop from '@/components/common/ModalBackdrop.vue'
import Chips from '@/components/common/Chips.vue'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus
import { operateAttrs, getAttrs } from '@/api/attsAPI'

import { required } from 'vuelidate/lib/validators'

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
      fetchStatus: IDLE,
      deleteStatus: IDLE,
      newAttr: {
        key: null,
        value: null,
        status: 'active'
      },
      form: null
    }
  },
  validations: {
    form: {
      $each: {
        value: { required }
      }
    }
  },
  computed: {
    ...apiStatusComputedFactory([
      'onSubmitStatus',
      'fetchStatus',
      'deleteStatus'
    ]),
    getCurrentTableId() {
      return +this.$route.query.table_id
    }
  },
  methods: {
    openCreateModal() {
      this.$modal.hide('editAttrModal')
      this.$modal.show('createAttrModal')
    },
    addNewValue(formItem) {
      this.form[formItem].values.push({
        key: null,
        value: this.form[formItem].value,
        status: 'active'
      })
      this.formItem.value = null
    },
    onRemoveOption(formIdx, rowIdx) {
      this.form[formIdx].values.splice(rowIdx, 1)
    },
    async removeAttr(formIdx) {
      this.deleteStatus = PENDING
      this.form[formIdx]['operation_type'] = 'drop'
      this.form[formIdx]['attribute_dict_id'] = +this.getCurrentTableId
      this.form[formIdx]['dictionary_row_id'] = +this.classifierId
      const { error } = await withAsync(operateAttrs, this.form[formIdx])
      if (error) {
        this.deleteStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.deleteStatus = SUCCESS
      this.$toast.success('Атрибуты успешно удалены')
      this.form.splice(formIdx, 1)
    },
    async fetchAttr() {
      this.fetchStatus = PENDING
      const { response, error } = await withAsync(getAttrs, {
        dictionary_id: this.getCurrentTableId,
        dictionary_row_id: this.classifierId
      })
      if (error) {
        this.fetchStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      let result = response.data.result
      if (result.length > 0) {
        this.form = result.map((item) => ({
          operation_type: 'alter',
          attribute_id: item.attrib_id,
          value: null,
          name: item.name,
          values: item.vals && item.vals.filter((val) => !!val.value)
        }))
      } else {
        this.form = {
          operation_type: 'add',
          attribute_id: null,
          name: null,
          values: []
        }
      }
      this.fetchStatus = SUCCESS
    },
    async onsubmit() {
      this.form.forEach(async (formItem, formIdx) => {
        await this.saveAttr(formIdx)
      })
    },
    async saveAttr(formIdx) {
      this.onSubmitStatus = PENDING
      this.form[formIdx]['attribute_dict_id'] = +this.getCurrentTableId
      this.form[formIdx]['dictionary_row_id'] = +this.classifierId
      const { error } = await withAsync(operateAttrs, this.form[formIdx])
      if (error) {
        this.onSubmitStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.onSubmitStatus = SUCCESS
      this.$modal.hide('editAttrModal')
    }
  }
}
</script>

<style lang="scss" scoped>
.btns {
  display: flex;
  align-items: center;
  gap: 2rem;
}
.row {
  display: grid;
  grid-template-columns: 3fr 1fr;
  align-items: flex-end;
  gap: 2rem;
}
.chips {
  margin-top: 0.3rem;
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
  margin-top: 2rem;
}

.attr {
  &__form {
    display: flex;
    flex-direction: column;
    border: 1px dashed var(--green);
    padding: 10px;
    border-radius: 8px;
    position: relative;

    &:not(:first-child) {
      margin-top: 4rem;
    }

    &-icon {
      position: absolute;
      right: 8px;
      cursor: pointer;
      top: 5px;
      z-index: 234;
    }
  }
}
</style>
