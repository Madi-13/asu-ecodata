<template>
  <div>
    <SearchController @searchClicked="onSearch($event)" placeholder="Найти" />
    <TableController
      :selected-row="selectedRow"
      :operations="operations"
      @create="openModal('createIntegrationRouteModal')"
      @edit="openModal('editIntegrationRouteModal')"
      resource-name="integration"
      permission-name="integration_operations"
      @delete="onDelete"
      class="mb--2"
    />
    <el-table
      highlight-current-row
      :data="tableData"
      @row-dblclick="openModal('editIntegrationRouteModal')"
      border
      ref="table"
      style="width: 100%"
      @row-click="handleClickRow"
      v-loading="fetchListStatusPending"
      empty-text="Данных нет"
    >
      <el-table-column prop="id" label="Код" width="60" />
      <el-table-column
        prop="property_name"
        label="Наименование"
        min-width="80"
      />
      <el-table-column prop="property_type" label="Тип " min-width="30" />
      <el-table-column prop="path" label="Путь" min-width="150" />
    </el-table>
    <PaginationContainer
      class="table__pagination"
      :pagination="params"
      @pageChange="paginate($event)"
    ></PaginationContainer>
    <IntegrationRouteCreateModal @onClose="fetchList" />
    <IntegrationRouteEditModal :selected="selectedRow" @onClose="fetchList" />
  </div>
</template>

<script>
import SearchController from '@/components/common/SearchController.vue'
import TableController from '@/components/common/TableController.vue'
import PaginationContainer from '@/components/common/PaginationContainer.vue'

import IntegrationRouteCreateModal from './components/modals/IntegrationRouteCreateModal.vue'
import IntegrationRouteEditModal from './components/modals/IntegrationRouteEditModal.vue'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

import {
  getIntegrationRoutes,
  operateIntegrationRoute
} from '@/api/integrationAPI'

import { operationsIntegrationTables } from '@/fixtures/options'

export default {
  components: {
    TableController,
    SearchController,
    PaginationContainer,
    IntegrationRouteEditModal,
    IntegrationRouteCreateModal
  },
  data() {
    return {
      operations: operationsIntegrationTables,
      params: {
        row_id: null,
        limit: 10,
        offset: 1,
        total: 0,
        search: null
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
      const { error } = await withAsync(operateIntegrationRoute, {
        operation_type: 'drop',
        property: {
          id: this.selectedRow.id,
          path: null,
          property_name: null,
          property_type: null,
          rq_group: null,
          rq_token: null
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
      const { response, error } = await withAsync(
        getIntegrationRoutes,
        this.params
      )
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
