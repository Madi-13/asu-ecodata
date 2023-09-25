<template>
  <div class="auth">
    <div class="auth__background"></div>
    <div class="auth__block">
      <div class="auth__logo display--row">
        <img src="../../assets/icons/logo.svg" alt="" />
        <img src="../../assets/icons/logo_text.svg" alt="" class="ml--md" />
      </div>
      <div class="auth__tabs">
        <MainTab
          :tabs="authTabs"
          width="70%"
          ref="tab"
          @activateComponent="onActivateComponent($event)"
        >
          <FadeTransition>
            <component :is="activeComponentName"></component>
          </FadeTransition>
        </MainTab>
      </div>
    </div>
  </div>
</template>

<script>
import MainTab from '@/components/common/tabs/MainTab.vue'
import FadeTransition from '@/components/transitions/FadeTransition.vue'
export default {
  components: {
    MainTab,
    FadeTransition,
    SignInForm: () => import(`./components/SignInForm.vue`),
    SignUpForm: () => import(`./components/SignUpForm.vue`)
  },
  data() {
    return {
      authTabs: [
        {
          name: 'Вход',
          componentName: 'SignInForm',
          id: 'signin'
        }
        // {
        //   name: 'Регистрация',
        //   componentName: 'SignUpForm',
        //   id: 'signup'
        // }
      ],
      activeComponentName: 'SignInForm'
    }
  },
  methods: {
    onActivateComponent(name) {
      this.activeComponentName = name
    }
  }
}
</script>

<style lang="scss" scoped>
.auth {
  width: 100%;
  display: flex;

  &__background {
    width: 50%;
    background: linear-gradient(212.93deg, #001b0b -17.74%, #002f13 101.27%);
    box-shadow: 0px 0px 10px rgba(1, 69, 28, 0.5);
    border-radius: 0px;
    min-height: 100vh;
  }
  &__logo {
    position: absolute;
    z-index: 1;
    top: 10%;
  }
  &__block {
    width: 50%;
    min-height: 100vh;
    display: flex;
    align-items: flex-start;
    padding: 9.5rem 9%;
    justify-content: center;
    flex-direction: column;
    position: relative;
  }
  &__tabs {
    width: 100%;
  }
}
</style>
