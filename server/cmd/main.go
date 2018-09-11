package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/nazo/binsen/server/app/http"
	"github.com/nazo/binsen/server/app/services"
	"github.com/nazo/binsen/server/lib/db"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Run HTTP server",
			Action: func(c *cli.Context) error {
				http.Server()
				return nil
			},
		},
		{
			Name:    "users",
			Aliases: []string{"u"},
			Usage:   "users",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new user",
					Action: func(c *cli.Context) error {
						_, err := db.NewDB(&db.NewDBConfig{
							Host:     "postgres",
							Port:     "5432",
							User:     "binsen",
							Password: "binsen",
							Database: "binsen_dev",
							SSLMode:  "disable",
						})
						if err != nil {
							panic(err)
						}
						group, err := services.NewGroupService().GetDefaultGroup()
						if err != nil {
							panic(err)
						}
						role, err := services.NewRoleService().GetDefaultRole()
						if err != nil {
							panic(err)
						}
						email := c.Args().First()
						_, err = services.NewUserService().CreateUser(email, email, role, group)
						if err != nil {
							return cli.NewExitError(fmt.Sprintf("user create failed: %s", err), 2)
						}
						fmt.Printf("new user created. %s\n", email)
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing user",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
