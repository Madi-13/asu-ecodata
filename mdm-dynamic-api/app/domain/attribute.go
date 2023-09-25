package domain

type GetAttributeEntiry struct {
	DictionaryID    int `json:"dictionary_id"`
	DictionaryRowID int `json:"dictionary_row_id"`
}

type AttributeFastCreate struct {
	AttributeID    int    `json:"attrib_id"`
	AttributeValue string `json:"value"`
}

type CUDAttribute struct {
	OperationType   string `json:"operation_type"`
	AttributeID     any    `json:"attribute_id"`
	Name            string `json:"name"`
	AttributeDictID int    `json:"attribute_dict_id"`
	DictionaryRowID int    `json:"dictionary_row_id"`
	Values          []struct {
		Key    *int    `json:"key"`
		Value  *string `json:"value"`
		Status string  `json:"status"`
	} `json:"values"`
}

type CUDAttributeValues struct {
	DictionaryID    int         `json:"dictionary_id"`
	AttributeDictID int         `json:"attribute_dict_id"`
	DictionaryRowID int         `json:"dictionary_row_id"`
	Attributes      []Attribute `json:"attributes"`
}

type Attribute struct {
	AttributeID int `json:"attribute_id"`
	Key         int `json:"key"`
}

type GetAttributeValues struct {
	DictionaryID             int  `json:"dictionary_id"`
	DictionaryRowID          *int `json:"dictionary_row_id"`
	AttributeDictionaryID    int  `json:"attribute_dictionary_id"`
	AttributeDictionaryRowID int  `json:"attribute_dictionary_row_id"`
}
