<template>
  <div>
    <form>
      <div class="reference__form">
        <template v-for="item in getReferenceControllers" :id="item.title">
          <template v-if="item.name != 'id'">
            <BaseInput
              v-if="item.type == 'input-text'"
              :label="item.title"
              :disabled="!item.allow_edit"
              placeholder="Введите значение"
              :width="100"
              v-model="item[item.name]"
              :name="item.name"
              class=""
            ></BaseInput>
            <BaseInput
              v-if="item.type == 'input-number'"
              :label="item.title"
              placeholder="Введите значение"
              :width="100"
              :disabled="!item.allow_edit || item.name == 'id'"
              mask="#######################"
              v-model="item[item.name]"
              :name="item.name"
              class=""
            ></BaseInput>
            <BaseInput
              v-if="item.type == 'input-money'"
              :label="item.title"
              placeholder="Введите значение"
              :width="100"
              :disabled="!item.allow_edit"
              mask="#######################"
              v-model="item[item.name]"
              :name="item.name"
              class=""
            ></BaseInput>
            <InputAutocomplete
              :name="item.name"
              v-if="item.type == 'autocomplete'"
              placeholder="Введите значение"
              :disabled="!item.allow_edit"
              :label="item.title"
              :searchMethod="getReferenceRecordsByName"
              :searchParams="{ table_name: item.ref_table_name }"
              v-model="item[item.name]"
            />
            <InputAutocomplete
              :name="item.name"
              v-if="item.type == 'tree'"
              :disabled="!item.allow_edit"
              placeholder="Введите значение"
              :label="item.title"
              :searchMethod="getReferenceRecordsByName"
              :searchParams="{ table_name: getCurrentObjectName }"
              v-model="item[item.name]"
            />
          </template>
        </template>
      </div>

      <p
        class="title--tertiary text--dark"
        v-if="getClassifierControllers.length > 0"
      >
        Классификация
      </p>
      <div class="reference__form mt--2">
        <template v-for="item in getClassifierControllers" :id="item.title">
          <template v-if="item.name != 'id'">
            <BaseInput
              v-if="item.type == 'input-text'"
              :label="item.title"
              placeholder="Введите значение"
              :width="100"
              :disabled="!item.allow_edit"
              v-model="item[item.name]"
              s
              :name="item.name"
              class=""
            ></BaseInput>
            <BaseInput
              v-if="item.type == 'input-number'"
              :label="item.title"
              placeholder="Введите значение"
              :disabled="!item.allow_edit || item.name == 'id'"
              :width="100"
              mask="#######################"
              v-model="item[item.name]"
              :name="item.name"
              class=""
            ></BaseInput>
            <BaseInput
              v-if="item.type == 'input-money'"
              :label="item.title"
              placeholder="Введите значение"
              :disabled="!item.allow_edit"
              :width="100"
              mask="#######################"
              v-model="item[item.name]"
              :name="item.name"
              class=""
            ></BaseInput>
            <InputAutocomplete
              :name="item.name"
              v-if="item.type == 'autocomplete'"
              placeholder="Введите значение"
              :disabled="!item.allow_edit"
              :label="item.title"
              :searchMethod="getReferenceRecordsByName"
              :searchParams="{ table_name: item.ref_table_name }"
              v-model="item[item.name]"
            />
            <InputAutocomplete
              :name="item.name"
              v-if="item.type == 'tree'"
              :disabled="!item.allow_edit"
              placeholder="Введите значение"
              :label="item.title"
              :searchMethod="getReferenceRecordsByName"
              :searchParams="{ table_name: getCurrentObjectName }"
              v-model="item[item.name]"
            />
          </template>
        </template>
      </div>
    </form>
    <div>
      <Characters
        v-for="(attr, idx) of attributeValues"
        :attribute="attr"
        @addedNewAttr="onNewAttrAdded(idx, $event)"
      />
    </div>
    <div class="reference__controllers mt--3">
      <BaseButton
        color="green"
        @click.prevent="onSubmit"
        :loading="updateRecordStatusPending"
        width="20rem"
        >Сохранить
      </BaseButton>
      <BaseButton color="outline-grey" width="20rem">Отмена</BaseButton>
    </div>
  </div>
</template>

