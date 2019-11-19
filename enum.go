package pgdb

import (
	"fmt"
	"github.com/atselvan/go-utils"
	"strings"
)

// Enum represents a Enum type in the database
type Enum struct {
	Name   string
	Values []string
}

// IsValidEnumName checks if enum name is not empty
// The method return an error if enum name is a empty value
func (e *Enum) IsValidEnumName() error {
	e.Name = strings.TrimSpace(e.Name)
	if e.Name == "" {
		return utils.Error{ErrStr: enumNameEmptyErr, ErrMsg: enumNameEmptyErrStr}.NewError()
	} else {
		return nil
	}
}

// Exists checks if a enum type already exists or not in the database
// The method returns a boolean value and an error depending on the result
func (e *Enum) Exists() (bool, error) {
	var dbConn DbConn
	if err := e.IsValidEnumName(); err != nil {
		return false, err
	}
	err := dbConn.Exec(fmt.Sprintf("select unnest (enum_range(NULL::%s));", e.Name))
	if err == nil {
		return true, err
	} else {
		return false, err
	}
}

// Get returns a list of enum type values
// The method returns the values of a enum or an error
func (e *Enum) Get() error {
	var dbConn DbConn
	if err := e.IsValidEnumName(); err != nil {
		return err
	}
	db, err := dbConn.Connect()
	if err != nil {
		return err
	}
	rows, err := db.Query(fmt.Sprintf("select unnest (enum_range(NULL::%s));", e.Name))
	if err != nil {
		return dbConn.QueryExecError(err)
	}
	for rows.Next() {
		var value string
		err = rows.Scan(&value)
		if err != nil {
			return dbConn.RowScanError(err)
		}
		e.Values = append(e.Values, value)
	}
	return err
}

// Create creates a new enum type in the database
// The method returns an error if something goes wrong
func (e *Enum) Create() error {
	var dbConn DbConn
	if err := e.IsValidEnumName(); err != nil {
		return err
	}
	return dbConn.Exec(fmt.Sprintf("create type %s as enum ();", e.Name))
}

// Add updates the enum type in the database
// The method returns an error if something goes wrong
func (e *Enum) Update() error {
	var dbConn DbConn
	if err := e.IsValidEnumName(); err != nil {
		return err
	}
	return dbConn.Exec(fmt.Sprintf("alter type %s add value '%s';", e.Name, e.Values[0]))
}

// Delete removes the enum type from the database
// The method returns an error if something goes wrong
func (e *Enum) Delete() error {
	var dbConn DbConn
	if err := e.IsValidEnumName(); err != nil {
		return err
	}
	return dbConn.Exec(fmt.Sprintf("drop type %s;", e.Name))
}
