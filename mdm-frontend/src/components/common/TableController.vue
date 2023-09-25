<template>
  <div class="operations mt--2">
    <BaseButton
      color="green"
      height="50px"
      v-can:[operation.permission_scope].hide="{
        resource: resourceName,
        permission: permissionName
      }"
      @click="$emit(operation.key)"
      :disabled="operation.is_need_selected && !selectedRow"
      :key="operation.key"
      v-for="operation in operations"
      >{{ operation.title }}
    </BaseButton>
  </div>
</template>

<script>
export default {
  props: {
    selectedRow: {
      type: Object
    },
    operations: { type: Array },
    permissionName: {
      type: String
    },
    resourceName: {
      type: String,
      default: 'references'
    }
  },
  data() {
    return {
      options: null
    }
  },
  watch: {
    operations: {
      handler(newValues) {
        this.options = newValues
      },
      immediate: true
    }
  },
  methods: {
    gotoPage(operationType) {
      switch (operationType) {
        case 'edit':
          this.$router.push({
            path: this.$route.path + '/edit/' + this.selectedRow.id,
            query: {
              reference_id: this.$route.query.table_id
            }
          })
          break
        case 'clone':
          this.$router.push({
            path: this.$route.path + '/clone/' + this.selectedRow.id,
            query: {
              reference_id: this.$route.query.table_id
            }
          })
          break
        case 'create':
          this.$router.push({
            path: this.$route.path + '/create',
            query: {
              reference_id: this.$route.query.table_id
            }
          })
          break
        default:
          break
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.controllers {
  margin: 2rem 0;
}

.operations {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(22rem, 1fr));
  gap: 1rem;
}
</style>
