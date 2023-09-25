<template>
  <div>
    <div class="event__header">
      {{ event.date }}
    </div>
    <div class="event__body">
      <div class="event__info">
        <b class="pl--2">{{ event.time }}</b>
        <p class="event__message">{{ event.message }}</p>
      </div>
      <div
        class="event__column mt--2"
        v-if="event.new_value || event.old_value"
      >
        <b>Значения полей:</b>
        <div class="event__content">
          <div class="event__item">
            <b class="event__item-title">Текущие данные</b>
            <template v-if="event.new_value">
              <div
                class="event__field"
                v-for="value in event.new_value"
                :key="value.column_name"
              >
                <span class="event__cell">{{ value.column_name }}</span>
                <b class="event__cell text--center">{{ value.column_value }}</b>
              </div>
            </template>
          </div>
          <div class="event__item">
            <b class="event__item-title">Ранее использованные данные</b>
            <template v-if="event.old_value">
              <div
                class="event__field"
                v-for="value in event.old_value"
                :key="value.column_name"
              >
                <span class="event__cell">{{ value.column_name }}</span>
                <b class="event__cell text--center">{{ value.column_value }}</b>
              </div>
            </template>
          </div>
        </div>
      </div>
      <div
        class="event__column"
        v-if="
          event.dictionary_attributes_new_values ||
          event.dictionary_attributes_old_values
        "
      >
        <b>Атрибуты:</b>
        <div class="event__content">
          <div class="event__item">
            <b class="event__item-title">Текущие данные</b>
            <template v-if="event.dictionary_attributes_new_values">
              <div
                class="event__field"
                v-for="value in event.dictionary_attributes_new_values"
                :key="value.attribute_name"
              >
                <span class="event__cell">{{ value.attribute_name }}</span>
                <b class="event__cell text--center"
                  >{{ value.attribute_value }}
                </b>
              </div>
            </template>
          </div>
          <div class="event__item">
            <b class="event__item-title">Ранее использованные данные</b>
            <template v-if="event.dictionary_attributes_old_values">
              <div
                class="event__field"
                v-for="value in event.dictionary_attributes_old_values"
                :key="value.attribute_name"
              >
                <span class="event__cell">{{ value.attribute_name }}</span>
                <b class="event__cell text--center"
                  >{{ value.attribute_value }}
                </b>
              </div>
            </template>
          </div>
        </div>
      </div>
      <div
        class="event__column"
        v-if="event.attributes_new_values || event.attributes_old_values"
      >
        <b>Значения атрибутов:</b>
        <div class="event__content">
          <div class="event__item">
            <b class="event__item-title">Текущие данные</b>
            <template v-if="event.attributes_new_values">
              <div class="event__field">
                <span class="event__cell">{{
                  event.attributes_new_values['attribute_name']
                }}</span>
                <b class="event__cell text--center">
                  <template
                    v-for="(attr, idx) in event.attributes_new_values
                      .attribute_list"
                  >
                    {{ attr.value }}
                    {{
                      idx + 1 !=
                      event.attributes_new_values.attribute_list.length
                        ? ', '
                        : ''
                    }}
                  </template>
                </b>
              </div>
            </template>
          </div>
          <div class="event__item">
            <b class="event__item-title">Ранее использованные данные</b>
            <template v-if="event.attributes_old_values">
              <div
                class="event__field"
                v-for="value in event.attributes_old_values"
                :key="value.attribute_name"
              >
                <span class="event__cell">{{ value.attribute_name }}</span>
                <b class="event__cell text--center">
                  <template v-for="(attr, idx) in value.attribute_list">
                    {{ attr.value }}
                    {{ idx + 1 != value.attribute_list.length ? ', ' : '' }}
                  </template>
                </b>
              </div>
            </template>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    event: {
      type: Object
    }
  }
}
</script>

<style lang="scss" scoped>
.event {
  &__header {
    background: #e9e9e9;
    padding: 12px 0 12px 2rem;
  }
  &__info {
    display: grid;
    align-items: center;
    padding: 12px 0;
    grid-template-columns: 15% 85%;
  }
  &__field {
    display: grid;
    grid-template-columns: 60% 40%;
  }
  &__content {
    display: grid;
    grid-template-columns: repeat(2, 49%);
    margin-bottom: 2rem;
    justify-content: space-between;
  }
  &__item {
    &-title {
      margin-bottom: 8px;
      display: inline-block;
    }
  }
  &__cell {
    border: 1px solid #acacac;
    padding: 6px 13px;
  }
  &__column {
    display: grid;
    align-items: flex-start;
    grid-template-columns: 15% 85%;
    padding: 0 15px;
  }
}
</style>
