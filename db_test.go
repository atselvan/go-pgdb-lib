package pgdb

import (
	"os"
	"testing"
)

func TestDbConn_GetConn_Default(t *testing.T) {
	var dbConn DbConn
	dbConn = dbConn.GetConn()

	// default values
	defaultHostname := "192.168.2.75"
	defaultPort := "5432"
	defaultName := "assets"
	defaultUsername := "postgres"
	defaultPassword := "postgres"

	// check default hostname
	if dbConn.Hostname != defaultHostname {
		t.Errorf("Database hostname was incorrect, got: %s, want: %s", dbConn.Hostname, defaultHostname)
	}
	// check default port
	if dbConn.Port != defaultPort {
		t.Errorf("Database port was incorrect, got: %s, want: %s", dbConn.Port, defaultPort)
	}
	// check default name
	if dbConn.Name != defaultName {
		t.Errorf("Database name was incorrect, got: %s, want: %s", dbConn.Name, defaultName)
	}
	// check default username
	if dbConn.Username != defaultUsername {
		t.Errorf("Database username was incorrect, got: %s, want: %s", dbConn.Username, defaultUsername)
	}
	// check default password
	if dbConn.Password != defaultPassword {
		t.Errorf("Database password was incorrect, got: %s, want: %s", dbConn.Password, defaultPassword)
	}
}

func TestDbConn_GetConn_Env(t *testing.T) {
	var dbConn DbConn

	// set environment values
	// default values
	envHostname := "example.com"
	_ = os.Setenv("DB_HOSTNAME", envHostname)
	envPort := "1234"
	_ = os.Setenv("DB_PORT", envPort)
	envName := "test"
	_ = os.Setenv("DB_NAME", envName)
	envUsername := "admin"
	_ = os.Setenv("DB_USERNAME", envUsername)
	envPassword := "admin123"
	_ = os.Setenv("DB_PASSWORD", envPassword)

	dbConn = dbConn.GetConn()

	// check env hostname
	if dbConn.Hostname != envHostname {
		t.Errorf("Database hostname was incorrect, got: %s, want: %s", dbConn.Hostname, envHostname)
	}
	// check env port
	if dbConn.Port != envPort {
		t.Errorf("Database port was incorrect, got: %s, want: %s", dbConn.Port, envPort)
	}
	// check env name
	if dbConn.Name != envName {
		t.Errorf("Database name was incorrect, got: %s, want: %s", dbConn.Name, envName)
	}
	// check env username
	if dbConn.Username != envUsername {
		t.Errorf("Database username was incorrect, got: %s, want: %s", dbConn.Username, envUsername)
	}
	// check env password
	if dbConn.Password != envPassword {
		t.Errorf("Database password was incorrect, got: %s, want: %s", dbConn.Password, envPassword)
	}
	_ = os.Setenv("DB_HOSTNAME", "example.com")

	dbConn = dbConn.GetConn()
	if dbConn.Hostname != "example.com" {
		t.Errorf("Hostname was incorrect, got: %s, want: %s", dbConn.Hostname, "example.com")
	}
}
