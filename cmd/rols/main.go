package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/lightclient/rols/node"
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
		Name:     "rols",
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
					&cli.StringFlag{Name: "feed-oracle", DefaultText: "http://localhost:8545", Required: true},
					&cli.StringFlag{Name: "sequencer", DefaultText: "http://localhost:8645", Required: true},
					&cli.StringFlag{Name: "key", Required: true},
				},
				Action: func(ctx *cli.Context) error {
					output := colorable.NewColorableStderr()
					ostream := log.StreamHandler(output, log.TerminalFormat(true))
					glogger.SetHandler(ostream)

					config := &node.Config{
						FeedOracle: ctx.String("feed-oracle"),
						Sequencer:  ctx.String("sequencer"),
						SigningKey: ctx.String("key"),
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
