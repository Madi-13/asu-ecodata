export const systemRefDeserializeReq = (
  data = {},
  operationType = 'create'
) => ({
  table_name: data.table_name,
  table_code: data.table_code,
  type_prefix: data.type_prefix,
  [operationType == 'edit' && 'id']: data.id,
  table_desc: data.table_desc,
  table_type: data.table_type,
  table_prefix: data.table_prefix,
  columns: data.columns
    .map((col) => systemRefDeserializeColumn(col))
    .filter((col) => {
      if (operationType == 'edit' && col.operation_type) {
        return col
      } else if (operationType == 'create') {
        return col
      }
    })
})

const systemRefDeserializeColumn = (data = {}) => ({
  id: data.id,
  operation_type: data.operation_type,
  column_name: data.column_name,
  column_desc: data.column_desc,
  data_type: data.data_type,
  max_symbol_length: !data.max_symbol_length
    ? data.max_symbol_length
    : +data.max_symbol_length,
  is_unique: data.is_unique,
  foreign_table: data.foreign_table,
  is_required: data.is_required,
  is_filterable: data.is_filterable,
  is_default_filterable: data.is_default_filterable
})

export const systemRefSerializer = (data = {}) => ({
  id: data.id,
  table_name: data.table_name,
  table_code: data.table_code,
  table_desc: data.table_desc,
  type_prefix: data.type_prefix,
  table_type: { id: +data.table_type.id, title: data.table_type.title },
  columns: data.columns
    .filter((col) => col.name != 'id')
    .map((col) => systemRefSerializeColumn(col))
})

const systemRefSerializeColumn = (column = {}) => ({
  operation_type: column.operation_type,
  id: column.Id,
  column_name: column.name_ru,
  column_desc: column.column_desc,
  data_type: column.data_type,
  max_symbol_length: !column.data_length
    ? column.data_length
    : +column.data_length,
  foreign_table: column.foreign_table,
  column_system_name: column.name,
  is_unique: column.is_unique,
  is_required: column.is_required,
  is_filterable: column.is_filterable,
  is_default_filterable: column.is_default_filterable
})
