<template>
  <div>
    <div class="mb--2 sticky">
      <SearchController @searchClicked="onSearch($event)" placeholder="Найти" />
      <TableController
        :selected-row="selectedRow"
        resource-name="structure"
        permission-name="structure_operations"
        :operations="operations"
        @create="$router.push({ name: 'ConstructorCreatePage' })"
        @edit="
          $router.push({
            name: 'ConstructorEditPage',
            params: { id: selectedRow.id }
          })
        "
        @delete="removeSystemRef"
        class="mb--2"
      />
    </div>
    <el-table
      :data="tableData"
      border
      ref="table"
      style="width: 100%"
      @row-dblclick="
        $router.push({
          name: 'ConstructorEditPage',
          params: { id: selectedRow.id }
        })
      "
      @row-click="handleClickRow"
      v-loading="fetchListStatusPending"
      empty-text="Данных нет"
      highlight-current-row
    >
      <el-table-column prop="table_code" label="Код" width="160" />
      <el-table-column prop="table_name" label="Наименование" min-width="80" />
      <el-table-column prop="table_type" label="Тип" min-width="160">
        <template slot-scope="scope">{{ scope.row.table_type.title }}</template>
      </el-table-column>
      <el-table-column prop="status" label="Статус" min-width="150" />
    </el-table>
    <PaginationContainer
      class="table__pagination"
      :pagination="params"
      @pageChange="paginate($event)"
    ></PaginationContainer>
  </div>
</template>

<script>
import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'

const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

import { getSystemRefs, deleteSystemRef } from '@/api/objectsAPI'

import PaginationContainer from '@/components/common/PaginationContainer.vue'
import TableController from '@/components/common/TableController.vue'
import SearchController from '@/components/common/SearchController.vue'

import { constructionOperations } from '@/fixtures/options'

export default {
  components: {
    PaginationContainer,
    TableController,
    SearchController
  },
  data() {
    return {
      tableData: null,
      selectedRow: null,
      params: {
        total: null,
        search: '',
        limit: 10,
        offset: 1
      },
      fetchListStatus: IDLE,
      deleteStatus: IDLE,
      operations: constructionOperations
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
    handleClickRow(row) {
      this.selectedRow = row
    },
    paginate(pageName) {
      this.params.offset = pageName
      this.fetchList()
    },
    gotoPage(pageName) {
      if (this.selectedRow) {
        this.$router.push({
          name: pageName,
          params: { id: this.selectedRow.id }
        })
      } else {
        this.$router.push({
          name: pageName
        })
      }
    },
    async fetchList() {
      this.fetchListStatus = PENDING
      const { response, error } = await withAsync(getSystemRefs, this.params)
      if (error) {
        this.fetchListStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.fetchListStatus = SUCCESS
      this.tableData = response.data.result
      this.params.total = response.data.total
    },
    async removeSystemRef() {
      const { error } = await withAsync(deleteSystemRef, {
        table_id: this.selectedRow.id
      })
      if (error) {
        this.deleteStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.deleteStatus = SUCCESS
      this.$toast.success('Запись успешна удалена!')
      this.fetchList()
    }
  }
}
</script>

<style lang="scss" scoped></style>
