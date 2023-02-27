package config

import (
	"path"
	"path/filepath"
	"testing"
)

type appConfig struct {
	Service serviceConfig
}

type serviceConfig struct {
	Name string
}

func TestConfig(t *testing.T) {
	var cfg *appConfig
	c := New("prefix")

	bindings := []string{"name"}
	c.SetBindings(bindings)

	err := c.Load("config", "yaml", filepath.Join("fixtures"), &cfg)
	if err != nil {
		t.Error(err)
	}

	if cfg.Service.Name != "name" {
		t.Error(err)
	}
}

func TestConfigFileError(t *testing.T) {
	var cfg *appConfig
	c := New("prefix")

	bindings := []string{"name"}
	c.SetBindings(bindings)

	err := c.Load("ccoonnffiigg", "yaml", path.Join(".", "fixtures"), &cfg)
	if err == nil {
		t.Fatal("should have failed to find config file")
	}
}

func TestConfigUnmarshallError(t *testing.T) {
	var cfg string
	c := New("prefix")

	bindings := []string{"name"}
	c.SetBindings(bindings)

	err := c.Load("config", "yaml", path.Join(".", "fixtures"), &cfg)
	if err == nil {
		t.Fatal("should have failed to unmarshall config")
	}
}
