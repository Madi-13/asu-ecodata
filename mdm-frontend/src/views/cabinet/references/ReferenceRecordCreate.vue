<template>
  <div>
    <form v-if="form">
      <div class="reference__form">
        <template
          v-for="(item, idx) in getReferenceControllers"
          :id="item.title"
        >
          <div v-if="item.name != 'id'">
            <BaseInput
              v-if="item.type == 'input-text'"
              :label="item.title"
              placeholder="Введите значение"
              :disabled="!item.allow_edit"
              :width="100"
              v-model="item[item.name]"
              :name="item.name"
            ></BaseInput>
            <BaseInput
              v-if="item.type == 'input-number'"
              :label="item.title"
              placeholder="Введите значение"
              :width="100"
              :disabled="item.name == 'id' || !item.allow_edit"
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
          </div>
        </template>
      </div>
      <p
        class="title--tertiary text--dark"
        v-if="getClassifierControllers && getClassifierControllers.length > 0"
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
              :disabled="!item.allow_edit"
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
              :disabled="!item.allow_edit"
              placeholder="Введите значение"
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
        @addedNewAttr="onNewAttrAdded(idx, $event)"
        v-for="(attr, idx) of attributeValues"
        :attribute="attr"
      />
    </div>
    <UploaderFile
      @onRemove="removeFile($event)"
      label="Документы и файлы"
      class="mt--2"
      isMultiple
      v-model="uploadFilesForm"
      acceptFileFormat="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.ms-excel, application/msword, 
        application/vnd.openxmlformats-officedocument.wordprocessingml.document, image/jpeg , image/png"
    />
    <div class="reference__controllers mt--3">
      <BaseButton
        color="green"
        @click.prevent="onSubmit"
        :loading="saveRecordStatusPending"
        width="20rem"
        >Сохранить
      </BaseButton>
      <BaseButton color="outline-grey" width="20rem"> Отмена </BaseButton>
    </div>
  </div>
</template>

<script>
import InputAutocomplete from '@/components/common/InputAutocomplete.vue'
import UploaderFile from '@/components/common/UploaderFile.vue'

import {
  getReferenceRecordsByName,
  createNewRecord
} from '@/api/referenceRecordsAPI'
import { uploadFiles } from '@/api/objectsAPI'

import { deserializeRecordForm } from '@/api/serializers/recordForm'

import { getReferenceFormForCreate } from '@/api/referenceRecordsAPI'
import { serializeRecordForms } from '@/api/serializers/recordForm'
import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
import Characters from '@/views/cabinet/references/tables/Characters'
import BaseColumnDropdown from '@/components/base/BaseColumnDropdown.vue'
import { getAttrValues, setDMLValues } from '@/api/attsAPI'
import { serializeAttr } from '@/api/serializers/attrs'

const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

export default {
  components: {
    BaseColumnDropdown,
    Characters,
    InputAutocomplete,
    UploaderFile
  },
  data() {
    return {
      form: null,
      uploadFilesForm: [],
      fetchFormStatus: IDLE,
      saveRecordStatus: IDLE,
      currentRowId: null,
      attributeValues: []
    }
  },
  computed: {
    ...apiStatusComputedFactory(['fetchFormStatus', 'saveRecordStatus']),
    getCurrentObjectName() {
      return this.$route.params.reference_name
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
      handler(controllers) {
        for (let ctrl of controllers) {
          this.fetchAttrValues(ctrl)
        }
      },
      deep: true
    }
  },
  async created() {
    await this.fetchForm()
  },
  methods: {
    removeFile(idx) {
      this.uploadFilesForm.splice(idx, 1)
    },
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
    async uploadRecordFiles() {
      let uploadFormPayload = {
        files: this.uploadFilesForm,
        object_id: +this.$route.query.reference_id,
        row_id: this.currentRowId
      }
      const { error } = await withAsync(uploadFiles, uploadFormPayload)
      if (error) {
        return error
      }
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
    getReferenceRecordsByName,
    async fetchForm() {
      this.fetchFormStatus = PENDING
      const { response, error } = await withAsync(getReferenceFormForCreate, {
        table_name: this.getCurrentObjectName,
        row_id: null
      })
      if (error) {
        this.fetchFormStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.fetchFormStatus = SUCCESS
      this.form = serializeRecordForms(response.data.result.form)
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
      const err = await this.uploadRecordFiles()
      if (err) {
        this.$toast.error('Не удалось сохранить файлы')
        return
      }
      this.$router.go(-1)
      this.$toast.success('Запись успешно обновлена!')
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
