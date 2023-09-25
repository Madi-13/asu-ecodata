export const serializeAttr = (attrs = []) => {
  return attrs.map((attr) => ({
    id: attr.attrib_id,
    name: attr.name,
    options: attr.attrib_list,
    selected: { key: attr.attrib_key, value: attr.attrib_value }
  }))
}
