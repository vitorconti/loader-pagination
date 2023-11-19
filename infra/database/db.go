package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

var DB *sql.DB

func Init() {
	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := "3306"
	dbName := "corgi"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
func Ping() {

	err := DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
}
func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("db", DB)
		return next(c)
	}
}
