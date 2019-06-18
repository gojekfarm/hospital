package main

import (
	"fmt"
	"hospital/server"
	"hospital/storage"
	"hospital/surgeon"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "Hospital - An autonomous healing System"
	app.Usage = "fix fault/failure in system"
	app.Author = "Jainam | Dilip"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "runserver",
			Aliases: []string{"s"},
			Usage:   "Starts the server",
			Action: func(c *cli.Context) {
				server.StartServer()
			},
		},
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "Database migration",
			Action: func(c *cli.Context) {
				storage.Migration()
				fmt.Println("Migrating...")
			},
		},
		{
			Name:    "downonestep",
			Aliases: []string{"dw1"},
			Usage:   "Database roll back",
			Action: func(c *cli.Context) {
				storage.DownOneStep()
				fmt.Println("Version Rolled back by 1 step...")
			},
		},
		{
			Name:    "startsurgeon",
			Aliases: []string{"surgeon"},
			Usage:   "Surgeon starts long polling",
			Action: func(c *cli.Context) {
				fmt.Println("Started Surgeon...")
				surgeon.LongPolling()
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
