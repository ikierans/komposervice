package main

import (
	"example/komposervice/cmd/server"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func NewClient() *cli.App {
	_app := &cli.App{
		Name:        "komposervice",
		Usage:       "komposervice",
		Version:     "0.0.1",
		Description: "API server",
		Commands:    server.Command,
		// Flags:       app.Flag,
	}

	sort.Sort(cli.FlagsByName(_app.Flags))
	sort.Sort(cli.CommandsByName(_app.Commands))

	return _app
}

func main() {
	client := NewClient()

	if err := client.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
