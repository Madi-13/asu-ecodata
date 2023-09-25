import { helpers } from 'vuelidate/lib/validators'

export function containsOnlyLatin(value) {
  return /^[a-zA-Z]*$/.test(value)
}

export function mustStartWithCyrillicLetter(value) {
  const cyrillicPattern = /^[\u0400-\u04FFa-zA-Z0-9\s\_]*$/
  return cyrillicPattern.test(value)
}
