<template>
    <header class="header">
        <p class="header__pagename" v-if="currentRouteTitle">
            {{currentRouteTitle}}
        </p>
        <div class="go-back" @click="goBack" v-else>
            <img src="../../../assets/icons/shevron.svg" alt="">
            <p>Назад</p>
        </div>
        <p class="header__info">
            {{ getUsername }}
        </p>
    </header>
</template>

<script>
import {  mapGetters } from 'vuex'
export default {
    computed: {
        ...mapGetters(['getUsername']),
        currentRouteTitle() {
            return this.$route.meta.title
        },
    },
    methods: {
        goBack() {
            this.$router.go(-1)
        }
    },
}
</script>

<style lang="scss" scoped>
.header{
    display: flex;
    justify-content: space-between;;
    margin-bottom: 5rem;
    &__info{
        font-weight: 400;
        font-size: 16px;
        color: #2B2B2B;
    }
    &__pagename{
        font-weight: 900;
        font-size: 28px;
        text-transform: uppercase;
        color: #2B2B2B;
    }
}
.go-back{
    display: flex;
    align-items: center;
    cursor: pointer;
    gap: 1rem
}
</style>