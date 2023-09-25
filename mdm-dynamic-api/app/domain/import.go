package domain

type RImport struct {
	TableName      string `json:"table_name,omitempty"`
	RFile          string `json:"file_body"`
	FileName       string `json:"file_name"`
	ImportTypeID   string `json:"import_type_id,omitempty"`
	ImportTypeCode string `json:"import_type_code,omitempty"`
}

type ImportColumnStruct struct {
	Form []struct {
		IsRequired   bool   `json:"is_required"`
		MaxLength    any    `json:"max_length"`
		Name         string `json:"name"`
		OrderNum     int    `json:"order_num"`
		RefTableID   any    `json:"ref_table_id"`
		RefTableName any    `json:"ref_table_name"`
		RefTableType any    `json:"ref_table_type"`
		Title        string `json:"title"`
		Type         string `json:"type"`
	} `json:"form"`
}

type IAttrubuteStructure struct {
	OperationType   string   `json:"operation_type"`
	AttributeID     *int     `json:"attribute_id"`
	Name            string   `json:"name"`
	AttributeDictID int      `json:"attribute_dict_id"`
	DictionaryRowID int      `json:"dictionary_row_id"`
	Values          []IValue `json:"values"`
}

type IValue struct {
	Key    *int   `json:"key"`
	Value  string `json:"value"`
	Status string `json:"status"`
}
type IFormat struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	IsRequired bool   `json:"is_required"`
}
