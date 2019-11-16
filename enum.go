package pgdb

import (
	"errors"
	"fmt"
)

type Enum struct {
	name   string
	values []string
}

// Exists checks if a enum type already exists or not in the database
// The method returns a boolean value and an error depending on the result
func (e *Enum) Exists() (bool, error) {
	var dbConn DbConn
	if e.name == "" {
		return false, errors.New("enum name cannot be empty")
	}
	err := dbConn.Exec(fmt.Sprintf("select unnest (enum_range(NULL::%s));", e.name))
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
	rows, err := db.Query(fmt.Sprintf("select unnest (enum_range(NULL::%s));", e.name))
	if err != nil {
		return err
	}
	for rows.Next() {
		var value string
		err = rows.Scan(&value)
		if err != nil {
			return err
		}
		e.values = append(e.values, value)
	}
	return err
}

// Create creates a new enum type in the database
// The method returns an error if something goes wrong
func (e *Enum) Create() error {
	var dbConn DbConn
	if e.name == "" {
		return errors.New("enum name cannot be empty")
	}
	return dbConn.Exec(fmt.Sprintf("create type %s as enum ();", e.name))
}

// Add updates the enum type in the database
// The method returns an error if something goes wrong
func (e *Enum) Update() error {
	var dbConn DbConn
	if e.name == "" || len(e.values) < 1 {
		return errors.New("enum name or value cannot be empty")
	}
	return dbConn.Exec(fmt.Sprintf("alter type %s add value '%s';", e.name, e.values[0]))
}

// Delete removes the enum type from the database
// The method returns an error if something goes wrong
func (e *Enum) Delete() error {
	var dbConn DbConn
	if e.name == "" {
		return errors.New("enum name cannot be empty")
	}
	return dbConn.Exec(fmt.Sprintf("drop type %s;", e.name))
}
