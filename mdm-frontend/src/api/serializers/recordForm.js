export const serializeRecordForms = (controllers = []) => {
  return controllers.map((ctr) => {
    ctr.is_required = true
    if (!ctr.value) {
      ctr[ctr.name] = null
    } else {
      ctr[ctr.name] =
        ctr.type == 'autocomplete' || ctr.type == 'tree'
          ? { id: ctr.value, value: ctr.value_show }
          : ctr.value
    }
    return ctr
  })
}

export const deserializeRecordForm = (data = {}) => {
  let columns = {}
  data.form.forEach((el) => {
    if (el.type == 'autocomplete' || el.type == 'tree') {
      columns[el.name || 'name'] = el[el.name] && el[el.name].id
    } else {
      columns[el.name || 'name'] = el[el.name]
    }
  })
  return {
    table_name: data.table_name,
    operation_type: data.operation_type,
    row_id: data.row_id,
    columns
  }
}
