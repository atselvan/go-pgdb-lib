package pgdb

import (
	"fmt"
	"github.com/atselvan/go-utils"
	"strings"
)

type Table struct {
	Name    string
	Columns []TableColumn
}

type TableColumn struct {
	Name        string
	DataType    string
	Constraints []string
}

// GetTableName returns the name of the table
func (t *Table) getTableName() string {
	return strings.TrimSpace(t.Name)
}

// isValidTableName checks if table name is not empty
// The method return an error if table name is a empty value
func (t *Table) isValidTableName() error {
	if t.getTableName() == "" {
		return utils.Error{ErrStr: tableNameEmptyErr, ErrMsg: tableNameEmptyErrStr}.NewError()
	} else {
		return nil
	}
}

// isValidColumnLength checks if at least one column definition is provided for the table
// The method returns an error if no column definition is specified
func (t *Table) isValidColumnLength() error {
	if len(t.Columns) < 1 {
		return utils.Error{ErrStr: tableColumnErr, ErrMsg: tableNoColumnErrStr}.NewError()
	} else {
		return nil
	}
}

// isValidColumnInfo checks if column information is correctly provided
// The method returns an error if empty values are provided
func (t *Table) isValidColumnInfo() error {
	for _, c := range t.Columns {
		if c.Name == "" || c.DataType == "" {
			return utils.Error{ErrStr: tableColumnErr, ErrMsg: tableColumnInfoErrStr}.NewError()
		}
	}
	return nil
}

// ValidateTableDefinition validates the table definition before creation of the table in the database
// The method returns an error if something goes wrong
func (t *Table) isValidTableInfo() error {
	if err := t.isValidTableName(); err != nil {
		return err
	} else if err := t.isValidColumnLength(); err != nil {
		return err
	} else if err := t.isValidColumnInfo(); err != nil {
		return err
	} else {
		return nil
	}
}

// GetQuery creates and returns a SQL table query which would be used to create a table in the database
func (t *Table) GetQuery() string {
	var query string
	query = fmt.Sprintf("CREATE TABLE %s (\n", t.Name)
	for i, c := range t.Columns {
		query += "\t" + c.Name + "\t" + c.DataType + "\t"
		for i, cs := range c.Constraints {
			query += cs
			if i < len(c.Constraints)-1 {
				query += "\t"
			}
		}
		if i < len(t.Columns)-1 {
			query += ",\n"
		} else {
			query += "\n);"
		}
	}
	return query
}

// Create creates a new table in the database
// The method returns an error if something goes wrong
func (t *Table) Create() error {
	var dbConn DbConn
	if err := t.isValidTableInfo(); err != nil {
		return err
	}
	return dbConn.Exec(t.GetQuery())
}

// Exists checks if a table already exists
// The method returns a boolean value or an error if something goes wrong
func (t *Table) Exists() (bool, error) {
	if err := t.isValidTableInfo(); err != nil {
		return false, err
	}
	var (
		dbConn DbConn
		values []string
	)
	db, err := dbConn.Connect()
	if err != nil {
		return false, dbConn.ConnectionError(err)
	}
	rows, err := db.Query(fmt.Sprintf("select exists (  select 1 from   information_schema.tables   where table_name = '%s');", t.Name))
	if err != nil {
		return false, dbConn.QueryExecError(err)
	}
	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			return false, dbConn.RowScanError(err)
		}
		values = append(values, value)
	}
	if err = dbConn.Close(db); err != nil {
		return false, dbConn.ClosureError(err)
	}
	if values[0] == "false" {
		return false, err
	} else {
		return true, err
	}
}

// Delete drops the table from tha database
// The method returns an error if something goes wrong
func (t *Table) Delete() error {
	var dbConn DbConn
	if err := t.isValidTableName(); err != nil {
		return err
	}
	return dbConn.Exec(fmt.Sprintf("drop table %s", t.Name))
}
