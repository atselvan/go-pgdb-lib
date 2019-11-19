package pgdb

import "testing"

func TestEnum_IsValidEnumName(t *testing.T) {
	e := Enum{
		Name: "",
	}
	if err := e.IsValidEnumName(); err == nil {
		t.Errorf("Expected a error as enum name is empty")
	}

	e = Enum{
		Name: " ",
	}
	if err := e.IsValidEnumName(); err == nil {
		t.Errorf("Expected a error as enum name only has a blank space")
	}

	e = Enum{
		Name: "test",
	}
	if err := e.IsValidEnumName(); err != nil {
		t.Errorf("Did not expect an error as enum name is valid")
	}
}
