package bootstrap

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"hi-gin/common/conf"
)

func InitConfig(f *pflag.FlagSet) {
	if f == nil {
		pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
		pflag.Parse()
		f = pflag.CommandLine
	}

	rootViper := viper.New()
	err := rootViper.BindPFlags(f)
	if err != nil {
		panic("配置文件解析失败")
	}

	cfg := rootViper.GetString("config")
	conf.InitCfg(cfg)
}
