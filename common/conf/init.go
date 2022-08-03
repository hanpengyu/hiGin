package conf

import "github.com/spf13/viper"

func InitCfg(cfg string) {
	if cfg == "" {
		cfg = "./config/test.toml" // 默认配置文件
	}

	v := viper.New()
	v.SetConfigFile(cfg)
	err := v.ReadInConfig()
	if err != nil {
		panic("配置文件解析失败")
	}

	conf := &Config{cfgViper: v}
	setConfig(conf)
}
