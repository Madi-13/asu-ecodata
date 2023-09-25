<template>
    <form class="auth__form">
        <BaseInput
          label="Сумма оплаты ₸"
          placeholder="Внесите сумму оплаты"
          :width="30"
          v-model="form.sum"
          class="bill-form__total"
          size="large"
        ></BaseInput>
    </form>
</template>

<script>
    export default {
        data() {
            return {
                form: {
                    username:null,
                    password:null
                }
            }
        },
    }
</script>

<style lang="scss" scoped>

</style>