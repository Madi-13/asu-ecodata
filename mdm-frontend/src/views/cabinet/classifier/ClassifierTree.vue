<template>
  <div>
    <div class="sticky">
      <p class="title--secondary mb--2">
        {{ `Классификатор "${tableShowName}"` }}
      </p>
      <SearchController @searchClicked="onSearch($event)" placeholder="Найти" />
      <TableController
        :selected-row="selectedRow"
        resource-name="references"
        :permission-name="currentRefName"
        :operations="operations"
        @create-attr="openModal('createAttrModal')"
        @edit-attr="openModal('editAttrModal')"
        @create="gotoPage('ReferenceRecordCreate')"
        @edit="gotoPage('ReferenceRecordEdit')"
        @delete="removeCategory"
        @export="exportData"
        class="mb--2"
      />
    </div>
    <template v-if="fetchTreeStatusSuccess && !data">
      <p class="title--secondary">Данных нет</p>
    </template>
    <template v-else>
      <TreeView
        v-for="(item, idx) in data"
        :key="idx"
        :item="item"
        :is-filtered="!!searchValue"
        @checked="setSelectedNode($event)"
        @openEditModal="gotoPage('ReferenceRecordEdit')"
        :selected="selectedRow"
      />
    </template>

    <CreateAttrModal :classifier-id="selectedRow && selectedRow.id" />
    <EditAttrModal :classifier-id="selectedRow && selectedRow.id" />
  </div>
</template>

<script>
import TableController from '@/components/common/TableController.vue'
import SearchController from '@/components/common/SearchController.vue'
import TreeView from './components/TreeView.vue'
import CreateAttrModal from './components/CreateAttrModal.vue'
import EditAttrModal from './components/EditAttrModal.vue'

import {
  getReferenceRecordsTree,
  deleteRecord
} from '@/api/referenceRecordsAPI'
import { downloadFile } from '@/api/importAPI'

import { operationsClassifierTable } from '@/fixtures/options'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'

const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

export default {
  components: {
    TableController,
    SearchController,
    TreeView,
    CreateAttrModal,
    EditAttrModal
  },
  data() {
    return {
      fetchTreeStatus: IDLE,
      data: null,
      selectedRow: null,
      currentRefName: null,
      tableShowName: null,
      searchValue: null,
      operations: operationsClassifierTable
    }
  },
  computed: {
    ...apiStatusComputedFactory(['fetchTreeStatus'])
  },
  watch: {
    $route: {
      handler(newValue) {
        this.tableShowName = newValue.query && newValue.query.table_show
        this.currentRefName = newValue.params && newValue.params.reference_name
        this.fetchTree()
      },
      deep: true,
      immediate: true
    }
  },
  methods: {
    gotoPage(pageName) {
      this.$router.push({
        name: pageName,
        query: {
          reference_id: this.$route.query.table_id,
          is_classifier: true
        },
        params: {
          reference_name: this.currentRefName,
          row_id: this.selectedRow && this.selectedRow.id
        }
      })
    },
    openModal(modalName) {
      this.$modal.show(modalName)
    },
    setSelectedNode(selectedNode) {
      this.selectedRow = selectedNode
    },
    async exportData() {
      const { response, error } = await withAsync(downloadFile, {
        import_type_code: 'CLASSIFICATION',
        table_name: this.currentRefName
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
    onSearch(searchObj) {
      this.searchValue = searchObj.search_value
      this.fetchTree()
    },
    async fetchTree() {
      this.fetchTreeStatus = PENDING
      const { response, error } = await withAsync(getReferenceRecordsTree, {
        table_name: this.currentRefName,
        [this.searchValue ? 'search' : 'parent_id']: this.searchValue || null
      })
      if (error) {
        this.fetchTreeStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.fetchTreeStatus = SUCCESS
      this.data = this.searchValue
        ? this.prettifyResponse(response.data.result, this.searchValue)
        : response.data.result
    },
    prettifyResponse(tree, value) {
      let pattern = /value/gi

      let childrens = tree.map((item) => {
        if (pattern.test(item.label) || item.children == null) {
          item.isHighlight = true
          item.isOpen = true
          return item
        }
        item.isHighlight = false
        item.isOpen = true
        return this.prettifyResponse(item.children, value)
      })
      tree.children = childrens
      return tree
    },
    async removeCategory() {
      this.removeCategoryStatus = PENDING
      const { error } = await withAsync(deleteRecord, {
        table_name: this.currentRefName,
        operation_type: 'drop',
        row_id: this.selectedRow.id
      })
      if (error) {
        this.removeCategoryStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.removeCategoryStatus = SUCCESS
      this.$toast.success('Категория успешно удалена!')
      this.fetchTree()
      this.selectedRow = null
    }
  }
}
</script>

<style lang="scss" scoped></style>
