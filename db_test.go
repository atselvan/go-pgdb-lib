package pgdb

import (
	"fmt"
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
	if dbConn.hostname != defaultHostname {
		t.Errorf("Database hostname was incorrect, got: %s, want: %s", dbConn.hostname, defaultHostname)
	}
	// check default port
	if dbConn.port != defaultPort {
		t.Errorf("Database port was incorrect, got: %s, want: %s", dbConn.port, defaultPort)
	}
	// check default name
	if dbConn.name != defaultName {
		t.Errorf("Database name was incorrect, got: %s, want: %s", dbConn.name, defaultName)
	}
	// check default username
	if dbConn.username != defaultUsername {
		t.Errorf("Database username was incorrect, got: %s, want: %s", dbConn.username, defaultUsername)
	}
	// check default password
	if dbConn.password != defaultPassword {
		t.Errorf("Database password was incorrect, got: %s, want: %s", dbConn.password, defaultPassword)
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
	if dbConn.hostname != envHostname {
		t.Errorf("Database hostname was incorrect, got: %s, want: %s", dbConn.hostname, envHostname)
	}
	// check env port
	if dbConn.port != envPort {
		t.Errorf("Database port was incorrect, got: %s, want: %s", dbConn.port, envPort)
	}
	// check env name
	if dbConn.name != envName {
		t.Errorf("Database name was incorrect, got: %s, want: %s", dbConn.name, envName)
	}
	// check env username
	if dbConn.username != envUsername {
		t.Errorf("Database username was incorrect, got: %s, want: %s", dbConn.username, envUsername)
	}
	// check env password
	if dbConn.password != envPassword {
		t.Errorf("Database password was incorrect, got: %s, want: %s", dbConn.password, envPassword)
	}
	_ = os.Setenv("DB_HOSTNAME", "example.com")

	fmt.Println(os.Getenv("DB_HOSTNAME"))

	dbConn = dbConn.GetConn()
	if dbConn.hostname != "example.com" {
		t.Errorf("Hostname was incorrect, got: %s, want: %s", dbConn.hostname, "example.com")
	}
}
