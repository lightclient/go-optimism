package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/lightclient/go-optimism/node"
	"github.com/mattn/go-colorable"
	"github.com/urfave/cli/v2"
)

var glogger *log.GlogHandler

func init() {
	glogger = log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.LvlInfo)
	log.Root().SetHandler(glogger)
}

func main() {
	app := &cli.App{
		Name:     "go-optimism",
		Usage:    "An adapter between L1 and rollup clients.",
		Version:  "v0.0.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "lightclient",
				Email: "lightclient@protonmail.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "start adapter",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "feed-oracle", DefaultText: "http://localhost:8545"},
					&cli.StringFlag{Name: "feed-contract", DefaultText: "0x4BF681894abEc828B212C906082B444Ceb2f6cf6"},
				},
				Action: func(ctx *cli.Context) error {
					output := colorable.NewColorableStderr()
					ostream := log.StreamHandler(output, log.TerminalFormat(true))
					glogger.SetHandler(ostream)

					config := &node.Config{
						FeedProvider: ctx.String("feed-provider"),
					}

					node, err := node.New(config)
					if err != nil {
						return err
					}

					node.Start()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
