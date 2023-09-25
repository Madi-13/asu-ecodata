<template>
  <BaseColumnKeyDropdown
    :option-title-field="optionTitleField"
    :option-value-field="optionValueField"
    :options="filteredBySearchKey"
    :value="value"
    @selected="selectHandler"
  >
    <template #input="{toggle}">
      <div class="inputs">
        <div class="select">
          <BaseInput :has-validation="false" @focus="toggle" name="search" placeholder="введите" v-model="search" />
          <img class="arrow" src="@/assets/icons/chevron_down.svg" alt="">
        </div>
        <button @click="addHandler" title="Создать новое значение" class="plus">+</button>
      </div>
    </template>
  </BaseColumnKeyDropdown>
</template>

<script>
import BaseColumnKeyDropdown from "@/components/base/BaseColumnKeyDropdown.vue";
import BaseInput from "@/components/base/BaseInput.vue";

export default {
  components: { BaseInput, BaseColumnKeyDropdown },
  model: {
    event: "select",
    prop: "value"
  },
  props: {
    optionTitleField: [String, Number],
    optionValueField: [String, Number],
    list: Array,
    value: String
  },
  data() {
    return {
      search: ""
    };
  },
  computed: {
    filteredBySearchKey() {
      return this.list.filter(item => item[this.optionTitleField]?.toLowerCase().includes(this.search?.toLowerCase()));
    },
    findItem() {
      return this.list.find(item => item[this.optionValueField] === this.value);
    }
  },
  watch: {
    value: {
      handler(key) {
        if (this.findItem) {
          this.search = this.findItem[this.optionTitleField];
        }
      },
      immediate: true
    }
  },
  methods: {
    addHandler() {
      this.$emit("add", this.search);
    },
    selectHandler(key) {
      if (this.findItem) {
        this.search = this.findItem[this.optionTitleField];
      }
      this.$emit("select", key);
    }
  }
};
</script>

<style lang="scss" scoped>
.inputs {
  display: flex;
  gap: 1rem;
  align-items: center;
}

.plus {
  font-size: 1.5rem;
  background: none;
  border: none;
  cursor: pointer;
  padding: .2rem 1rem;

  &:hover {
    background: #dfe3e8;
  }
}

.select {
  position: relative;
  width: 30rem;

  .arrow {
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
    right: 1rem;
  }
}
</style>
