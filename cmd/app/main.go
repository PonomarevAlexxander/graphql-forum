package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/PonomarevAlexxander/graphql-forum/internal/app"
	"github.com/PonomarevAlexxander/graphql-forum/internal/config"
	"github.com/PonomarevAlexxander/graphql-forum/internal/logger"
)

var (
	configFilePath = "config/app/config.yaml"
)

func main() {
	conf, err := config.ReadConfigFromYAML[app.Config](configFilePath)
	if err != nil {
		panic(fmt.Errorf("Read of config from '%s' failed: %w", configFilePath, err))
	}

	err = config.ValidateConfig(conf)
	if err != nil {
		panic(fmt.Errorf("'%s' parsing failed: %w", configFilePath, err))
	}

	logger.SetupGlobalLogger(&conf.LoggerConfig)
	slog.Info("Starting...")

	notify := make(chan error, 1)
	defer close(notify)

	app, err := app.NewApp(conf, notify)
	if err != nil {
		log.Fatal(err)
	}

	app.Start()
	defer func() {
		err := app.Stop()
		if err != nil {
			slog.Error("Error while shutting down", "error", err.Error())
		}
	}()

	interupt := make(chan os.Signal, 1)
	defer close(interupt)

	signal.Notify(interupt, os.Interrupt, syscall.SIGTERM)

	select {
	case serr := <-notify:
		slog.Error("Notified with app error", "error", serr.Error())
	case signl := <-interupt:
		slog.Info("Cought signal while App running", "signal", signl.String())
	}

	slog.Info("Shutting down...")
}
