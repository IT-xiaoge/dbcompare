package dbcompare

import (
	"database/sql"

	"github.com/wnote/worm"
)

type CompareConfig struct {
	Db1Dn    string
	Db1Table string
	Db2Dn    string
	Db2Table string
}

// Database diff
type DatabaseDiff struct {
	DbTables         [2][]string
	TablesDiffResult map[string]TableDiff
}

// Table diff
type TableDiff struct {
	Fields          [2][]string
	FieldDiffResult map[string]FieldDiff
}

// Field diff
type FieldDiff struct {
	Type    [2]string
	Null    [2]string
	Key     [2]string
	Default [2]string
	Extra   [2]string
}

type Field struct {
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

// Compare return the pointer of the struct DatabaseDiff{}
func Compare(config CompareConfig) (*DatabaseDiff, error) {
	return compareTable(config)
}

func compareTable(config CompareConfig) (*DatabaseDiff, error) {
	db1, err := worm.OpenDb("mysql", config.Db1Dn)
	if err != nil {
		return nil, err
	}
	db1TablesMap, err := getTablesMap(db1)
	if err != nil {
		return nil, err
	}
	db2, err := worm.OpenDb("mysql", config.Db2Dn)
	if err != nil {
		return nil, err
	}
	db2TablesMap, err := getTablesMap(db2)
	if err != nil {
		return nil, err
	}
	var databaseDiff DatabaseDiff
	var db1MoreTable []string
	var db2MoreTable []string
	tablesDiffResult := make(map[string]TableDiff)
	for table1 := range db1TablesMap {
		if _, exist := db2TablesMap[table1]; !exist {
			db1MoreTable = append(db1MoreTable, table1)
		}
	}
	for table2 := range db2TablesMap {
		if _, exist := db1TablesMap[table2]; !exist {
			db2MoreTable = append(db2MoreTable, table2)
		} else {
			fieldResult, err := compareField(db1, db2, table2)
			if err != nil {
				return nil, err
			}
			tablesDiffResult[table2] = *fieldResult
		}
	}
	databaseDiff.TablesDiffResult = tablesDiffResult
	databaseDiff.DbTables = [2][]string{db1MoreTable, db2MoreTable}
	return &databaseDiff, nil
}

func getTablesMap(db *sql.DB) (map[string]bool, error) {
	rows, err := db.Query("show tables;")
	if err != nil {
		return nil, err
	}
	dbTablesMap := make(map[string]bool)
	for rows.Next() {
		var tableName string
		rows.Scan(&tableName)
		dbTablesMap[tableName] = true
	}
	return dbTablesMap, nil
}

func compareField(db1 *sql.DB, db2 *sql.DB, tableName string) (*TableDiff, error) {
	db1FieldsMap, err := getFieldsMap(db1, tableName)
	if err != nil {
		return nil, err
	}
	db2FieldsMap, err := getFieldsMap(db2, tableName)
	if err != nil {
		return nil, err
	}
	var tableDiff TableDiff
	var db1MoreField []string
	var db2MoreField []string
	fieldsDiffResult := make(map[string]FieldDiff)
	for fieldName := range db1FieldsMap {
		if _, exist := db2FieldsMap[fieldName]; !exist {
			db1MoreField = append(db1MoreField, fieldName)
		}
	}
	for fieldName, field := range db2FieldsMap {
		if db1Field, exist := db1FieldsMap[fieldName]; !exist {
			db2MoreField = append(db2MoreField, fieldName)
		} else {
			fieldDiff := FieldDiff{
				Type:    [2]string{db1Field.Type, field.Type},
				Null:    [2]string{db1Field.Null, field.Null},
				Key:     [2]string{db1Field.Key, field.Key},
				Default: [2]string{db1Field.Default, field.Default},
				Extra:   [2]string{db1Field.Extra, field.Extra},
			}
			fieldsDiffResult[fieldName] = fieldDiff
		}
	}
	tableDiff.Fields = [2][]string{db1MoreField, db2MoreField}
	tableDiff.FieldDiffResult = fieldsDiffResult
	return &tableDiff, nil
}

func getFieldsMap(db *sql.DB, tableName string) (map[string]Field, error) {
	rows, err := db.Query("show fields from " + tableName + ";")
	if err != nil {
		return nil, err
	}
	fieldsMap := make(map[string]Field)
	for rows.Next() {
		var fieldName, fType, fNull, fKey, fDefault, fExtra string
		rows.Scan(&fieldName, &fType, &fNull, &fKey, &fDefault, &fExtra)
		fieldsMap[fieldName] = Field{
			Type:    fType,
			Null:    fNull,
			Key:     fKey,
			Default: fDefault,
			Extra:   fExtra,
		}
	}
	return fieldsMap, nil
}
