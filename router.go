package main

import (
	"database/sql"
	"fmt"
	"hospital/doctor"
	"hospital/reception"
	"log"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "Hospital - An autonomous healing System"
	app.Usage = "fix fault/failure in system"
	app.Author = "Dilip"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "runserver",
			Aliases: []string{"s"},
			Usage:   "Starts the server",
			Action: func(c *cli.Context) {
				port := "8088"
				if c.Args().Present() {
					port = c.Args()[0]
				}
				startserver(port)
			},
		},
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "Database migration",
			Action: func(c *cli.Context) {
				migration()
				fmt.Println("Migrating...")
			},
		},
		{
			Name:    "rollback",
			Aliases: []string{"r"},
			Usage:   "Database roll back",
			Action: func(c *cli.Context) {
				rollBack()
				fmt.Println("Version Rolled back by 1 step...")
			},
		},
	}
}

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func startserver(port string) {
	http.HandleFunc("/reception", reception.ReceptionHandler)
	http.HandleFunc("/doctor", doctor.DoctorHandler)
	// http.HandleFunc("/operation", OperationHandler)
	// http.HandleFunc("/reporting", ReportingHandler)
	http.ListenAndServe(":"+port, nil)
}

func migration() {
	databaseHost := os.Getenv("DOCTOR_DB_HOST")
	databasePort := os.Getenv("DOCTOR_DB_PORT")
	databaseUser := os.Getenv("DOCTOR_DB_USER")
	databasePass := os.Getenv("DOCTOR_DB_PASS")
	databaseName := os.Getenv("DOCTOR_DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		databaseHost, databasePort, databaseUser, databasePass, databaseName)
	//var err error
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
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

func rollBack() {
	databaseHost := os.Getenv("DOCTOR_DB_HOST")
	databasePort := os.Getenv("DOCTOR_DB_PORT")
	databaseUser := os.Getenv("DOCTOR_DB_USER")
	databasePass := os.Getenv("DOCTOR_DB_PASS")
	databaseName := os.Getenv("DOCTOR_DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		databaseHost, databasePort, databaseUser, databasePass, databaseName)
	//var err error
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
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
