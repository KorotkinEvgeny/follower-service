package viper

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var aliases = map[string]string{
	"ENV":             "env",
	"SERVER_SRV_HOST": "server.host",
	"SERVER_SRV_PORT": "server.port",

	"DB_HOST":     "db.host",
	"DB_PORT":     "db.port",
	"DB_NAME":     "db.name",
	"DB_USER":     "db.username",
	"DB_PASSWORD": "db.password",

	"LOGGER_LOG_LEVEL": "logger.level",

	"PARTITIONS": "partitions.count",
}

func ReadAndReturn() (*viper.Viper, error) {
	cfg := viper.New()

	viper.Set("server.read_timeout", 60*time.Second)
	viper.Set("server.write_timeout", 60*time.Second)

	if err := registerEnvironment(aliases, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

type ViperOptions struct {
	Func  func(*viper.Viper, string)
	Param string
}

func ReadFile(options ...ViperOptions) (*viper.Viper, error) {
	cfg := viper.New()
	for _, option := range options {
		option.Func(cfg, option.Param)
	}
	err := cfg.ReadInConfig()
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}
	registerAlias(aliases, cfg)
	return cfg, nil
}

func registerEnvironment(envList map[string]string, cfg *viper.Viper) error {
	for param, alias := range envList {
		if err := cfg.BindEnv(alias, param); err != nil {
			return err
		}
	}
	return nil
}

func registerAlias(envList map[string]string, cfg *viper.Viper) {
	for param, alias := range envList {
		cfg.RegisterAlias(alias, param)
	}
}
