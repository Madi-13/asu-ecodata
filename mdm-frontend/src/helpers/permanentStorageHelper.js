export function getDataFromLocalStorage(key) {
  return localStorage.getItem(key)
}

export function setDataToLocalStorage(key, value) {
  const valueAsString = typeof value == 'object' ? JSON.stringify(value) : value
  localStorage.setItem(key, valueAsString)
}

export function clearDataLocalStorage(key) {
  localStorage.removeItem(key)
}
