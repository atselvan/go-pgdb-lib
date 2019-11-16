package pgdb

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

// DBConn database connection details
type DbConn struct {
	Hostname string
	Port     string
	Name     string
	Username string
	Password string
}

// getConn gets the database connection details from environment variables
// If the documented environment variables are not set the method return default values
func (dbConn *DbConn) GetConn() DbConn {
	// set db hostname
	if dbConn.Hostname = os.Getenv("DB_HOSTNAME"); dbConn.Hostname == "" {
		dbConn.Hostname = "192.168.2.75"
	}
	// set db port
	if dbConn.Port = os.Getenv("DB_PORT"); dbConn.Port == "" {
		dbConn.Port = "5432"
	}
	// set db name
	if dbConn.Name = os.Getenv("DB_NAME"); dbConn.Name == "" {
		dbConn.Name = "assets"
	}
	// set db username
	if dbConn.Username = os.Getenv("DB_USERNAME"); dbConn.Username == "" {
		dbConn.Username = "postgres"
	}
	// set db password
	if dbConn.Password = os.Getenv("DB_PASSWORD"); dbConn.Password == "" {
		dbConn.Password = "postgres"
	}
	return *dbConn
}

// GetDBConnString gets the database connection details and returns the database connection string
func (dbConn *DbConn) GetConnString() string {
	*dbConn = dbConn.GetConn()
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbConn.Hostname, dbConn.Port, dbConn.Name, dbConn.Username, dbConn.Password)
}

// Connect establishes a connection with the database and validates the connection
// The method return the db connection or an error if something goes wrong
func (dbConn *DbConn) Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", dbConn.GetConnString())
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	return db, nil
}

// Close closes the connection with the database
// The method return an error if something goes wrong
func (dbConn *DbConn) Close(db *sql.DB) error {
	return db.Close()
}

// Exec executes a query on the database
// The method return an error if something goes wrong
func (dbConn *DbConn) Exec(query string) error {
	db, err := dbConn.Connect()
	if err != nil {
		return err
	}
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	err = dbConn.Close(db)
	if err != nil {
		return err
	}
	return nil
}
