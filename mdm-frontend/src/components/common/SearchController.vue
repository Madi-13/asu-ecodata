<template>
  <div class="search">
    <div class="search__controller">
      <input
        type="text"
        v-bind="{ ...$attrs }"
        v-if="selectedFilterField?.frontend_type == 'datepicker'"
        v-mask="'####.##.##'"
        class="search__input"
        v-model="searchValue"
      />
      <input
        type="text"
        v-bind="{ ...$attrs }"
        v-else
        class="search__input"
        v-model="searchValue"
      />
      <img src="../../assets/icons/search.svg" alt="" class="search__icon" />
    </div>
    <BaseButton color="outline-dark" @click="onSearch">Поиск</BaseButton>
    <BaseButton
      color="outline-dark"
      @click="openFilterList($event)"
      v-if="filterableColumns && filterableColumns.length > 0"
      v-click-outside="close"
    >
      <img src="../../assets/icons/setting-config.svg" alt="" />
    </BaseButton>
    <ul
      class="filter__list"
      :style="{ top: `${listTopPosition}px`, left: `${listLeftPosition}px` }"
      v-if="filterableColumns && isListOpen"
    >
      <li
        :class="[
          'filter__item',
          searchBy == option.col_name ? 'filter__item--active' : ''
        ]"
        v-for="option in filterableColumns"
        :key="option.col_name"
        @click="onSelect(option)"
      >
        {{ option.title }}
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  props: {
    filterOptions: {
      type: Array
    }
  },
  data() {
    return {
      searchValue: null,
      searchBy: null,
      listTopPosition: null,
      listLeftPosition: null,
      selectedFilterField: null,
      isListOpen: false
    }
  },
  computed: {
    filterableColumns() {
      return (
        this.filterOptions &&
        this.filterOptions.filter((el) => el.is_filterable)
      )
    }
  },
  methods: {
    close() {
      this.isListOpen = false
    },
    onSelect(option) {
      this.isListOpen = false
      this.selectedFilterField = option
      this.searchBy = option.col_name
    },
    onSearch() {
      this.$emit('searchClicked', {
        search_value: this.searchValue,
        search_by: this.searchBy
      })
    },
    openFilterList(event) {
      let controllerPosition = event.target.getBoundingClientRect()
      this.listTopPosition =
        controllerPosition.y + controllerPosition.height + 10
      this.listLeftPosition = controllerPosition.x
      this.isListOpen = true
    }
  }
}
</script>

<style lang="scss" scoped>
.search {
  display: grid;
  grid-template-columns: 40% 20% 4%;
  gap: 1rem;
  &__input {
    border: 2px solid #2b2b2b;
    border-radius: 10px;
    padding: 15px;
    padding-left: 75px;
    font-weight: 400;
    font-size: 16px;
    color: #2b2b2b;
    height: 100%;
    position: relative;
    width: 100%;
    &::placeholder {
      color: #2b2b2b;
      font-weight: 400;
      opacity: 0.5;
    }
  }
  &__controller {
    position: relative;
  }
  &__icon {
    position: absolute;
    top: 50%;
    left: 30px;
    transform: translateY(-50%);
  }
  &__button {
    border: 2px solid #2b2b2b;
    border-radius: 10px;
    display: flex;
    width: 14rem;
    align-items: center;
    background-color: white;
    padding: 15px 15px;
    justify-content: center;
    font-weight: 400;
    font-size: 16px;
    color: #2b2b2b;
    cursor: pointer;
    margin-left: 33px;
  }
}
.filter {
  &__list {
    position: fixed;
    list-style: none;
    border: 2px solid #2b2b2b;
    border-radius: 10px;
    outline: none;
    font-weight: 400;
    background-color: white;
    overflow: hidden;
    font-size: 16px;
    z-index: 12;
    color: #2b2b2b;
  }
  &__item {
    min-height: 40px;
    background-color: white;
    cursor: pointer;
    display: flex;
    outline: none;
    align-items: center;
    padding: 0 15px;

    &--active {
      background-color: #332c2c;
      color: white;
    }

    &:hover {
      background-color: #a19d9d;
      color: #2b2b2b;
    }
  }
}
</style>
