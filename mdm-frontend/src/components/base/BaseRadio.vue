<template>
  <div class="radio">
    <p class="radio__label mr--3" v-if="label != ''">{{ label }}</p>
    <div class="radio__controllers">
      <label
        :class="[
          'radio__wrapper',
          index == radioValues.length - 1 ? '' : 'mr--3'
        ]"
        v-for="(radio, index) in radioValues"
        :key="radio.name"
        :for="radio.name"
      >
        <input
          type="radio"
          :id="radio.name"
          class="radio__input"
          :name="radio.group"
          :value="radio.value"
          :checked="index === 0"
          @click="onClicked($event.target.value)"
        />
        <span class="checkmark"></span>
        <span>{{ radio.name }}</span>
      </label>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BaseRadio',
  props: {
    radioValues: {
      type: Array,
      default: () => []
    },
    label: {
      type: String,
      default: ''
    },
    value: {
      type: String,
      required: true
    }
  },
  methods: {
    onClicked(value) {
      this.$emit('input', value)
    }
  }
}
</script>

<style scoped lang="scss">
.radio {
  display: flex;
  flex-direction: column;
  width: 100%;
  margin-bottom: 2rem;
  &__controllers {
    display: flex;
    align-items: center;
  }
  &__label {
    margin-bottom: 1.4rem;
  }
  &__input {
    opacity: 0;
    width: 0;
    height: 0;
  }
  &__wrapper {
    display: flex;
    align-items: center;
    font-weight: 400;
    font-size: 16px;
    color: #2b2b2b;
    display: flex;
    align-items: center;
    cursor: pointer;
    width: fit-content;
    width: max-content;
  }
}
.checkmark {
  background: linear-gradient(109.8deg, #dee6ec 0%, #fcfdfe 100%);
  box-shadow: -2px -2px 4px #ffffff, 2px 2px 4px #d1dce3;
  width: 2rem;
  height: 2rem;
  display: block;
  border-radius: 50%;
  margin-right: 1.5rem;
  position: relative;
  transition: all 0.3s;
  &::before {
    content: '';
    width: 2rem;
    height: 2rem;
    background: linear-gradient(109.8deg, #dee6ec 0%, #fcfdfe 100%);
    position: absolute;
    top: 50%;
    left: 50%;
    border-radius: 50%;
    transform: translate(-50%, -50%);
  }
}
.radio__input:checked {
  & ~ .checkmark {
    background: linear-gradient(84.07deg, #42d814 19.97%, #47ce1e 90.11%);
    box-shadow: inset 2px 2px 4px #4ace38, inset -2px -2px 4px #36b11d;
    position: relative;
    &::before {
      content: '';
      width: 1rem;
      height: 1rem;
      background: linear-gradient(84.07deg, #15bd2c 19.97%, #2dda2d 90.11%);
      position: absolute;
      top: 50%;
      left: 50%;
      border-radius: 50%;
      transform: translate(-50%, -50%);
    }
  }
}
</style>
