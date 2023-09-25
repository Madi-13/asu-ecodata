<template>
  <form class="form">
    <div class="table__info">
      <BaseInput
        label="Наименование структуры"
        placeholder="Введите название"
        :width="100"
        type="textarea"
        name="table_name"
        :error="$v.form.table_name"
        @blur="$v.form.table_name.$touch()"
        v-model="form.table_name"
        size="large"
      ></BaseInput>
      <BaseDropdown
        label="Тип структуры"
        placeholder="Выберите из списка"
        :error="$v.form.table_type"
        @blur="$v.form.table_type.$touch()"
        v-model="form.table_type"
        :options="tableTypes"
        :width="100"
      />
      <BaseInput
        label="Префикс"
        placeholder="Введите префикс"
        :width="100"
        disabled
        type="textarea"
        name="table_prefix"
        :error="$v.form.table_prefix"
        @blur="$v.form.table_prefix.$touch()"
        v-model="form.table_prefix"
        size="large"
      ></BaseInput>
    </div>
    <el-table
      :data="form.columns"
      border
      empty-text="Данных нет"
      highlight-current-row
      style="width: 100%"
      @row-click="handleClickRow"
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
            name="column_name"
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
            :error="$v.form.columns.$each[scope.$index].data_type"
            @blur="$v.form.columns.$each[scope.$index].data_type.$touch()"
            @selected="onDateTypeChanged($event, scope.$index)"
            v-model="scope.row.data_type"
            option-title-field="name_ru"
            option-value-field="id"
            v-if="objectTypes"
            :width="80"
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
            :error="$v.form.columns.$each[scope.$index].foreign_table"
            @blur="$v.form.columns.$each[scope.$index].foreign_table.$touch()"
            :disabled="
              scope.row.data_type && !scope.row.data_type.is_need_foreign_table
            "
            v-model="scope.row.foreign_table"
            option-title-field="table_name"
            option-value-field="table_code"
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
            :disabled="
              !scope.row.data_type || !scope.row.data_type.is_need_length
            "
            v-model="scope.row.max_symbol_length"
            mask="############"
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
            type="danger"
            icon="el-icon-delete"
            circle
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
        :disabled="$v.form.columns.$invalid && $v.form.columns.$dirty"
        @click.prevent="addNewRow"
        >Добавить</BaseButton
      >
    </div>

    <div class="form__buttons mt--3">
      <BaseButton
        color="green"
        :disabled="$v.form.$invalid && $v.form.$dirty"
        @click.prevent="createNewRef"
        width="20rem"
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
import { objectTypesComputed } from '@/services/columnTypesService'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

import { createSystemRef } from '@/api/objectsAPI'
import { systemRefDeserializeReq } from '@/api/serializers/systemRefs'

import ColumnAutocomplete from '@/components/common/ColumnAutocomplete.vue'

import { updateChildrens } from '@/services/navs.js'

export default {
  components: {
    ColumnAutocomplete
  },
  data() {
    return {
      selectedRow: null,
      createStatus: IDLE,
      form: {
        table_name: null,
        table_type: null,
        table_prefix: null,
        columns: [this.newRow()]
      },
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
    ...objectTypesComputed,
    ...apiStatusComputedFactory(['createStatus'])
  },
  methods: {
    deleteRow(idx) {
      if (this.form.columns.length == 1) {
        this.$toast.warning('Структура не может быть пустым')
        return
      }
      this.form.columns.splice(idx, 1)
    },
    onDateTypeChanged(dataType, idx) {
      if (!dataType.is_need_foreign_table) {
        this.form.columns[idx].foreign_table = null
      }
    },
    handleClickRow(row) {
      this.selectedRow = row
    },
    addNewRow() {
      this.$v.form.columns.$touch()
      if (!this.$v.form.columns.$invalid) {
        this.form.columns.push(this.newRow())
      }
    },
    newRow() {
      return {
        column_name: null,
        column_desc: null,
        data_type: null,
        foreign_table: null,
        max_symbol_length: null,
        is_required: false,
        is_unique: false,
        is_filterable: false,
        is_default_filterable: false
      }
    },
    async createNewRef() {
      this.$v.form.$touch()
      if (!this.$v.form.$invalid) {
        this.createStatus = PENDING
        const { error } = await withAsync(
          createSystemRef,
          systemRefDeserializeReq(this.form)
        )
        if (error) {
          this.createStatus = ERROR
          this.$toast.error(error.message)
          return
        }
        await updateChildrens(this.form.table_type.id)
        this.$toast.success('Новая запись успешно сохранена')
        this.$router.push({ name: 'ConstructorList' })
        this.createStatus = SUCCESS
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
