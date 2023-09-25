<template>
  <div>
    <form class="import__form">
      <BaseDropdown
        label="Тип загрузки"
        placeholder="Выберите из списка"
        optionValueField="code"
        :error="$v.form.table_type"
        @blur="$v.form.table_type.$touch()"
        v-model="form.table_type"
        :options="tableTypesForImport"
        :width="100"
      />
      <div class=""></div>
      <BaseDropdown
        label="Тип объекта"
        placeholder="Введите название"
        optionValueField="table_id"
        optionTitleField="table_show"
        is-searchable
        :disabled="
          !form.table_type ||
          fetchObjectsStatusPending ||
          form.table_type.id == null
        "
        :error="$v.form.object_type"
        @blur="$v.form.object_type.$touch()"
        v-model="form.object_type"
        :options="objects"
        :width="100"
      />
      <UploaderFile
        v-model="uploadFileForm"
        @onRemove="removeFile"
        :disabled="
          !form.table_type || (!form.table_type?.id ? false : !form.object_type)
        "
      />
      <BaseButton
        color="green"
        @click.prevent="downloadTemplateFile"
        :is-loading="downloadTemplateStatusPending"
        height="50px"
        :disabled="
          !form.table_type || (!form.table_type?.id ? false : !form.object_type)
        "
      >
        Выгрузить шаблон
      </BaseButton>
      <BaseButton
        color="green"
        @click.prevent="uploadFile"
        height="50px"
        :is-loading="uploadFileStatusPending"
        :disabled="
          !form.table_type ||
          (!form.table_type?.id ? false : !form.object_type) ||
          !uploadFileForm
        "
      >
        Загрузить шаблон
      </BaseButton>
      <!-- <BaseButton
        color="green"
        @click.prevent=""
        :is-loading="downloadTemplateStatusPending"
        height="50px"
        :disabled="
          !form.table_type || (!form.table_type?.id ? false : !form.object_type)
        "
      >
        Выгрузить шаблон
      </BaseButton> -->
    </form>
    <el-table
      class="mt--2"
      highlight-current-row
      :data="monitoringData"
      border
      ref="table"
      style="width: 100%"
      empty-text="Данных нет"
    >
      <el-table-column prop="id" label="Код" width="340" />
      <el-table-column label="Название сущности" prop="file_name" width="300">
      </el-table-column>
      <el-table-column label="Дата начало" width="150">
        <template slot-scope="scope">
          {{ scope.row.date_start | fullDateTime }}
        </template>
      </el-table-column>
      <el-table-column label="Дата окончания" width="150">
        <template slot-scope="scope">
          {{ scope.row.date_end | fullDateTime }}
        </template>
      </el-table-column>
      <el-table-column prop="states" label="Процент загрузки" min-width="100" />
      <el-table-column prop="status" label="Состояние" min-width="150" />
      <el-table-column prop="messages" label="Сообщение" min-width="300" />
      <el-table-column fixed="right" label="Операции" width="120">
        <template slot-scope="scope">
          <el-button
            type="danger"
            icon="el-icon-close"
            circle
            @click="terminateDownload(scope.row.id)"
          ></el-button>
        </template>
      </el-table-column>
    </el-table>
    <PaginationContainer
      class="table__pagination"
      :pagination="params"
      @pageChange="paginate($event)"
    ></PaginationContainer>
  </div>
  </div>
</template>

<script>
import InputAutocomplete from '@/components/common/InputAutocomplete.vue'
import UploaderFile from '@/components/common/UploaderFile.vue'
import PaginationContainer from '@/components/common/PaginationContainer.vue'

import { tableTypesForImport } from '@/fixtures/options'
import { required } from 'vuelidate/lib/validators'
import { getMenuContent } from '@/api/menuAPI'

import { uploadFile, downloadFileTemplate, getMonitoring, cancelMonitoring } from '@/api/importAPI'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

export default {
  components: {
    InputAutocomplete,
    UploaderFile,
    PaginationContainer
  },
  data() {
    return {
      monitoringData: null,
      params: {
        limit: 10,
        offset: 1,
        total: 0
      },

      tableTypesForImport,
      objects: null,

      fetchObjectsStatus: IDLE,
      downloadTemplateStatus: IDLE,
      uploadFileStatus: IDLE,

      form: {
        table_type: null,
        object_type: null
      },

      uploadFileForm: null
    }
  },
  validations: {
    form: {
      table_type: {
        required
      },
      object_type: {
        required
      }
    }
  },
  computed: {
    ...apiStatusComputedFactory([
      'fetchObjectsStatus',
      'downloadTemplateStatus',
      'uploadFileStatus'
    ])
  },
  watch: {
    'form.table_type': {
      handler(newValue) {
        if (newValue) {
          this.form.object_type = null
          this.fetchMenu()
        }
      }
    }
  },
  created() {
    this.fetchMonitoringData()
  },
  methods: {
    removeFile(){
      this.uploadFileForm = null
    },
    paginate(pageName) {
      this.params.offset = pageName
      this.fetchMonitoringData()
    },
    async fetchMonitoringData() {
      const { response, error } = await withAsync(getMonitoring, this.params)
      if (error) {
        return
      }
      if (response && response.data) {
        this.params.total = response.data.total
        this.monitoringData = response.data.import
      }
    },
    async terminateDownload(rowId){
      const {response, error} = await withAsync(cancelMonitoring, {
        row_id: rowId
      })
      if(error){
        this.$toast.error('Не удалось отменить запрос')
        return
      }
      this.$toast.success('Загрузка успешно отменена!')
      this.fetchMonitoringData()
    },
    async fetchMenu() {
      this.fetchObjectsStatus = PENDING
      const { response, error } = await withAsync(getMenuContent, {
        table_type_id: this.form.table_type.id
      })
      if (error) {
        this.fetchObjectsStatus = ERROR
        return
      }
      this.fetchObjectsStatus = SUCCESS
      this.objects = response.data
    },
    async downloadTemplateFile() {
      this.downloadTemplateStatus = PENDING
      const { response, error } = await withAsync(downloadFileTemplate, {
        table_name: this.form.table_type.id
          ? this.form.object_type.table_name
          : this.form.table_type.table_name,
        import_type_code: this.form.table_type.code
      })
      if (error) {
        this.downloadTemplateStatus = ERROR
        this.$toast.error('Не удалось выгрузить шаблон')
        return
      }
      this.downloadTemplateStatus = SUCCESS
      if (!response.data.file_body) {
        this.$toast.warning('Отсуствует шаблон файла')
        return
      }
      const source = `data:application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;base64,${response.data.file_body}`
      const link = document.createElement('a')
      link.href = source
      link.download = response.data.file_name
      link.click()
      link.remove()
    },
    async uploadFile() {
      this.uploadFileStatus = PENDING
      const { error } = await withAsync(uploadFile, {
        table_name: this.form.object_type && this.form.object_type.table_name,
        file_body: this.uploadFileForm.file_body,
        import_type_code: this.form.table_type.code
      })
      if (error) {
        this.uploadFileStatus = ERROR
        this.$toast.error('Не удалось загрузить файл')
        return
      }
      this.uploadFileStatus = SUCCESS
      this.fetchMonitoringData()
      this.$toast.success('Файл успешно загружен!')
    }
  }
}
</script>

<style lang="scss" scoped>
.import {
  &__form {
    display: grid;
    grid-template-columns: repeat(2, 50%);
    gap: 1rem;
  }
}
</style>
