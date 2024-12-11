package config

import (
	"errors"

	"github.com/BurntSushi/toml"
)

var (
	Meta toml.MetaData
	conf Configuration = Configuration{}

	ErrMissingServerDirective  = errors.New("missing 'server' key in the configuration")
	ErrMissingWorkerDirectives = errors.New("missing 'worker' key/s in the configuration")
)

func New(path string) (*Configuration, error) {

	var err error

	Meta, err = toml.DecodeFile(path, &conf)

	return &conf, err
}

func (cfg *Configuration) TestConfig() (bool, error) {

	if !Meta.IsDefined("server") {
		return false, ErrMissingServerDirective
	}

	if !Meta.IsDefined("workers") {
		return false, ErrMissingWorkerDirectives
	}

	return true, nil
}
