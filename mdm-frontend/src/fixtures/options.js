const newOperation = (
  title,
  key,
  is_need_selected,
  permission_scope = null
) => ({ title, key, is_need_selected, permission_scope })

export const agreementOptions = [
  {
    name: 'Да',
    id: 'Да'
  },
  {
    name: 'Нет',
    id: 'Нет'
  }
]

export const tableTypes = [
  {
    title: 'Справочники',
    id: 1
  },
  {
    title: 'Классификаторы',
    id: 2
  }
]

export const tableTypesForImport = [
  {
    title: 'Справочники',
    id: 1,
    code: 'DICTIONARY'
  },
  {
    title: 'Классификаторы',
    id: 2,
    code: 'CLASSIFICATION'
  },
  {
    title: 'Атрибуты',
    id: null,
    code: 'ATTRIBUTE'
  },
  {
    title: 'Значения атрибутов',
    id: null,
    code: 'ATTRIBUTE_VALUE'
  }
]

export const operationsRefTable = [
  newOperation('Создать запись', 'create', false, 'write'),
  newOperation('Клонировать запись', 'clone', true, 'write'),
  newOperation('Изменить запись', 'edit', true, 'read'),
  newOperation('Сменить статус', 'set-state', true, 'write'),
  newOperation('Удалить запись', 'delete', true, 'delete'),
  newOperation('Реплицировать запись', 'replicate', true, 'write'),
  newOperation('Экспорт', 'export', false, 'read')
]

export const operationsClassifierTable = [
  newOperation('Создать запись', 'create', false, 'write'),
  newOperation('Изменить запись', 'edit', true, 'read'),
  newOperation('Создать аттрибуты', 'create-attr', true, 'update'),
  newOperation('Изменить аттрибуты', 'edit-attr', true, 'update'),
  newOperation('Удалить запись', 'delete', true, 'delete'),
  newOperation('Экспорт', 'export', false, 'read')
]

export const operationsRoleTable = [
  newOperation('Создать роль', 'create', false, 'write'),
  newOperation('Изменить роль', 'edit', true, 'update'),
  newOperation('Заблокировать роль', 'block', true, 'delete')
]

export const operationsUserTable = [
  newOperation('Создать пользователя', 'create', false, 'write'),
  newOperation('Изменить пользователя', 'edit', true, 'read'),
  newOperation('Заблокировать пользователя', 'block', true, 'delete')
]

export const constructionOperations = [
  newOperation('Создать структуру', 'create', false, 'write'),
  newOperation('Изменить структуру', 'edit', true, 'read'),
  newOperation('Удалить структуру', 'delete', true, 'delete')
]

export const operationsIntegrationTables = [
  newOperation('Создать запись', 'create', false, 'write'),
  newOperation('Изменить запись', 'edit', true, 'read'),
  newOperation('Удалить запись', 'delete', true, 'delete')
]
