package pgdb

import (
	"errors"
	"fmt"
)

type Enum struct {
	Name   string
	Values []string
}

// Exists checks if a enum type already exists or not in the database
// The method returns a boolean value and an error depending on the result
func (e *Enum) Exists() (bool, error) {
	var dbConn DbConn
	if e.Name == "" {
		return false, errors.New("enum name cannot be empty")
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
	db, err := dbConn.Connect()
	if err != nil {
		return err
	}
	rows, err := db.Query(fmt.Sprintf("select unnest (enum_range(NULL::%s));", e.Name))
	if err != nil {
		return err
	}
	for rows.Next() {
		var value string
		err = rows.Scan(&value)
		if err != nil {
			return err
		}
		e.Values = append(e.Values, value)
	}
	return err
}

// Create creates a new enum type in the database
// The method returns an error if something goes wrong
func (e *Enum) Create() error {
	var dbConn DbConn
	if e.Name == "" {
		return errors.New("enum name cannot be empty")
	}
	return dbConn.Exec(fmt.Sprintf("create type %s as enum ();", e.Name))
}

// Add updates the enum type in the database
// The method returns an error if something goes wrong
func (e *Enum) Update() error {
	var dbConn DbConn
	if e.Name == "" || len(e.Values) < 1 {
		return errors.New("enum name or value cannot be empty")
	}
	return dbConn.Exec(fmt.Sprintf("alter type %s add value '%s';", e.Name, e.Values[0]))
}

// Delete removes the enum type from the database
// The method returns an error if something goes wrong
func (e *Enum) Delete() error {
	var dbConn DbConn
	if e.Name == "" {
		return errors.New("enum name cannot be empty")
	}
	return dbConn.Exec(fmt.Sprintf("drop type %s;", e.Name))
}
