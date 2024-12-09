package cli

import (
	"flag"
	"fmt"
	"os"
)

func New() Flags {
	flags := Flags{
		Verbose:    flag.Bool("v", false, "Display detailed server activity."),
		ConfigPath: flag.String("f", "./conf.toml", "The configuration file."),
	}

	flag.Parse()

	return flags
}

func (args Flags) ValidateFlags() (bool, string) {

	// Check if the given config file path exist.
	if _, err := os.Stat(*args.ConfigPath); err != nil {
		return false, fmt.Sprintf("The config file does not exist in %s", *args.ConfigPath)
	}

	return true, ""
}
