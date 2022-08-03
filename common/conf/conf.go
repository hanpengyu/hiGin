package conf

import "github.com/spf13/viper"

var _conf *Config

func setConfig(c *Config) {
	_conf = c
}

func GetConfig() *viper.Viper {
	return _conf.cfgViper
}

func GetLogConfig() LogConfig {
	v := GetConfig().Sub("logger")
	if v == nil {
		panic("no log config")
	}

	var cfg LogConfig
	err := v.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
