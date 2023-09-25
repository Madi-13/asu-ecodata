<template>
  <paginate
    v-if="pageCount > 1"
    :page-count="pageCount"
    :page-range="2"
    :margin-pages="1"
    :click-handler="paginateHandler"
    :prev-text="'<'"
    :next-text="'>'"
    :container-class="'pagination'"
    :page-class="'pagination'"
    v-model="pagination.offset"
  >
  </paginate>
</template>

<script>
export default {
  props: {
    pagination: {
      type: Object,
      required: true
    }
  },
  computed: {
    pageCount() {
      return Math.ceil(this.pagination.total / this.pagination.limit)
    }
  },

  methods: {
    paginateHandler(page) {
      this.$emit('pageChange', page)
    }
  }
}
</script>

<style lang="scss" scoped>
.pagination {
  &::v-deep {
    display: flex;
    flex-direction: row;
    width: 100%;

    align-items: center;
    li {
      list-style-type: none;
      padding: 0;
      margin: 0;
      a {
        width: 32px;
        height: 32px;

        background: #ffffff;
        border: 1px solid #dfe3e8;
        border-radius: 4px;

        margin: 8px;
        display: flex;
        align-items: center;
        justify-content: center;

        font-weight: bold;
        font-size: 14px;
        line-height: 20px;
        &:hover {
          color: var(--green-light);
          border-color: var(--green-light);
        }
      }
    }
    .active a {
      color: var(--green);
      border-color: var(--green);
    }
  }
  margin-top: 20px;
  @media (max-width: 560px) {
    padding: 0 24px;
  }
}
</style>
