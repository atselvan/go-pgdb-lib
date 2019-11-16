package pgdb

import (
	"errors"
	"fmt"
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
	err := t.ValidateTableDefinition()
	if err != nil {
		return err
	}
	return dbConn.Exec(t.GetQuery())
}

// ValidateTableDefinition validates the table definition before creation of the table in the database
func (t *Table) ValidateTableDefinition() error {
	if t.Name == "" {
		return errors.New("table name cannot be empty")
	}
	if len(t.Columns) < 1 {
		return errors.New("the table should have at least one column")
	}
	for _, c := range t.Columns {
		if c.Name == "" || c.DataType == "" {
			return errors.New("column name and datatype cannot be empty")
		}
	}
	return nil
}

// Exists checks if a table already exists
// The method returns a boolean value or an error if something goes wrong
func (t *Table) Exists() (bool, error) {
	err := t.ValidateTableDefinition()
	if err != nil {
		return false, err
	}
	var (
		dbConn DbConn
		values []string
	)
	db, err := dbConn.Connect()
	if err != nil {
		return false, err
	}
	rows, err := db.Query(fmt.Sprintf("select exists (  select 1 from   information_schema.tables   where table_name = '%s');", t.Name))
	if err != nil {
		return false, err
	}
	for rows.Next() {
		var value string
		err = rows.Scan(&value)
		if err != nil {
			return false, err
		}
		values = append(values, value)
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
	if t.Name == "" {
		return errors.New("table name cannot be empty")
	}
	return dbConn.Exec(fmt.Sprintf("drop table %s", t.Name))
}
