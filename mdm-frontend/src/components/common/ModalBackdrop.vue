<template>
  <FadeTransition>
    <div class="modal__backdrop-wrapper" v-if="visible">
      <div class="modal__backdrop" @click.self="onClose"></div>
      <div class="modal">
        <h3 class="modal-title" v-if="title">{{ title }}</h3>
        <slot></slot>
      </div>
    </div>
  </FadeTransition>
</template>

<script>
import FadeTransition from '@/components/transitions/FadeTransition.vue'

export default {
  components: {
    FadeTransition
  },
  props: {
    name: {
      type: String,
      required: true
    },
    title: {
      type: String
    }
  },
  data() {
    return {
      visible: false,
      params: {}
    }
  },
  created() {
    const escapeHandler = (e) => {
      if (e.key == 'Escape') {
        this.onClose()
      }
    }
    document.addEventListener('keydown', escapeHandler)
    this.$once('hook:destroyed', () => {
      document.removeEventListener('keydown', escapeHandler)
    })
  },
  watch: {
    visible: {
      handler(newValue) {
        newValue
          ? (document.body.style.overflow = 'hidden')
          : (document.body.style.overflow = 'initial')
      },
      immediate: true
    }
  },
  beforeMount() {
    this.$modal.$event.$on('show', (name, params) => {
      if (this.name == name) {
        this.visible = true
        this.$emit('onOpen')
        this.params = params
      }
    })
    this.$modal.$event.$on('hide', (name) => {
      if (this.name == name) {
        this.visible = false
        this.$emit('onClose')
      }
    })
  },
  methods: {
    onClose() {
      this.$modal.hide(this.name)
    },
    onSaveClicked() {
      this.$emit('saveTriggered')
    }
  }
}
</script>

<style lang="scss" scoped>
.modal {
  width: 55%;
  overflow-y: auto;
  background-color: white;
  z-index: 23;
  position: absolute;
  padding: 3rem 2rem;
  top: 50%;
  max-height: 66%;

  left: 50%;
  transform: translate(-50%, -50%);

  &-title {
    margin-bottom: 5rem;
  }

  &__content {
    padding: 24px;
  }

  &__backdrop {
    background: rgba(#000000, 0.5);
    transition: all 0.3s;
    position: fixed;
    height: 100vh;
    width: 100vw;
    z-index: 10;
    top: 0;
    left: 0;
  }
}
</style>
