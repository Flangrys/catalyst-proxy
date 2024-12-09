package config

import (
	"github.com/BurntSushi/toml"
)

var (
	Meta toml.MetaData
)

func New(path string) (Configuration, error) {

	var (
		err  error
		conf Configuration = Configuration{}
	)

	Meta, err = toml.DecodeFile(path, &conf)

	return conf, err
}

func (cfg *Configuration) ValidateConfig() (bool, string) {

	if !Meta.IsDefined("server") {
		return false, "The 'server' table is not defined."
	}

	return true, ""
}
