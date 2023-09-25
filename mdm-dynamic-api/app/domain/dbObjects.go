package domain

import (
	"encoding/json"
	"time"
)

type Table struct {
	Id          int       `json:"id"`
	Name        *string   `json:"table_name"`
	Code        *string   `json:"table_code"`
	Description *string   `json:"table_descr"`
	Status      *string   `json:"status"`
	TableType   TableType `json:"table_type"`
	TypePrefix  *string   `json:"type_prefix"`
	Columns     []Column  `json:"columns"`
}

type InfoPanel struct {
	DictionaryID *int `json:"dictionary_id"`
	FilterID     *int `json:"filter_id"`
}
type FilterListOut struct {
	ID         int    `json:"id"`
	FilterName string `json:"filter_name"`
}

type TableType struct {
	TypeId   *string `json:"id"`
	TypeName *string `json:"title"`
}

type RTableOut struct {
	TableId   *string `json:"table_id"`
	TableName *string `json:"table_name"`
	TableShow *string `json:"table_show"`
	IsTree    *bool   `json:"is_tree"`
}

type OutTable struct {
	Total  int     `json:"total"`
	Result []Table `json:"result"`
}

type Column struct {
	Id                  int           `json:"id,omitempty`
	DictionariesId      *int          `json:"dictionaries_id,omitempty"`
	Name                *string       `json:"name"`
	NameRu              *string       `json:"name_ru"`
	Description         *string       `json:"column_desc"`
	DataType            ColumnType    `json:"data_type"`
	DataLength          *int          `json:"data_length"`
	ForeignTable        ForeignTables `json:"foreign_table"`
	IsRequired          *bool         `json:"is_required"`
	IsUnique            *bool         `json:"is_unique"`
	IsFilterable        *bool         `json:"is_filterable"`
	IsDefaultFilterable *bool         `json:"is_default_filterable"`
}

type ColumnType struct {
	Id               int     `json:"id"`
	Name             *string `json:"column_name"`
	NameDesc         *string `json:"column_desc"`
	NameRu           *string `json:"name_ru"`
	SqlType          *string `json:"sql_type,omitempty"`
	NeedLength       *bool   `json:"is_need_length"`
	NeedForeignTable *bool   `json:"is_need_foreign_table"`
	NeedParentId     *bool   `json:"need_parent_id,omitempty"`
}

type ForeignTables struct {
	TableName *string `json:"table_name"`
	TableCode *string `json:"table_code"`
}

type CreateTable struct {
	TableName   string            `json:"table_name"`
	TablePrefix string            `json:"table_prefix"`
	TableDesc   string            `json:"table_desc"`
	Columns     []json.RawMessage `json:"columns"`
}

type AlterTable struct {
	TableName string            `json:"table_name"`
	TableDesc string            `json:"table_desc"`
	TableType int               `json:"table_type"`
	TableID   int               `json:"table_id"`
	Columns   []json.RawMessage `json:"columns"`
}

type GetTables struct {
	TableID     *int   `json:"table_id"`
	TableTypeID *int   `json:"table_type_id"`
	Search      string `json:"search"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
}

type CreateTableOut struct {
	DbCode      int     `json:"db_code"`
	DbErrorDesc *string `json:"db_error_desc"`
}

type NCreateTable struct {
	TableName  string `json:"table_name"`
	TableID    *int   `json:"table_id"`
	ID         *int   `json:"id"`
	TableDescr string `json:"table_descr"`
	TableCode  string `json:"table_code"`
	TableType  struct {
		ID    int     `json:"id"`
		Title *string `json:"title"`
	} `json:"table_type"`
	TablePrefix string `json:"table_prefix"`
	Columns     []struct {
		ID            int     `json:"id"`
		OperationType *string `json:"operation_type"`
		ColumnName    *string `json:"column_name"`
		ColumnDesc    *string `json:"column_desc"`
		DataType      struct {
			ID                 int     `json:"id"`
			NameRu             *string `json:"name_ru"`
			Name               *string `json:"name"`
			IsNeedLength       *bool   `json:"is_need_length"`
			IsNeedForeignTable *bool   `json:"is_need_foreign_table"`
		} `json:"data_type"`
		ForeignTable struct {
			TableName *string `json:"table_name"`
			TableCode *string `json:"table_code"`
		} `json:"foreign_table"`
		MaxSymbolLength     *int  `json:"max_symbol_length"`
		IsRequired          *bool `json:"is_required"`
		IsUnique            *bool `json:"is_unique"`
		IsFilterable        *bool `json:"is_filterable"`
		IsDefaultFilterable *bool `json:"is_default_filterable"`
		IsNullable          *bool `json:"is_nullable"`
	} `json:"columns"`
}

type OutFiles struct {
	Result []struct {
		ID          int    `json:"id"`
		UUID        string `json:"uuid"`
		DateCreated string `json:"date_created"`
		UserCreated string `json:"user_created"`
	} `json:"result"`
}

type RFiles struct {
	ID          int       `json:"file_id"`
	UUID        string    `json:"file_name"`
	FileBody    string    `json:"file_body"`
	DateCreated time.Time `json:"date_created,omitempty"`
	UserCreated string    `json:"user_created,omitempty"`
}
