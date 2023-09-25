export const serializePermissions = (resources = []) => {
  let serializedPermisssions = {}

  for (let resource of resources) {
    serializedPermisssions[resource.resource_code] = {}

    if (!resource.permissions) {
      continue
    }

    for (const perm of resource.permissions) {
      for (const [key, value] of Object.entries(perm)) {
        if (
          serializedPermisssions[resource.resource_code][key] &&
          value.role_name != 'head_of_nsi'
        ) {
          continue
        } else {
          serializedPermisssions[resource.resource_code][key] = value
        }
      }
    }
  }
  return serializedPermisssions
}