<script>
import InputAutocomplete from '@/components/common/InputAutocomplete.vue'
import {
  createNewRecord,
  getReferenceFormForEdit,
  getReferenceRecordsByName,
  updateRecord
} from '@/api/referenceRecordsAPI'
import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'

const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

import {
  deserializeRecordForm,
  serializeRecordForms
} from '@/api/serializers/recordForm'
import { getAttrValues, setDMLValues } from '@/api/attsAPI'
import Characters from '@/views/cabinet/references/tables/Characters'
import { serializeAttr } from '@/api/serializers/attrs'

export default {
  components: {
    Characters,
    InputAutocomplete
  },
  data() {
    return {
      form: null,
      fetchFormStatus: IDLE,
      updateRecordStatus: IDLE,
      currentRowId: null,
      attributeValues: []
    }
  },
  computed: {
    ...apiStatusComputedFactory(['fetchFormStatus', 'updateRecordStatus']),
    getCurrentObjectName() {
      return this.$route.params.reference_name
    },
    getCurrentRowId() {
      return +this.$route.params.row_id
    },
    getReferenceControllers() {
      return (
        this.form &&
        this.form.filter((el) => !el.ref_table_type || el.ref_table_type == 1)
      )
    },
    getClassifierControllers() {
      return this.form && this.form.filter((el) => el.ref_table_type == 2)
    }
  },
  watch: {
    getClassifierControllers: {
      deep: true,
      handler(controllers) {
        for (let ctrl of controllers) {
          this.fetchAttrValues(ctrl)
        }
      }
    }
  },
  created() {
    this.fetchForm()
  },
  methods: {
    async onNewAttrAdded(attrIdx, { id, attr }) {
      this.attributeValues[attrIdx] = this.attributeValues[attrIdx].map(
        (attribute) => {
          if (attribute.id == id) {
            attribute.options.push(attr)
          }
          return attribute
        }
      )
    },
    async fetchAttrValues(classifier) {
      const { response, error } = await withAsync(getAttrValues, {
        dictionary_row_id: this.getCurrentRowId,
        dictionary_id: +this.$route.query.reference_id,
        attribute_dictionary_id: classifier.ref_table_id,
        attribute_dictionary_row_id: classifier[classifier.name]?.id
      })
      if (error) {
        this.fetchAttrStatus = ERROR
        return
      }
      this.fetchAttrStatus = SUCCESS
      if (response.data.result) {
        this.attributeValues.push(serializeAttr(response.data.result))
      } else {
        this.attributeValues = []
      }
    },
    async bindDMLAttributes(classifier, idx) {
      await withAsync(setDMLValues, {
        dictionary_id: +this.$route.query.reference_id,
        attribute_dict_id: classifier.ref_table_id,
        dictionary_row_id: this.currentRowId,
        attributes: this.attributeValues[idx].map((value) => ({
          attribute_id: value.id,
          key: value.selected.key
        }))
      })
    },
    getReferenceRecordsByName,
    async fetchForm() {
      this.fetchFormStatus = PENDING
      const { response, error } = await withAsync(getReferenceFormForEdit, {
        table_name: this.getCurrentObjectName,
        row_id: this.getCurrentRowId
      })
      if (error) {
        this.fetchFormStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.fetchFormStatus = SUCCESS
      this.form = serializeRecordForms(response.data.result.form)
    },
    async onSubmit() {
      this.saveRecordStatus = PENDING
      const { error, response } = await withAsync(
        createNewRecord,
        deserializeRecordForm({
          table_name: this.getCurrentObjectName,
          row_id: null,
          operation_type: 'add',
          form: this.form
        })
      )
      if (error) {
        this.saveRecordStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.saveRecordStatus = SUCCESS
      this.currentRowId = response.data.row_id
      if (this.getClassifierControllers?.length) {
        await this.getClassifierControllers.forEach((classifier, idx) => {
          this.bindDMLAttributes(classifier, idx)
        })
      }
      this.$router.go(-1)
      this.$toast.success('Запись успешно склонирована!')
    }
  }
}
</script>

<style lang="scss" scoped>
.reference {
  &__form {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 1rem;
  }

  &__controllers {
    display: flex;
    gap: 1rem;
  }
}
</style>
