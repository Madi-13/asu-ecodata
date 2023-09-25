import { withAsync } from "@/api/helpers/withAsync";
import { getAttrValues } from "@/api/attsAPI";

const actions = {
  fetchAttrValues({ commit }, { payload, getCurrentRowId, referenceId }) {
    return Promise.all(payload.map(controller => withAsync(getAttrValues, {
      dictionary_id: referenceId,
      dictionary_row_id: getCurrentRowId,
      attribute_dictionary_id: controller.ref_table_id,
      attribute_dictionary_row_id: controller[controller.name]?.id
    })))
  },
}

export default {
  actions
}
