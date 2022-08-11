package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := NewApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
func NewApp() *cli.App {
	return &cli.App{
		Name: "cars-client",
		Commands: []*cli.Command{
			{
				Name: "create",
			},
			{
				Name: "get",
			},
			{
				Name: "list",
			},
			{
				Name: "get",
			},
		},
	}
}
