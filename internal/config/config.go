package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Test string `mapstructure:test`
}

func MustLoadConfig(fullpath string) Config {
	viper.SetConfigFile(fullpath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		panic("config file is not found in " + fullpath)
	}

	var cfg = Config{}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic("error at parsing config, check your config file")
	}
	return cfg
}
