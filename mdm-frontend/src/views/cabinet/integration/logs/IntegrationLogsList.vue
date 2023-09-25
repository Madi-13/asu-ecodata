<template>
  <div>
    <SearchController @searchClicked="onSearch($event)" placeholder="Найти" />
    <el-table
      highlight-current-row
      :data="tableData"
      border
      ref="table"
      class="mt--2"
      style="width: 100%"
      v-loading="fetchListStatusPending"
      empty-text="Данных нет"
    >
      <el-table-column prop="id" label="Код" width="60" />
      <el-table-column prop="process_id" label="ID процесса " min-width="20" />
      <el-table-column prop="status" label="Статус" min-width="30" />
      <el-table-column prop="description" label="Описание" min-width="150" />
      <el-table-column label="Дата" min-width="80">
        <template slot-scope="scope">
          {{ scope.row.send_date | fullDateTime }}
        </template>
      </el-table-column>
    </el-table>
    <PaginationContainer
      class="table__pagination"
      :pagination="params"
      @pageChange="paginate($event)"
    ></PaginationContainer>
  </div>
</template>

<script>
import PaginationContainer from '@/components/common/PaginationContainer.vue'
import SearchController from '@/components/common/SearchController.vue'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

import { getIntegrationLogs } from '@/api/integrationAPI'

export default {
  components: {
    PaginationContainer,
    SearchController
  },
  data() {
    return {
      operations: [
        {
          title: 'Создать запись',
          key: 'create',
          is_need_selected: false
        },
        {
          title: 'Изменить запись',
          key: 'edit',
          is_need_selected: true
        },
        {
          title: 'Удалить запись',
          key: 'delete',
          is_need_selected: true
        }
      ],
      params: {
        row_id: null,
        limit: 10,
        offset: 1,
        search: null,
        total: 0
      },
      tableData: null,
      fetchListStatus: IDLE
    }
  },
  computed: {
    ...apiStatusComputedFactory(['fetchListStatus'])
  },
  created() {
    this.fetchList()
  },
  methods: {
    onSearch(searchObj) {
      this.params.search = searchObj.search_value
      this.fetchList()
    },
    paginate(pageName) {
      this.params.offset = pageName
      this.fetchList()
    },
    async fetchList() {
      this.fetchListStatus = PENDING
      const { response, error } = await withAsync(
        getIntegrationLogs,
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
