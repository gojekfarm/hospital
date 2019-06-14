package storage

import (
	"database/sql"
	"fmt"
	"os"

	// postgres
	_ "github.com/lib/pq"
)

var db *sql.DB

//Initialize initializes the database
func Initialize() {
	databaseHost := os.Getenv("DB_HOST")
	databasePort := os.Getenv("DB_PORT")
	databaseUser := os.Getenv("DB_USER")
	databasePass := os.Getenv("DB_PASS")
	databaseName := os.Getenv("DB_NAME")

	postgresConnectionURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", databaseUser, databasePass, databaseHost, databasePort, databaseName)

	var err error
	db, err = sql.Open("postgres", postgresConnectionURL)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected!")

}

// Ping check connection to DB
func Ping() {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
