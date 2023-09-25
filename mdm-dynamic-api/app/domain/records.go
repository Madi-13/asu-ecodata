package domain

type GetRecordsInput struct {
	TableName   string      `json:"table_name"`
	TableID     *int        `json:"table_id"`
	RowID       *int        `json:"row_id"`
	ParentID    *int        `json:"parent_id"`
	Search      interface{} `json:"search"`
	PopupSearch string      `json:"search_popup"`
	Offset      int         `json:"offset"`
	Limit       int         `json:"limit"`
}

type RSearch struct {
	SearchBy    string      `json:"search_by"`
	SearchValue interface{} `json:"search_value"`
}

type GetCUDRecordsInput struct {
	TableName     string                 `json:"table_name"`
	OperationType string                 `json:"operation_type"`
	RowId         *int                   `json:"row_id,omitempty"`
	RColumns      map[string]interface{} `json:"columns"`
}

type LogHistoryEntity struct {
	DictionaryID *int    `json:"dictionary_id"`
	RowID        *int    `json:"row_id"`
	DateFrom     *string `json:"date_from"`
	DateTo       *string `json:"date_to"`
}

type ChangeStatus struct {
	DictionaryID string `json:"dictionary_id"`
	RowID        int    `json:"row_id"`
	NewStatus    string `json:"new_status"`
	ChangeReason string `json:"change_reason"`
}

type TablePopup struct {
	RowID    int    `json:"id"`
	RowValue string `json:"value"`
}

type GetTablePopup struct {
	Total      int          `json:"total"`
	TablePopup []TablePopup `json:"records"`
}

type GetRecordsOut struct {
	DbCode      int     `json:"db_code"`
	DbErrorDesc *string `json:"db_error_desc"`
}
