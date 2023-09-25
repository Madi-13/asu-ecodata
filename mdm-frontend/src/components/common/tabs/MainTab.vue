<template>
  <div>
    <div v-if="type == 'large'">
      <ul class="tab__list">
        <li
          :class="[
            selectedTab.id != tab.id ? '' : 'tab__item--active',
            'tab__item'
          ]"
          @click="switchTab(tab.id, true)"
          v-for="(tab, index) in tabs"
          :key="index"
        >
          {{ tab.name }}
        </li>
      </ul>
      <slot></slot>
    </div>
    <div v-else-if="type == 'small'">
      <ul class="tab__list">
        <li
          :class="[
            selectedTab.id != tab.id ? '' : 'tab__item--sm--active',
            'tab__item--sm'
          ]"
          @click="switchTab(tab.id, false)"
          v-for="(tab, index) in tabs"
          :key="index"
        >
          {{ tab.name }}
        </li>
      </ul>
      <div class="tab__content">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    tabs: {
      type: Array,
      required: true
    },
    width: {
      type: String,
      default: 100
    },
    type: {
      type: String,
      default: 'large'
    }
  },
  data() {
    return {
      selectedTab: null
    }
  },
  watch: {
    $route: {
      handler(route) {
        if (route.hash) {
          this.switchTab(route.hash, false)
        } else {
          this.switchTab(this.tabs[0].id, false)
        }
      },
      immediate: true
    }
  },
  methods: {
    findTab(hash) {
      return this.tabs.find((tab) => {
        return tab.id == hash || `#${tab.id}` == hash
      })
    },
    cleanHash(hash) {
      return hash.replace(/^#/, '')
    },
    switchTab(hash, switchHash) {
      this.selectedTab = this.findTab(hash)
      if (this.selectedTab) {
        this.$emit('activateComponent', this.selectedTab.componentName)
      }
      if (switchHash) {
        this.$router.replace({
          hash: `#${this.selectedTab.id}`
        })
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.tab {
  &__list {
    display: flex;
    align-items: center;
    list-style: none;
    overflow: hidden;
    margin-bottom: 6px;
    margin-bottom: 3rem;
    overflow-x: auto;
    gap: 20px;
    -ms-overflow-style: none; /* IE and Edge */
    scrollbar-width: none; /* Firefox */
    &::-webkit-scrollbar {
      display: none;
    }
  }
  &__content {
    background: #ffffff;
    box-shadow: 0px 0px 10px rgb(43 43 43 / 20%);
    border-radius: 20px 20px 0px 0px;
  }
  &__item {
    font-size: 24px;
    color: #acacac;
    transition: all 0.2s;
    cursor: pointer;

    &--sm {
      font-weight: 500;
      font-size: 20px;
      color: #acacac;
      cursor: pointer;
      padding-bottom: 8px;
      transition: all 0.2s;
      position: relative;

      &--active {
        color: #0a7937;
        &::after {
          content: '';
          position: absolute;
          border-bottom: 3px solid #0a7937;
          bottom: 0;
          width: 100%;
          left: 0;
        }
      }
    }

    &--active {
      font-size: 32px;
      color: #2b2b2b;
      font-weight: 900;
    }
  }
}
</style>
