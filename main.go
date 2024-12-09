package main

import (
	"os"

	"github.com/flangrys/catalyst-proxy/cli"
	"github.com/flangrys/catalyst-proxy/config"
	logger "github.com/sirupsen/logrus"
	formatter "github.com/x-cray/logrus-prefixed-formatter"
)

const (
	Version = "0.0.0-dev"
)

func main() {

	// Setup logger.
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logger.DebugLevel)
	logger.SetFormatter(&formatter.TextFormatter{
		ForceColors: true,
	})

	// Setup flags.
	flags := cli.New()

	logger.Infof("Locating config file at: %s", *flags.ConfigPath)

	if ok, message := flags.ValidateFlags(); !ok {
		logger.Fatalf("Failed to validate flags: %s", message)
	}

	// Setup config.
	conf, err := config.New(*flags.ConfigPath)

	if err != nil {
		logger.Fatalf("An error ocurred during the parsing: %s", err)
	}

	if ok, message := conf.ValidateConfig(); !ok {
		logger.Fatalf("Invalid config file: %s", message)

	} else {
		logger.Info("Flags succesfully validated.")
	}
}
