package config

import (
	"strings"

	"github.com/spf13/viper"
)

type config struct {
	v *viper.Viper
}

func (c *config) SetBindings(bindings []string) {
	for _, binding := range bindings {
		c.v.BindEnv(binding)
	}
}

func (c *config) Load(name string, configType string, path string, o interface{}) error {
	c.v.SetConfigName(name)
	c.v.SetConfigType(configType)
	c.v.AddConfigPath(path)
	if err := c.v.ReadInConfig(); err != nil {
		return err
	}
	if err := c.v.Unmarshal(&o); err != nil {
		return err
	}
	c.v.AutomaticEnv()
	return nil
}

func New(prefix string) *config {
	v := viper.New()
	v.SetEnvPrefix(prefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return &config{v: v}
}
