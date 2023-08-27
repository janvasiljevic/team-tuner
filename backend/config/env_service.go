package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Enviroment string

const (
	EnvDev  Enviroment = "dev"
	EnvProd Enviroment = "prod"
)

type Config struct {
	Environment Enviroment `validate:"required"`
	Server      struct {
		Port int `validate:"required,numeric"`
	}
	Database struct {
		ConnectionUrl string `validate:"required"`
	}
	Github struct {
		ClientId     string `validate:"required"`
		ClientSecret string `validate:"required"`
		RedirectUrl  string `validate:"required"`
	}
}

var LoadedConfig Config

func NewEnv() (*Config, error) {
	// new instance of viper, for reading and handling configuration
	vp := viper.New()

	// set config file options
	vp.SetConfigName("config")
	vp.SetConfigType("toml")
	vp.AddConfigPath("./env")

	var cfg Config

	// Read config file
	if err := vp.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file %v", err)
	}

	// Parse the config file
	if err := vp.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error parsing env file %v", err)
	}

	validate := validator.New()

	// Validate the the config struct
	if err := validate.Struct(&cfg); err != nil {
		return nil, fmt.Errorf("missing required attributes %v", err)
	}

	// Check if the environment is set correctly
	if cfg.Environment != "prod" && cfg.Environment != "dev" {
		return nil, fmt.Errorf("environment incorrectly set. Expecting 'prod' or 'dev'")
	}

	return &cfg, nil
}
