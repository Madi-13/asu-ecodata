import Vue from 'vue'

Vue.filter('moneyNormalizer', (value) => {
  if (value == null) return 'Не указан'
  else if (!value) return 0
  return String(value).replace(/(?<!\..*)(\d)(?=(?:\d{3})+(?:\.|$))/g, '$1 ')
})

Vue.filter('phoneNumberNormalizer', (value) => {
  if (value == null) return 'Не указан'
  else if (value.length == 12)
    return (
      '+7' +
      ' (' +
      value.slice(2, 5) +
      ') ' +
      value.slice(5, 8) +
      '-' +
      value.slice(8, 10) +
      '-' +
      value.slice(10, 12)
    )
  else return value
})
