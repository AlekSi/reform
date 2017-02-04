package main

import (
	"database/sql"
	"fmt"
)

//go:generate reform

type yesNo bool

func (yn *yesNo) Scan(src interface{}) error {
	var str string
	switch s := src.(type) {
	case string:
		str = s
	case []byte:
		str = string(s)
	default:
		return fmt.Errorf("unexpected type %T (%#v)", src, src)
	}

	switch str {
	case "YES":
		*yn = true
	case "NO":
		*yn = false
	default:
		return fmt.Errorf("unexpected %q", str)
	}
	return nil
}

// check interface
var _ sql.Scanner = (*yesNo)(nil)

//reform:information_schema.tables
type table struct {
	Catalog string `reform:"table_catalog"`
	Schema  string `reform:"table_schema"`
	Name    string `reform:"table_name"`
	Type    string `reform:"table_type"`
}

//reform:information_schema.columns
type column struct {
	TableCatalog string `reform:"table_catalog"`
	TableSchema  string `reform:"table_schema"`
	TableName    string `reform:"table_name"`
	Name         string `reform:"column_name"`
	IsNullable   yesNo  `reform:"is_nullable"`
	Type         string `reform:"data_type"`
}

//reform:sqlite_master
type sqliteMaster struct {
	Name string `reform:"name"`
}

// TODO This "dummy" table name is ugly. We should do better.
//reform:dummy
type sqliteTableInfo struct {
	CID          int     `reform:"cid"`
	Name         string  `reform:"name"`
	Type         string  `reform:"type"`
	NotNull      bool    `reform:"notnull"`
	DefaultValue *string `reform:"dflt_value"`
	PK           bool    `reform:"pk"`
}
