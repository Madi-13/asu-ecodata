<template>
  <aside class="sidebar">
    <img
      src="../../../assets/icons/logo_header.svg"
      alt="Logo"
      class="header_logo"
    />
    <nav class="nav" v-if="navs">
      <PuSkeleton :count="5" height="40px" v-if="isLoading"></PuSkeleton>
      <ul class="nav__list mt--5" v-else>
        <li
          class="nav__item"
          v-for="(nav, idx) in navs"
          :key="nav.routeTo.name"
          v-can:read.hide="{ resource: nav.resourceName }"
        >
          <router-link
            :to="{
              name: nav.routeTo.name
            }"
            @click.prevent="openSubNavs(idx, nav.refType)"
            :is="nav.childrens ? 'span' : 'router-link'"
            class="nav__link"
          >
            <inline-svg
              :src="require(`@/assets/icons/${nav.iconName}.svg`)"
              class="mr--2 nav__icon"
            />
            {{ nav.name }}
            <inline-svg
              v-if="nav.childrens"
              :src="require(`@/assets/icons/shevron.svg`)"
              :class="['nav__toggle', nav.isOpen ? 'nav__toggle--open' : '']"
            />
          </router-link>
          <FromRightSideTransition>
            <ul
              class="nav__list nav__list--inner"
              v-if="nav.hasSubMenu && nav.isOpen"
            >
              <template v-if="nav.childrens.length > 0">
                <router-link
                  :to="{
                    name: child.routeTo.name,
                    ...(child.routeTo.params && {
                      params: child.routeTo.params,
                      query: child.routeTo.query
                    })
                  }"
                  class="nav__item--inner nav__link"
                  v-for="child in nav.childrens"
                  :key="child.routeTo.name"
                >
                  {{ child.name }}
                </router-link>
              </template>
              <li class="nav__item--inner nav__link" v-else>Нет данных</li>
            </ul>
          </FromRightSideTransition>
        </li>
        <li class="nav__link mt--3" @click="logout">
          <inline-svg
            :src="require(`@/assets/icons/align-text-top.svg`)"
            class="mr--2 nav__icon"
          />Выход
        </li>
      </ul>
    </nav>
  </aside>
</template>

<script>
import FromRightSideTransition from '@/components/transitions/FromRightSideTransition.vue'
import { mapMutations } from 'vuex'

import { navsComputed, toggleMenu } from '@/services/navs.js'

export default {
  components: {
    FromRightSideTransition
  },
  data() {
    return {
      isLoading: false
    }
  },
  computed: {
    ...navsComputed
  },
  methods: {
    ...mapMutations({
      logoutMutation: 'SET_LOGOUT'
    }),
    openSubNavs(idx, refType) {
      toggleMenu(idx, refType)
    },
    logout() {
      this.logoutMutation()
      this.$router.push({ name: 'Login' })
    }
  }
}
</script>

<style lang="scss" scoped>
.sidebar {
  padding: 4.3rem 0 4.3rem 4rem;
  height: 100vh;
  width: 25%;
  overflow-y: auto;
}
.nav {
  &__list {
    list-style: none;
    font-weight: 400;
    overflow: hidden;
    padding-right: 0.5rem;
    font-size: 20px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;

    &--inner {
      margin-left: 4.7rem;
    }
  }
  &__item {
    cursor: pointer;
    display: flex;
    align-items: flex-start;
    flex-direction: column;
    width: 100%;

    &:not(:first-child) {
      margin-top: 3rem;
    }
    &--inner {
      margin-top: 1rem;
      &:not(:first-child) {
        margin-top: 1rem;
      }
    }
  }
  &__link {
    text-decoration: none;
    display: flex;
    align-items: center;
    cursor: pointer;
    color: #0a7937;
    opacity: 0.5;
    width: 100%;
  }
  &__toggle {
    transform: rotate(-180deg);
    transition: all 0.3s;
    margin-left: auto;
    &--open {
      transform: rotate(-90deg);
    }
  }
}
a.router-link-active {
  font-weight: 700;
  opacity: 1;
}
</style>
