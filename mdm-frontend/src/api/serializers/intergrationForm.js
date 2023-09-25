export const processDeserialize = (data) => ({
  operation_type: data.operation_type,
  process: {
    dictionary_id: data.process.dictionary.id,
    process_name: data.process.process_name,
    process_order: data.process.process_order,
    property_id: data.process.property.id,
    scenario_id: data.process.scenario.id
  }
})

export const processBuildForm = (data) => ({
  id: data.id,
  process_name: data.process_name,
  dictionary: {
    id: data.dictionary_id,
    table_name: data.dictionary_local
  },
  process_order: null,
  property: {
    id: data.property_id,
    property_name: data.property_name
  },
  scenario: {
    id: data.scenario_id,
    scenario_name: data.scenario_name
  }
})
