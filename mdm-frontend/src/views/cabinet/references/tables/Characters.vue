<template>
  <div>
    <el-table highlight-current-row empty-text="Данных нет" :data="form" border>
      <el-table-column label="Характеристика" prop="name"></el-table-column>
      <el-table-column label="Значение">
        <template slot-scope="scope">
          <div class="display--row">
            <BaseColumnDropdown
              placeholder="Выберите из списка"
              v-model="scope.row.selected"
              optionValueField="key"
              optionTitleField="value"
              :options="scope.row.options"
              :width="100"
            />
            <button
              @click="openCreateAttrModal(scope.row.id)"
              title="Создать новое значение"
              class="plus"
            >
              +
            </button>
          </div>
        </template>
      </el-table-column>
    </el-table>
    <CreateAttrModal
      :attribute-id="selectedAttrID"
      @onClose="$emit('addedNewAttr', { id: selectedAttrID, attr: $event })"
    />
  </div>
</template>

<script>
import CreateAttrModal from '../modals/CreateAttrModal.vue'
export default {
  components: {
    CreateAttrModal
  },
  props: {
    attribute: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      form: null,
      selectedAttrID: null
    }
  },
  watch: {
    attribute: {
      handler(newValue) {
        this.form = newValue
      },
      immediate: true,
      deep: true
    }
  },
  methods: {
    openCreateAttrModal(attrID) {
      this.selectedAttrID = attrID
      this.$modal.show('createAttrModal')
    }
  }
}
</script>

<style lang="scss" scoped>
.plus {
  font-size: 1.5rem;
  background: none;
  border: none;
  cursor: pointer;
  padding: 0.2rem 1rem;

  &:hover {
    background: #dfe3e8;
  }
}
</style>
