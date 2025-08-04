package config

import (
	"bytes"
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"os"
	"time"
)

//go:embed Defaults.yaml
var Defaults []byte

type Config struct {
	GRPCServer struct {
		Port     int           `mapstructure:"Port" validate:"required,min=200,max=65535"`
		Interval time.Duration `mapstructure:"Interval" validate:"required,min=1,max=99999"`
	} `mapstructure:"GRPC"`
	Logger struct {
		Level string `mapstructure:"Level"`
	} `mapstructure:"Logger"`
}

func LoadConfig() (Config, error) {
	var (
		cfg = Config{}
	)
	fullpath := fetchConfigPath()
	// There should be a good comment here.
	if fullpath == "" {
		return Config{}, errors.New("path to configuration file is not specified")
	}
	// There should be a good comment here.
	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		return Config{}, err
	}
	// There should be a good comment here.
	if err := LoadDefaults(); err != nil {
		return Config{}, err
	}
	viper.AutomaticEnv()
	viper.SetConfigFile(fullpath)
	// There should be a good comment here.
	if err := viper.MergeInConfig(); err != nil {
		fmt.Println("Configuration file is not specified. Default values will be used.")
	}
	// There should be a good comment here.
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}
	// There should be a good comment here.
	return cfg, nil
}

func fetchConfigPath() string {
	var configPath string
	flag.StringVar(&configPath, "config", "", "Path to configuration file")
	flag.Parse()

	if configPath == "" {
		fmt.Println("The path to the configuration file was not specified. " +
			"An Environment variable called CONFIG_PATH will be used.")
		configPath = os.Getenv("CONFIG_PATH")
	}

	return configPath
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
