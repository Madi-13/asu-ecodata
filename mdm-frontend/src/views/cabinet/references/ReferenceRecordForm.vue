<template>
  <div class="reference__form-wrapper">
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
              :disabled="!item.allow_edit || item.name == 'id'"
              placeholder="Введите значение"
              type="number"
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
              :width="100"
              mask="#######################"
              :disabled="!item.allow_edit"
              v-model="item[item.name]"
              :name="item.name"
              class=""
            ></BaseInput>
            <InputAutocomplete
              :name="item.name"
              v-if="item.type == 'autocomplete'"
              placeholder="Введите значение"
              :label="item.title"
              :searchMethod="getReferenceRecordsByName"
              :searchParams="{ table_name: item.ref_table_name }"
              v-model="item[item.name]"
            />
            <InputAutocomplete
              :name="item.name"
              v-if="item.type == 'tree'"
              placeholder="Введите значение"
              :disabled="!item.allow_edit"
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
              type="number"
              :disabled="!item.allow_edit && item.name == 'id'"
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
    </form>
    <div>
      <Characters
        v-for="(value, key) in attributeValues"
        :attribute="value"
        @addedNewAttr="onNewAttrAdded(key, $event)"
      />
    </div>
    <UploaderFile
      label="Документы и файлы"
      class="mt--2"
      @onRemove="removeFile($event)"
      isMultiple
      @change="addFileToUpload"
      :value="uploadFilesForm"
      acceptFileFormat="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.ms-excel, application/msword, 
        application/vnd.openxmlformats-officedocument.wordprocessingml.document, image/jpeg , image/png"
    />
    <div class="reference__controllers mt--3">
      <BaseButton
        color="green"
        @click.prevent="onSubmit"
        :loading="updateRecordStatusPending"
        v-can:update.disable="{
          resource: 'references',
          permission: getCurrentObjectName
        }"
        width="20rem"
        >Сохранить
      </BaseButton>
    </div>
    <template v-if="isClassifier">
      <EditAttr
        v-for="item in getClassifierControllers"
        :id="item.title"
        :classifier="item"
        @fetchAttrs="fetchAttrValues(item)"
      />
    </template>
  </div>
</template>

<script>
import InputAutocomplete from '@/components/common/InputAutocomplete.vue'
import UploaderFile from '@/components/common/UploaderFile.vue'
import Characters from '@/views/cabinet/references/tables/Characters'
import EditAttr from '@/views/cabinet/classifier/components/EditAttr.vue'

import {
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
import { serializeAttr } from '@/api/serializers/attrs'

import { getAttrValues, setDMLValues } from '@/api/attsAPI'
import { uploadFiles, getFiles, removeFile } from '@/api/objectsAPI'

export default {
  components: {
    EditAttr,
    Characters,
    InputAutocomplete,
    UploaderFile
  },
  data() {
    return {
      form: null,
      uploadFilesForm: [],
      fetchFormStatus: IDLE,
      updateRecordStatus: IDLE,
      fetchAttrStatus: IDLE,
      attributeValues: {}
    }
  },
  computed: {
    ...apiStatusComputedFactory([
      'fetchFormStatus',
      'updateRecordStatus',
      'fetchAttrStatus'
    ]),
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
    },
    isClassifier() {
      return this.$route.query && this.$route.query.is_classifier
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
    this.fetchFiles()
  },
  methods: {
    async removeFile(idx) {
      let selectedFileId = this.uploadFilesForm[idx].file_id
      if (selectedFileId) {
        const { error } = await withAsync(removeFile, {
          file_id: selectedFileId
        })
        if (error) {
          this.$toast.error(error.message)
          return
        }
      }
      this.uploadFilesForm.splice(idx, 1)
    },
    async fetchFiles() {
      const { response, error } = await withAsync(getFiles, {
        object_id: +this.$route.query.reference_id,
        row_id: this.getCurrentRowId
      })
      if (error) {
        this.$toast.error(error.message)
        return
      }
      this.uploadFilesForm = response.data.result || []
    },
    async addFileToUpload(file) {
      this.uploadFilesForm.push(file)
      await this.uploadRecordFiles(file)
    },
    async uploadRecordFiles(file) {
      let uploadFormPayload = {
        files: [file],
        object_id: +this.$route.query.reference_id,
        row_id: this.getCurrentRowId
      }
      const { error } = await withAsync(uploadFiles, uploadFormPayload)
      if (error) {
        return error
      }
    },
    async onNewAttrAdded(attrKey, { id, attr }) {
      this.attributeValues[attrKey] = this.attributeValues[attrKey].map(
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
        this.$set(
          this.attributeValues,
          classifier.ref_table_name,
          serializeAttr(response.data.result)
        )
      } else {
        this.attributeValues = {}
      }
    },
    async bindDMLAttributes(classifier) {
      await withAsync(setDMLValues, {
        dictionary_id: +this.$route.query.reference_id,
        attribute_dict_id: classifier.ref_table_id,
        dictionary_row_id: this.getCurrentRowId,
        attributes: this.attributeValues[classifier.ref_table_name].map(
          (value) => ({
            attribute_id: value.id,
            key: value.selected.key
          })
        )
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
    async updateClassifier() {
      this.updateRecordStatus = PENDING
      const { error } = await withAsync(
        updateRecord,
        deserializeRecordForm({
          table_name: this.getCurrentObjectName,
          row_id: this.getCurrentRowId,
          operation_type: 'alter',
          form: this.form
        })
      )
      if (error) {
        this.updateRecordStatus = ERROR
        this.$toast.error(error.message)
        return error
      }
      this.updateRecordStatus = SUCCESS
    },
    async onSubmit() {
      let err = await this.updateClassifier()
      if (err) return

      if (this.getClassifierControllers?.length) {
        await this.getClassifierControllers.forEach((classifier, idx) => {
          this.bindDMLAttributes(classifier, idx)
        })
      }
      // this.$router.go(-1)
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
    &-wrapper {
      padding: 20px;
    }
  }

  &__controllers {
    display: flex;
    gap: 1rem;
  }
}
</style>
