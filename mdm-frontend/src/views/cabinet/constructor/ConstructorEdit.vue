<template>
  <form class="form" v-if="form">
    <div class="table__info">
      <BaseInput
        label="Наименование структуры"
        placeholder="Введите название"
        :width="100"
        name="table_name"
        type="textarea"
        :error="$v.form.table_name"
        @blur="$v.form.table_name.$touch()"
        v-model="form.table_name"
        size="large"
      ></BaseInput>
      <BaseDropdown
        label="Тип структуры"
        placeholder="Выберите из списка"
        v-model="form.table_type"
        :options="tableTypes"
        :error="$v.form.table_type"
        @blur="$v.form.table_type.$touch()"
        :width="100"
      />
      <BaseInput
        label="Префикс"
        placeholder="Введите префикс"
        :width="100"
        name="type_prefix"
        v-model="form.type_prefix"
        type="textarea"
        size="large"
        disabled
      ></BaseInput>
    </div>
    <el-table
      :data="form.columns"
      border
      style="width: 100%"
      @row-click="handleClickRow"
      empty-text="Данных нет"
      highlight-current-row
    >
      >
      <el-table-column prop="column_name" label="Наименование" min-width="160">
        <template slot-scope="scope">
          <BaseInput
            name="column_name"
            appereance="transparent"
            v-model="scope.row.column_name"
            :error="$v.form.columns.$each[scope.$index].column_name"
            @blur="$v.form.columns.$each[scope.$index].column_name.$touch()"
          />
        </template>
      </el-table-column>
      <el-table-column prop="column_desc" label="Описание" min-width="160">
        <template slot-scope="scope">
          <BaseInput
            name="column_desc"
            type="textarea"
            appereance="transparent"
            v-model="scope.row.column_desc"
          />
        </template>
      </el-table-column>
      <el-table-column prop="data_type" label="Тип" min-width="200">
        <template slot-scope="scope">
          <BaseColumnDropdown
            name="data_type"
            v-model="scope.row.data_type"
            :error="$v.form.columns.$each[scope.$index].data_type"
            @blur="$v.form.columns.$each[scope.$index].data_type.$touch()"
            @selected="onDateTypeChanged($event, scope.$index)"
            option-title-field="name_ru"
            option-value-field="id"
            v-if="objectTypes"
            :width="200"
            :options="objectTypes"
          />
        </template>
      </el-table-column>
      <el-table-column
        prop="foreign_table"
        label="Связанный классификатор/справочник"
        min-width="320"
      >
        <template slot-scope="scope">
          <ColumnAutocomplete
            type="transparent"
            name="foreign_table"
            v-model="scope.row.foreign_table"
            :disabled="
              scope.row.data_type && !scope.row.data_type.is_need_foreign_table
            "
            :error="$v.form.columns.$each[scope.$index].foreign_table"
            @blur="$v.form.columns.$each[scope.$index].foreign_table.$touch()"
            option-title-field="table_name"
            option-value-field="table_code"
          />
        </template>
      </el-table-column>
      <el-table-column
        prop="name"
        label="Системное наименование поля"
        min-width="260"
      >
        <template slot-scope="scope">
          <BaseInput
            name="column_system_name"
            disabled
            appereance="transparent"
            v-model="scope.row.column_system_name"
          />
        </template>
      </el-table-column>
      <el-table-column
        prop="max_symbol_length"
        label="Длина поля"
        min-width="150"
      >
        <template slot-scope="scope">
          <BaseInput
            name="max_symbol_length"
            appereance="transparent"
            :error="$v.form.columns.$each[scope.$index].max_symbol_length"
            @blur="
              $v.form.columns.$each[scope.$index].max_symbol_length.$touch()
            "
            type="number"
            v-model="scope.row.max_symbol_length"
          />
        </template>
      </el-table-column>
      <el-table-column
        prop="is_required"
        label="Обязательность заполнения"
        min-width="250"
      >
        <template slot-scope="scope">
          <BaseCheckbox
            v-model="scope.row.is_required"
            name="is_required"
            color="green"
          >
          </BaseCheckbox>
        </template>
      </el-table-column>
      <el-table-column prop="is_unique" label="Уникальность" min-width="130">
        <template slot-scope="scope">
          <BaseCheckbox
            v-model="scope.row.is_unique"
            name="is_unique"
            color="green"
          >
          </BaseCheckbox>
        </template>
      </el-table-column>
      <el-table-column prop="is_filterable" label="Фильтрация" min-width="130">
        <template slot-scope="scope">
          <BaseCheckbox
            v-model="scope.row.is_filterable"
            name="is_filterable"
            color="green"
          >
          </BaseCheckbox>
        </template>
      </el-table-column>
      <el-table-column
        prop="is_default_filterable"
        label="По умолчанию фильтрировать"
        min-width="250"
      >
        <template slot-scope="scope">
          <BaseCheckbox
            v-model="scope.row.is_default_filterable"
            name="is_default_filterable"
            color="green"
          >
          </BaseCheckbox>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="Операции" width="120">
        <template slot-scope="scope">
          <el-button
            type="success"
            icon="el-icon-refresh-left"
            v-if="scope.row.operation_type == 'drop'"
            circle
            @click="rollbackRow(scope.$index)"
          ></el-button>
          <el-button
            type="danger"
            icon="el-icon-delete"
            circle
            v-else
            @click="deleteRow(scope.$index)"
          ></el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="table__btns">
      <BaseButton
        color="green"
        width="15rem"
        class="table__button"
        @click.prevent="addNewRow"
        :disabled="$v.form.columns.$invalid && $v.form.columns.$dirty"
        >Добавить</BaseButton
      >
    </div>
    <div class="form__buttons mt--3">
      <BaseButton
        color="green"
        :disabled="$v.form.$invalid && $v.form.$dirty"
        @click.prevent="editRef"
        width="20rem"
        :loading="updateStatusPending"
        >Сохранить</BaseButton
      >
      <BaseButton color="outline-grey" width="20rem">Отмена</BaseButton>
    </div>
  </form>
