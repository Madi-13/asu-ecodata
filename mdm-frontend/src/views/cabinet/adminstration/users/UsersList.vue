<template>
  <div>
    <div class="sticky">
      <TableController
        :selected-row="selectedUser"
        :operations="operations"
        @create="$modal.show('createUserModal')"
        @edit="$modal.show('editUserModal')"
        @block="removeUser"
        resource-name="adminstration"
        permission-name="user_operations"
        class="mb--2"
      />
      <el-table
        highlight-current-row
        :data="users"
        @row-dblclick="$modal.show('editUserModal')"
        border
        ref="table"
        style="width: 100%"
        @row-click="handleClickRow"
        v-loading="fetchStatusPending"
        empty-text="Данных нет"
      >
        <el-table-column prop="id" label="Код" width="340" />
        <el-table-column label="ФИО" width="300">
          <template slot-scope="scope">
            {{ scope.row.firstName }} {{ scope.row.lastName }}
          </template>
        </el-table-column>
        <el-table-column prop="username" label="Логин" width="300" />
        <el-table-column
          prop="email"
          label="Электронная почта"
          min-width="300"
        />
      </el-table>
      <PaginationContainer
        class="table__pagination"
        :pagination="params"
        @pageChange="paginate($event)"
      ></PaginationContainer>
    </div>
    <CreateUserModal @onClose="fetchUsers" />
    <EditUserModal :selected-user="selectedUser" />
  </div>
</template>

<script>
import TableController from '@/components/common/TableController.vue'
import PaginationContainer from '@/components/common/PaginationContainer.vue'
import CreateUserModal from './modals/CreateUserModal.vue'
import EditUserModal from './modals/EditUserModal.vue'

import { operationsUserTable } from '@/fixtures/options'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
import { getUsers, deleteUser } from '@/api/userAPI'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

export default {
  components: {
    TableController,
    PaginationContainer,
    CreateUserModal,
    EditUserModal
  },
  data() {
    return {
      operations: operationsUserTable,
      blockStatus: IDLE,
      fetchStatus: IDLE,
      selectedUser: null,
      users: null,
      params: {
        limit: 10,
        offset: 1,
        total: 0
      }
    }
  },
  created() {
    this.fetchUsers()
  },
  computed: {
    ...apiStatusComputedFactory(['fetchStatus'])
  },
  methods: {
    handleClickRow(row) {
      this.selectedUser = row
    },
    paginate(pageName) {
      this.params.offset = pageName
      this.fetchUsers()
    },
    async removeUser() {
      this.blockStatus = PENDING
      const { error } = await withAsync(deleteUser, {
        headers: { 'user-id': this.selectedUser.id }
      })
      if (error) {
        this.blockStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.blockStatus = SUCCESS
      this.$toast.success('Пользователь успешно удален!')
      this.fetchUsers()
    },
    async fetchUsers() {
      this.fetchRolesStatus = PENDING
      const { response, error } = await withAsync(getUsers)
      if (error) {
        this.fetchRolesStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.fetchRolesStatus = SUCCESS
      this.params.total = response.data.total
      this.users = response.data
    }
  }
}
</script>

<style lang="scss" scoped></style>
