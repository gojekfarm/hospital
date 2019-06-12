package storage

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

// Migration will migrate all the way up
func Migration() {
	// databaseHost := os.Getenv("DOCTOR_DB_HOST")
	// databasePort := os.Getenv("DOCTOR_DB_PORT")
	// databaseUser := os.Getenv("DOCTOR_DB_USER")
	// databasePass := os.Getenv("DOCTOR_DB_PASS")
	// databaseName := os.Getenv("DOCTOR_DB_NAME")

	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	databaseHost, databasePort, databaseUser, databasePass, databaseName)
	// //var err error
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
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
	err = m.Up()
	if err != nil {
		panic(err)
	}
}

//RollBack will migrate one version down
func RollBack() {
	// databaseHost := os.Getenv("DOCTOR_DB_HOST")
	// databasePort := os.Getenv("DOCTOR_DB_PORT")
	// databaseUser := os.Getenv("DOCTOR_DB_USER")
	// databasePass := os.Getenv("DOCTOR_DB_PASS")
	// databaseName := os.Getenv("DOCTOR_DB_NAME")

	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	databaseHost, databasePort, databaseUser, databasePass, databaseName)
	// //var err error
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
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
