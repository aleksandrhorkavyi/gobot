package main

import (
	"fmt"
	"github.com/urfave/cli"
	"gobot/config"
	"gobot/db"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	conf := config.New()

	app.Commands = []*cli.Command{
		{
			Name:    "run",
			Usage:   "Run gobot demon",
			Action:  func(c *cli.Context) error {
				Run(conf)
				return nil
			},
		},
		{
			Name:        "db",
			Usage:       "options for task templates",
			Subcommands: []*cli.Command{
				{
					Name:  "deploy",
					Usage: "Deploy new db schema",
					Action: func(c *cli.Context) error {
						//c.Args().First()
						db.Deploy(c, conf)
						return nil
					},
				},
				{
					Name:  "migrate",
					Usage: "Update DB schema",
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
