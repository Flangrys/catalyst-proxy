package config_test

import (
	"os"
	"testing"

	"github.com/flangrys/catalyst-proxy/config"
)

const TEST_CONFIG_PATH = "../config.toml"

func TestConfigConstructor_withValidConfPath_expectingNoFailure(t *testing.T) {

	var (
		err  error
		conf *config.Configuration
	)

	if _, err = os.Stat(TEST_CONFIG_PATH); err != nil {
		t.Fail()
	}

	conf, err = config.New(TEST_CONFIG_PATH)

	if err != nil {
		t.Errorf("test failed: %s", err)
	}

	t.Logf("Displaying server config: %#v", conf.Server)
}

func TestConfigValidation_withValidConfig_expectingNoFailure(t *testing.T) {
	var (
		err  error
		ok   bool
		conf *config.Configuration
	)

	if _, err = os.Stat(TEST_CONFIG_PATH); err != nil {
		t.Error("Failed to locate the config file.")
	}

	conf, err = config.New(TEST_CONFIG_PATH)

	if err != nil {
		t.Errorf("Config constructor failed because: %s", err)
	}

	if ok, err = conf.TestConfig(); !ok {
		t.Errorf("Failed to validate the config: %s", err)
	}

}
