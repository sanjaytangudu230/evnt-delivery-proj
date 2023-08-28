package main

import (
	"eventDelivery/commands"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Event Delivery App"
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		commands.StartCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Print(err)
	}
}