</template>

<script>
import { required, requiredIf, minLength } from 'vuelidate/lib/validators'
import { tableTypes } from '@/fixtures/options'
import { mustStartWithCyrillicLetter } from '@/helpers/customValidators'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus
import {
  systemRefSerializer,
  systemRefDeserializeReq
} from '@/api/serializers/systemRefs'

import { getSystemRefById, editSystemRef } from '@/api/objectsAPI'

import { objectTypesComputed } from '@/services/columnTypesService'

import ColumnAutocomplete from '@/components/common/ColumnAutocomplete.vue'

export default {
  components: {
    ColumnAutocomplete
  },
  data() {
    return {
      getByIdStatus: IDLE,
      updateStatus: IDLE,
      form: null,
      selectedRow: null,
      deletedRows: [],
      tableTypes
    }
  },
  validations: {
    form: {
      table_name: { mustStartWithCyrillicLetter, required },
      table_type: { required },
      columns: {
        required,
        minLength: minLength(1),
        $each: {
          column_name: {
            mustStartWithCyrillicLetter,
            required
          },
          data_type: {
            required
          },
          foreign_table: {
            required: requiredIf(function (nestedModel) {
              return (
                nestedModel.data_type &&
                nestedModel.data_type.is_need_foreign_table
              )
            })
          },
          max_symbol_length: {
            required: requiredIf(function (nestedModel) {
              return (
                nestedModel.data_type && nestedModel.data_type.is_need_length
              )
            })
          },
          is_default_filterable: {
            required: requiredIf(function () {
              return (
                this.form.columns.filter((el) => !el.is_default_filterable)
                  .length == 0
              )
            })
          }
        }
      }
    }
  },
  computed: {
    ...apiStatusComputedFactory(['updateStatus']),
    ...objectTypesComputed,
    getCurrentId() {
      return this.$route.params.id
    }
  },
  created() {
    this.getById()
  },
  methods: {
    deleteRow(idx) {
      if (this.form.columns.length == 1) {
        this.$toast.warning('Структура не может быть пустым')
        return
      }
      this.form.columns[idx]['operation_type'] = 'drop'
      this.deletedRows.push(this.form.columns[idx])
      this.form.columns.splice(idx, 1)
    },
    rollbackRow(idx) {
      this.form.columns[idx]['operation_type'] = 'alter'
    },
    onDateTypeChanged(dataType, idx) {
      if (!dataType.is_need_foreign_table) {
        this.form.columns[idx].foreign_table = null
      }
    },
    addNewRow() {
      this.$v.form.columns.$touch()
      if (!this.$v.form.columns.$invalid) {
        this.form.columns.push(this.newRow())
      }
    },
    newRow() {
      return {
        id: null,
        column_system_name: null,
        column_name: null,
        column_desc: null,
        data_type: null,
        max_symbol_length: null,
        foreign_table: null,
        is_required: false,
        is_unique: false,
        is_filterable: false,
        is_default_filterable: false,
        operation_type: 'add'
      }
    },
    handleClickRow(row) {
      this.selectedRow = row
      this.form.columns.forEach((col) => {
        if (col.id == row.id && col['operation_type'] != 'drop')
          col['id']
            ? (col['operation_type'] = 'alter')
            : (col['operation_type'] = 'add')
      })
    },
    async getById() {
      this.getByIdStatus = PENDING
      const { response, error } = await withAsync(
        getSystemRefById,
        this.getCurrentId
      )
      if (error) {
        this.getByIdStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.getByIdStatus = SUCCESS
      this.form = systemRefSerializer(response.data.result[0]) || null
    },
    async editRef() {
      this.$v.form.$touch()
      let columns = [...this.form.columns, ...this.deletedRows]
      let payload = JSON.parse(JSON.stringify(this.form))
      payload.columns = columns
      if (!this.$v.form.$invalid) {
        this.updateStatus = PENDING
        const { error } = await withAsync(
          editSystemRef,
          systemRefDeserializeReq(payload, 'edit')
        )
        if (error) {
          this.updateStatus = ERROR
          this.$toast.error(error.message)
          return
        }
        this.updateStatus = SUCCESS
        this.$toast.success('Запись успено обновлена!')
        this.$router.push({ name: 'ConstructorList' })
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.table {
  &__info {
    display: grid;
    gap: 1.5rem;
    grid-template-columns: repeat(3, 1fr);
  }
  &__button {
    margin: 1rem auto;
  }
  &__btns {
    display: flex;
    align-items: center;
    gap: 1rem;
    width: 30%;
    margin: 0 auto;
  }
}
.form {
  &__buttons {
    display: flex;
    gap: 1rem;
  }
}
</style>
