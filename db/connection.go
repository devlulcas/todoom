package db

import (
	"database/sql"
	"fmt"

	"github.com/devlulcas/todoom/configs"

	// Importing the postgres driver (the underscore is to avoid unused import error)
	_ "github.com/lib/pq"
)

// Opens a connection with the database
func OpenConnection(conf configs.DBConfig) (*sql.DB, error) {
	stringConnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	// Opens the connection
	conn, err := sql.Open("postgres", stringConnection)
	if err != nil {
		return nil, err
	}

	// Checks if the connection is alive
	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	// Returns the connection
	return conn, nil
}

// Closes the connection with the database
func CloseConnection(conn *sql.DB) error {
	err := conn.Close()

	return err
}
