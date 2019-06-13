package storage

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

// Migration will migrate all the way up
func Migration() {
	// Initialize()
	// driver, err := postgres.WithInstance(db, &postgres.Config{})
	// if err != nil {
	// 	panic(err)
	// }
	databaseHost := os.Getenv("DB_HOST")
	databasePort := os.Getenv("DB_PORT")
	databaseUser := os.Getenv("DB_USER")
	databasePass := os.Getenv("DB_PASS")
	databaseName := os.Getenv("DB_NAME")

	postgresConnectionURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", databaseUser, databasePass, databaseHost, databasePort, databaseName)

	m, err := migrate.New(
		"file://./migrations", postgresConnectionURL)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil && err.Error() != "no change" {
		panic(err)
	}
}

//DownOneStep will migrate one version down
func DownOneStep() {
	Initialize()
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	err = m.Steps(-1)
	if err != nil {
		panic(err)
	}
}
