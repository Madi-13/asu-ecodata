<template>
  <div>
    <div
      :class="[
        'tree__item',
        !!(selected && selected.id != node.id) && 'tree__item--blur'
      ]"
      :style="{ marginLeft: `${nodeMargin}rem` }"
    >
      <BaseCheckbox
        :name="node.value"
        @input="changeCheckState"
        :value="isChecked"
        color="green"
        class="mr--2"
      />
      <inline-svg
        class="tree__folder"
        :src="require(`@/assets/icons/folder-open.svg`)"
      />
      <span
        :class="['tree__text', node.isHighlight && 'tree__text--green']"
        @dblclick="onGotoEdit(node)"
      >
        {{ node.label }}
      </span>
      <inline-svg
        :class="['tree__arrow', node.isOpen && 'tree__arrow--open']"
        @click="toggleChildrenElements"
        v-if="Array.isArray(node.children)"
        :src="require(`@/assets/icons/shevron.svg`)"
      />
    </div>
    <div class="children" v-if="node.isOpen && node.children?.length">
      <node
        v-for="(child, index) in node.children"
        :key="node.id + index"
        :item="child"
        :is-filtered="isFiltered"
        @openEditModal="onGotoEdit(child)"
        :node-margin="nodeMargin + 3"
        @checked="$emit('checked', $event)"
        :selected="selected"
      />
    </div>
  </div>
</template>

<script>
import { getReferenceRecordsTree } from '@/api/referenceRecordsAPI'

import { withAsync } from '@/api/helpers/withAsync'
import { apiStatusComputedFactory } from '@/api/helpers/apiStatusComputedFactory'
import { apiStatus } from '@/api/constants/apiStatus'
const { IDLE, PENDING, SUCCESS, ERROR } = apiStatus

export default {
  name: 'node',
  emits: ['checked'],
  props: {
    item: {
      type: Object,
      default: () => ({})
    },
    nodeMargin: {
      type: Number,
      default: 0
    },
    selected: {
      type: Object
    },
    isFiltered: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      fetchTreeStatus: IDLE,
      isChecked: false,
      node: null
    }
  },
  computed: {
    ...apiStatusComputedFactory(['fetchTreeStatus']),
    getCurrentRefName() {
      return this.$route.params.reference_name
    }
  },
  watch: {
    item: {
      handler(newValue) {
        this.node = newValue
      },
      immediate: true
    },
    selected: {
      handler(newValue) {
        if (newValue && this.node.id != newValue.id) {
          this.isChecked = false
        }
      }
    }
  },
  methods: {
    onGotoEdit(node) {
      this.$emit('checked', {
        id: node.id
      })
      this.$emit('openEditModal')
    },
    changeCheckState() {
      this.$nextTick(() => {
        this.isChecked = !this.isChecked
        if (this.isChecked) {
          this.$emit('checked', {
            id: this.node.id
          })
        } else {
          this.$emit('checked', {})
        }
      })
    },
    async fetchTree() {
      this.fetchTreeStatus = PENDING
      const { response, error } = await withAsync(getReferenceRecordsTree, {
        table_name: this.getCurrentRefName,
        parent_id: this.node.id
      })
      if (error) {
        this.fetchTreeStatus = ERROR
        this.$toast.error(error.message)
        return
      }
      this.fetchTreeStatus = SUCCESS
      this.node.children = response.data.result
    },
    toggleChildrenElements() {
      this.node.isOpen = !this.node.isOpen
      if (this.node.children && !this.isFiltered) {
        this.fetchTree()
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.tree {
  &__text {
    font-weight: 400;
    cursor: pointer;
    font-size: 18px;
    color: #2b2b2b;
    &--green {
      background: #19a74e;
      display: inline;
      padding: 2px 4px;
      color: white;
    }
  }
  &__arrow {
    transform: rotate(180deg);
    transition: all 0.2s;
    margin-left: 1.5rem;
    &--open {
      transform: rotate(270deg);
    }
  }
  &__folder {
    margin-right: 1.5rem;
    width: 25px;
    height: 25px;
  }
  &__item {
    display: flex;
    align-items: center;
    padding: 1rem 0;
    &--blur {
      opacity: 0.4;
    }
  }
}
</style>
