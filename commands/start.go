package commands

import (
	"context"
	"eventDelivery/server"
	"github.com/urfave/cli"
)

var StartCommand = cli.Command{
	Name:    "start",
	Aliases: []string{"up"},
	Usage:   "Starts web server",
	Action:  startAction,
}

// startAction start the web server and initializes the daemon
func startAction(ctx *cli.Context) error {
	ctxt := context.Background()

	server.Start(ctxt)
	return nil
}
