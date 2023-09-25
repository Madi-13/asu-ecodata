<template>
  <div>
    <div class="sticky">
      <TableController
        :selected-row="selectedRole"
        :operations="operations"
        @create="$modal.show('createRoleModal')"
        @edit="$modal.show('editRoleModal')"
        @block="blockRole"
        class="mb--2"
      />
    </div>
    <el-table
      highlight-current-row
      :data="tableData"
      border
      @row-dblclick="$modal.show('editRoleModal')"
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
    <CreateRoleModal />
    <EditRoleModal :selected-role="selectedRole" />
  </div>
</template>

<script>
import TableController from '@/components/common/TableController.vue'
import PaginationContainer from '@/components/common/PaginationContainer.vue'
import CreateRoleModal from '../roles/modals/CreateRoleModal.vue'
import EditRoleModal from '../roles/modals/EditRoleModal.vue'

import { operationsRoleTable } from '@/fixtures/options'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
import { getRoles } from '@/api/roleAPI'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

export default {
  components: {
    TableController,
    PaginationContainer,
    CreateRoleModal,
    EditRoleModal
  },
  data() {
    return {
      operations: operationsRoleTable,
      blockStatus: IDLE,
      fetchStatus: IDLE,
      roles: null,
      selectedRole: null,
      params: {
        limit: 10,
        offset: 1,
        total: 0
      }
    }
  },
  //   watch: {
  //     params: {
  //       handler() {
  //         this.fetchRoles()
  //       },
  //       immediate: true
  //     }
  //   },
  computed: {
    ...apiStatusComputedFactory(['fetchStatus'])
  },
  methods: {
    handleClickRow(row) {
      this.selectedRole = row
    },
    paginate(pageName) {
      this.params.offset = pageName
      this.fetchRoles()
    },
    async blockRole() {
      this.blockStatus = PENDING
      const { error } = await withAsync(blockRole)
      if (error) {
        this.blockStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.blockStatus = SUCCESS
      this.$toast.success('Роль успешно заблокирована!')
    },
    async fetchRoles() {
      this.fetchRolesStatus = PENDING
      const { response, error } = await withAsync(getRoles, {
        params: this.params
      })
      if (error) {
        this.fetchRolesStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.roles = response.data
      this.params.total = response.data.total
      this.fetchRolesStatus = SUCCESS
    }
  }
}
</script>

<style lang="scss" scoped></style>
