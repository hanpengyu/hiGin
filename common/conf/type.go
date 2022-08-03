package conf

import "github.com/spf13/viper"

type Config struct {
	cfgViper *viper.Viper
}

type LogConfig struct {
	LogFile string `mapstructure:"log_file" json:"log_file"`
}
