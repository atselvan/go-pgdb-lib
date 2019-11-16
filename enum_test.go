package pgdb

import "testing"

func TestEnum_GetEnumName(t *testing.T) {
	e := Enum{
		Name:   " test ",
	}
	expected := "test"
	if e.GetEnumName() != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, e.GetEnumName())
	}
}

func TestEnum_isValidEnumName(t *testing.T) {
	e := Enum{
		Name:   "",
	}
	if err := e.isValidEnumName(); err == nil {
		t.Errorf("Expected a error as enum name is empty")
	}

	e = Enum{
		Name:   " ",
	}
	if err := e.isValidEnumName(); err == nil {
		t.Errorf("Expected a error as enum name only has a blank space")
	}

	e = Enum{
		Name:   "test",
	}
	if err := e.isValidEnumName(); err != nil {
		t.Errorf("Did not expect an error as enum name is valid")
	}
}
