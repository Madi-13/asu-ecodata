<template>
  <div v-if="tableData">
    <div class="sticky">
      <p class="title--secondary mb--2">
        {{
          `${isClassifier ? 'Классификатор' : 'Справочник'} "${tableShowName}"`
        }}
      </p>
      <SearchController
        :filterOptions="tableData.header"
        placeholder="Найти"
        @searchClicked="onSearch($event)"
      />
      <TableController
        :selected-row="selectedRow"
        :permissionName="params.table_name"
        :operations="getOperations"
        @create="gotoPage('ReferenceRecordCreate')"
        @edit="gotoPage('ReferenceRecordEdit')"
        @clone="gotoPage('ReferenceRecordClone')"
        @set-state="$modal.show('blockReasonModal')"
        @create-attr="$modal.show('createAttrModal')"
        @edit-attr="$modal.show('editAttrModal')"
        @delete="deleteRef"
        @export="onExportReference"
      />
    </div>
    <el-table
      highlight-current-row
      :data="tableData.body.result"
      class="mt--2"
      @row-dblclick="gotoPage('ReferenceRecordEdit')"
      border
      ref="table"
      empty-text="Данных нет"
      style="width: 100%"
      @row-click="handleClickRow"
      v-loading="fetchTableDataStatusPending"
    >
      <template v-for="column in tableData.header">
        <el-table-column
          :prop="column.col_name"
          :label="column.title"
          v-if="
            column.col_name != 'row_status_code' &&
            column.frontend_type != 'datepicker'
          "
        >
          <template slot-scope="scope">
            {{ scope.row[column.col_name] }}
          </template>
        </el-table-column>
        <el-table-column
          :prop="column.col_name"
          :label="column.title"
          v-if="
            column.col_name != 'row_status_code' &&
            column.frontend_type == 'datepicker'
          "
        >
          <template slot-scope="scope">
            {{ scope.row[column.col_name] | normalizeDate }}
          </template>
        </el-table-column>
      </template>
    </el-table>
    <PaginationContainer
      class="table__pagination"
      :pagination="params"
      @pageChange="paginate($event)"
    ></PaginationContainer>
    <CreateAttrModal :classifier-id="selectedRow && selectedRow.id" />
    <EditAttrModal :classifier-id="selectedRow && selectedRow.id" />
    <BlockReasonModal :selected="selectedRow" @onClose="onCloseModal" />
  </div>
</template>

<script>
import { deleteRecord, getReferenceRecords } from '@/api/referenceRecordsAPI'
import { downloadFile } from '@/api/importAPI'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'

const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

import TableController from '@/components/common/TableController.vue'
import SearchController from '@/components/common/SearchController.vue'
import PaginationContainer from '@/components/common/PaginationContainer.vue'
import CreateAttrModal from '@/views/cabinet/classifier/components/CreateAttrModal.vue'
import EditAttrModal from '@/views/cabinet/classifier/components/EditAttrModal.vue'
import BlockReasonModal from './modals/BlockReasonModal.vue'

import {
  operationsRefTable,
  operationsClassifierTable
} from '@/fixtures/options'

export default {
  components: {
    TableController,
    SearchController,
    PaginationContainer,
    CreateAttrModal,
    EditAttrModal,
    BlockReasonModal
  },
  data() {
    return {
      selectedRow: null,
      fetchTableDataStatus: IDLE,
      tableData: null,
      isClassifier: false,
      tableShowName: null,
      operationsRefTable,
      operationsClassifierTable,
      params: {
        table_name: this.getCurrentRefName,
        limit: 15,
        offset: 1,
        total: null,
        search: []
      }
    }
  },
  computed: {
    ...apiStatusComputedFactory(['fetchTableDataStatus']),
    getCurrentRefName() {
      return this.$route.params.reference_name
    },
    getCurrentRefId() {
      return this.$route.query.table_id
    },
    getOperations() {
      return this.isClassifier ? operationsClassifierTable : operationsRefTable
    }
  },
  watch: {
    $route: {
      handler(newValue) {
        this.isClassifier =
          newValue.query && String(newValue.query.is_classifier) != 'false'
        this.tableShowName = newValue.query && newValue.query.table_show
        this.params.table_name =
          newValue.params && newValue.params.reference_name
        this.fetchTableData()
      },
      deep: true,
      immediate: true
    }
  },
  methods: {
    onCloseModal() {
      this.$modal.hide('blockReasonModal')
      this.fetchTableData()
    },
    onSearch(searchObj) {
      this.params.search = [searchObj]
      this.fetchTableData()
    },
    gotoPage(pageName) {
      this.$router.push({
        name: pageName,
        params: {
          reference_name: this.params.table_name,
          row_id: this.selectedRow && this.selectedRow.id
        },
        query: {
          reference_id: this.getCurrentRefId
        }
      })
    },
    handleClickRow(row) {
      this.selectedRow = row
      this.operationsRefTable = this.operationsRefTable.map((operation) => {
        if (operation.key == 'set-state') {
          operation.title =
            this.selectedRow.row_status_code == 'ACTIVE'
              ? 'Блокировать запись'
              : 'Разблокировать запись'
        }
        return operation
      })
    },
    paginate(pageNum) {
      this.params.offset = pageNum
      this.fetchTableData()
    },
    deleteRef() {
      withAsync(deleteRecord, {
        table_name: this.getCurrentRefName,
        operation_type: 'drop',
        row_id: this.selectedRow.id
      }).then(() => {
        this.$toast.success('Вы успешно удалили ' + this.selectedRow.imya)
        this.fetchTableData()
      })
    },
    async onExportReference() {
      const { response, error } = await withAsync(downloadFile, {
        import_type_code: 'DICTIONARY',
        table_name: this.getCurrentRefName
      })
      if (error) {
        this.$toast.error('Не удалось выгрузить данные')
        return
      }
      if (!response.data.file_body) {
        this.$toast.warning('Отсуствует файл')
        return
      }
      const source = `data:application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;base64,${response.data.file_body}`
      const link = document.createElement('a')
      link.href = source
      link.download = response.data.file_name
      link.click()
      link.remove()
    },
    async fetchTableData() {
      this.fetchTableDataStatus = PENDING
      const { response, error } = await withAsync(
        getReferenceRecords,
        this.params
      )
      if (error) {
        this.$toast.error(error.message)
        this.fetchTableDataStatus = ERROR
        return
      }
      this.tableData = response.data.result
      this.fetchTableDataStatus = SUCCESS
      this.params.total = response.data.result.body.total_count
    }
  }
}
</script>

<style lang="scss" scoped></style>
