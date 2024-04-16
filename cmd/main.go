package main

import (
	"log/slog"
	"os"
	"time"

	"ipfs-file-api/internal/config"
	fileMigration "ipfs-file-api/internal/file/migration"
	"ipfs-file-api/internal/route"
	"ipfs-file-api/pkg/graceful"
	"ipfs-file-api/pkg/logger"

	"github.com/urfave/cli/v2"
)

func main() {
	config.LoadConfig()
	app := &cli.App{
		Name:     "ipfs file server",
		Usage:    "Start the ipfs file",
		Compiled: time.Now(),
		Before: func(ctx *cli.Context) error {
			logger.SetupLogger(true)

			if err := fileMigration.Init(ctx.Context); err != nil {
				return err
			}

			return nil
		},
		Action: func(c *cli.Context) error {
			manager := graceful.GetManager()
			route.Init()
			<-manager.Done()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
