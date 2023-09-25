<template>
  <MainTab
    type="small"
    :tabs="recordEditTabs"
    width="70%"
    ref="tab"
    @activateComponent="onActivateComponent($event)"
  >
    <FadeTransition>
      <component :is="activeComponentName"></component>
    </FadeTransition>
  </MainTab>
</template>

<script>
import MainTab from '@/components/common/tabs/MainTab.vue'
import FadeTransition from '@/components/transitions/FadeTransition.vue'

export default {
  components: {
    MainTab,
    FadeTransition,
    ReferenceRecordForm: () => import('./ReferenceRecordForm.vue'),
    ReferenceRecordHistory: () => import('./ReferenceRecordHistory.vue')
  },
  data() {
    return {
      activeComponentName: 'ReferenceRecordHistory',
      recordEditTabs: [
        {
          name: 'Изменить сведения',
          componentName: 'ReferenceRecordForm',
          id: 'recordForm'
        },
        {
          name: 'История изменений',
          componentName: 'ReferenceRecordHistory',
          id: 'recordHistory'
        }
      ]
    }
  },
  methods: {
    onActivateComponent(name) {
      this.activeComponentName = name
    }
  }
}
</script>

<style lang="scss" scoped></style>
