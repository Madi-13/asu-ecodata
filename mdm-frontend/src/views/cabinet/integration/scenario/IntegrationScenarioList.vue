<template>
  <div>
    <SearchController @searchClicked="onSearch($event)" placeholder="Найти" />
    <TableController
      :selected-row="selectedRow"
      :operations="operations"
      @create="openModal('createScenarioModal')"
      @edit="openModal('editScenarioModal')"
      resource-name="integration"
      permission-name="integration_operations"
      @delete="onDelete"
      class="mb--2"
    />
    <el-table
      highlight-current-row
      :data="tableData"
      border
      ref="table"
      style="width: 100%"
      @row-dblclick="openModal('editScenarioModal')"
      @row-click="handleClickRow"
      v-loading="fetchListStatusPending"
      empty-text="Данных нет"
    >
      <el-table-column prop="id" label="Код" width="160" />
      <el-table-column
        prop="scenario_name"
        label="Наименование"
        min-width="80"
      />
      <el-table-column prop="status" label="Статус" min-width="150" />
    </el-table>
    <PaginationContainer
      class="table__pagination"
      :pagination="params"
      @pageChange="paginate($event)"
    ></PaginationContainer>
    <CreateScenarioModal @onClose="fetchList" />
    <EditScenarioModal :selected="selectedRow" @onClose="fetchList" />
  </div>
</template>

<script>
import SearchController from '@/components/common/SearchController.vue'
import TableController from '@/components/common/TableController.vue'
import PaginationContainer from '@/components/common/PaginationContainer.vue'
import CreateScenarioModal from './components/CreateScenarioModal.vue'
import EditScenarioModal from './components/EditScenarioModal.vue'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

import { getScenarios, operateScenario } from '@/api/integrationAPI'

import { operationsIntegrationTables } from '@/fixtures/options'

export default {
  components: {
    SearchController,
    EditScenarioModal,
    CreateScenarioModal,
    PaginationContainer,
    TableController
  },
  data() {
    return {
      operations: operationsIntegrationTables,
      params: {
        row_id: null,
        limit: 10,
        offset: 1,
        search: null,
        total: 0
      },
      tableData: null,
      selectedRow: null,
      fetchListStatus: IDLE,
      deleteStatus: IDLE
    }
  },
  computed: {
    ...apiStatusComputedFactory(['fetchListStatus', 'deleteStatus'])
  },
  created() {
    this.fetchList()
  },
  methods: {
    onSearch(searchObj) {
      this.params.search = searchObj.search_value
      this.fetchList()
    },
    openModal(modalName) {
      this.$modal.show(modalName)
    },
    handleClickRow(row) {
      this.selectedRow = row
    },
    paginate(pageName) {
      this.params.offset = pageName
      this.fetchList()
    },
    async onDelete() {
      this.deleteStatus = PENDING
      const { error } = await withAsync(operateScenario, {
        operation_type: 'drop',
        scenario: {
          id: this.selectedRow.id,
          scenario_name: null,
          status: 'ACTIVE'
        }
      })
      if (error) {
        this.deleteStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.deleteStatus = SUCCESS
      this.fetchList()
    },
    async fetchList() {
      this.fetchListStatus = PENDING
      const { response, error } = await withAsync(getScenarios, this.params)
      if (error) {
        this.fetchListStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.fetchListStatus = SUCCESS
      this.tableData = response.data.result
      this.params.total = response.data.total
    }
  }
}
</script>

<style lang="scss" scoped></style>
