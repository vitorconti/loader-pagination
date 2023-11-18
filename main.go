package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vitorconti/loader-pagination/seed"
)

func main() {
	// Connection parameters
	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := "3306"
	dbName := "corgi"

	// Create a DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	// Open a database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Example: Querying the database
	seed.Seed(db)
}
