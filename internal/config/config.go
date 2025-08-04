package config

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

//go:embed Defaults.yaml
var Defaults []byte

type Config struct {
	Server struct {
		Port     int `mapstructure:"Port" validate:"required,min=200,max=65535"`
		Interval int `mapstructure:"Interval" validate:"required,min=1,max=99999"`
	} `mapstructure:"Server"`
	Logger struct {
		Level string `mapstructure:"Level"`
	} `mapstructure:"Logger"`
}

func LoadConfig(fullpath string) (Config, error) {
	var (
		cfg = Config{}
	)
	if err := LoadDefaults(); err != nil {
		return Config{}, err
	}
	viper.AutomaticEnv()
	viper.SetConfigFile(fullpath)
	if err := viper.MergeInConfig(); err != nil {
		fmt.Println("Configuration file is not specified. Default values will be used.")
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func LoadDefaults() error {
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewReader(Defaults)); err != nil {
		return errors.New(fmt.Sprintf("Failed to read configuration defaults fields (%v)", err))
	}
	return nil
}

func ValidateConfig(config Config) error {
	validate := validator.New()
	return validate.Struct(config)
}
