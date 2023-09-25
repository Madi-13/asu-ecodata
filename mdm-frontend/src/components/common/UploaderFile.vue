<template>
  <div>
    <p :class="['uploader__label', disabled && 'uploader__label--disabled']">
      {{ label }}
    </p>
    <label id="uploader" :class="['uploder__input']">
      <p
        :class="[
          'uploader__value',
          disabled && 'uploader__value--disabled',
          'uploader__value--placeholder'
        ]"
      >
        <template v-if="!value || value.length == 0">
          {{ placeholder }}
        </template>
        <template v-else>
          <div class="chips">
            <template v-if="isMultiple">
              <Chips
                v-for="(item, idx) in value"
                :key="item.file_name"
                @clickOnChip="downloadFile(item)"
                @remove="$emit('onRemove', idx)"
                :title="item.file_name"
              />
            </template>
            <template v-else>
              <Chips
                :key="value.file_name"
                @clickOnChip="downloadFile(item)"
                @remove="$emit('onRemove')"
                :title="value.file_name"
              />
            </template>
          </div>
        </template>
      </p>
      <input
        type="file"
        id="uploader"
        class="uploader__ctrl"
        :disabled="disabled"
        :accept="acceptFileFormat"
        @change="uploadFile($event.target)"
      />
    </label>
  </div>
</template>

<script>
import Chips from '@/components/common/Chips.vue'
import { b64toBlob } from '@/helpers/fileManipulations'

export default {
  components: {
    Chips
  },
  props: {
    label: {
      type: String,
      default: 'Файл загрузки'
    },
    disabled: {
      type: Boolean,
      default: false
    },
    placeholder: {
      type: String,
      default: 'Выберите файл'
    },
    isMultiple: {
      type: Boolean,
      default: false
    },
    acceptFileFormat: {
      type: String,
      default:
        'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.ms-excel'
    },
    value: {
      type: [Array, Object, String],
      default: () => []
    }
  },
  methods: {
    downloadFile(file) {
      const blob = b64toBlob(file.file_body)
      const link = document.createElement('a')
      let url = window.URL.createObjectURL(blob)
      link.href = url
      link.download = file.file_name
      link.click()
      link.remove()
      window.URL.revokeObjectURL(url)
    },
    uploadFile(eventTarget) {
      if (eventTarget.files.length == 0) return
      let fileName = eventTarget.files[0].name
      const reader = new FileReader()
      reader.onloadend = () => {
        const base64String = reader.result
          .replace('data:', '')
          .replace(/^.+,/, '')
        eventTarget.value = ''
        this.$emit(
          'input',
          this.isMultiple
            ? [
                ...this.value,
                {
                  file_name: fileName,
                  file_body: base64String
                }
              ]
            : { file_name: fileName, file_body: base64String }
        )
        this.$emit('change', {
          file_name: fileName,
          file_body: base64String
        })
      }
      reader.readAsDataURL(eventTarget.files[0])
    },
    b64DecodeUnicode(str) {
      return decodeURIComponent(
        window
          .atob(str)
          .split('')
          .map(function (c) {
            return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)
          })
          .join('')
      )
    }
  }
}
</script>

<style lang="scss" scoped>
.chips {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
}
.uploader {
  &__ctrl {
    opacity: 0;
    position: absolute;
    z-index: -1;
  }
  &__value {
    border-bottom: 2.7113px solid #2b2b2b;
    padding: 20px 0 10px;
    width: 100%;
    transition: all 0.3s;
    cursor: pointer;
    font-weight: 400;
    padding: 20px 0 10px;
    font-size: 16px;
    &--placeholder {
      color: rgba(17, 17, 17, 0.5);
    }
    &--disabled {
      color: rgba(17, 17, 17, 0.5);
      border-bottom: 2.7113px solid rgba(17, 17, 17, 0.5);
      cursor: not-allowed;
    }
  }
  &__label {
    font-weight: 400;
    font-size: 16px;
    color: #2b2b2b;
    width: 100%;
    transition: all 0.3s;
    &--disabled {
      color: rgba(17, 17, 17, 0.5);
    }
  }
}
</style>
