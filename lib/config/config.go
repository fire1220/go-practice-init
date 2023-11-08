package config

import "github.com/spf13/viper"

type config struct {
	v *viper.Viper
}

var (
	_config = new(config)
)

func SetViper(v *viper.Viper) {
	getConfig().v = v
}

func getConfig() *config {
	return _config
}

func Viper() *viper.Viper {
	return getConfig().v
}

func (c *config) SetViper(v *viper.Viper) {
	c.v = v
}

func (c *config) GetViper() *viper.Viper {
	return c.v
}
