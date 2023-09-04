package config

import (
	"strings"

	"github.com/spf13/viper"
)

type viperCfg struct {
	v *viper.Viper
}

func NewViperConfig(prefix string) *viperCfg {
	v := viper.New()
	v.SetEnvPrefix(prefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return &viperCfg{
		v: v,
	}
}

func (c *viperCfg) SetBindings(bindings []string) {
	for _, binding := range bindings {
		c.v.BindEnv(binding)
	}
}

func (c *viperCfg) Load(name string, cfgType string, path string) error {
	c.v.SetConfigName(name)
	c.v.SetConfigType(cfgType)
	c.v.AddConfigPath(path)
	if err := c.v.ReadInConfig(); err != nil {
		return err
	}
	c.v.AutomaticEnv()
	return nil
}

func (c *viperCfg) Unmarshal(o interface{}) error {
	if err := c.v.Unmarshal(&o); err != nil {
		return err
	}
	return nil
}

func (c *viperCfg) LoadAndUnmarshal(name string, cfgType string, path string, o interface{}) error {
	if err := c.Load(name, cfgType, path); err != nil {
		return err
	}
	if err := c.Unmarshal(o); err != nil {
		return err
	}
	return nil
}
