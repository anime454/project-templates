package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	HTTP       HTTPConfig       `mapstructure:"http"`
	App        AppConfig        `mapstructure:"app"`
	PostgreSQL PostgreSQLConfig `mapstructure:"postgresql"`
}

type HTTPConfig struct {
	ListenAddr string `mapstructure:"listen_addr" validate:"required"`
}

type AppConfig struct {
	Name string `mapstructure:"name" validate:"required"`
}

type PostgreSQLConfig struct {
	Host            string        `mapstructure:"host" validate:"required"`
	Port            string        `mapstructure:"port" validate:"required"`
	User            string        `mapstructure:"user" validate:"required"`
	Password        string        `mapstructure:"password" validate:"required"`
	DB              string        `mapstructure:"db" validate:"required"`
	MaxConns        int           `mapstructure:"max_conns"`
	MinConns        int           `mapstructure:"min_conns"`
	MaxConnLifetime time.Duration `mapstructure:"max_conn_lifetime"`
	MaxConnIdleTime time.Duration `mapstructure:"max_conn_idle_time"`
}

func Load() (Config, error) {
	v := viper.New()

	v.SetDefault("postgresql.max_conns", 10)
	v.SetDefault("postgresql.min_conns", 1)
	v.SetDefault("postgresql.max_conn_lifetime", "30m")
	v.SetDefault("postgresql.max_conn_idle_time", "5m")

	// File config/config.yaml (relative to working dir)
	APP_ENV := os.Getenv("APP_ENV")
	if APP_ENV == "" {
		APP_ENV = "dev"
	}
	v.SetConfigName(APP_ENV)
	v.SetConfigType("yaml")
	v.AddConfigPath("config") // <repo-root>/config/config.yaml
	v.AddConfigPath(".")      // optionally allow ./config.yaml

	// ENV overrides
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return Config{}, fmt.Errorf("read config: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return Config{}, fmt.Errorf("unmarshal config: %w", err)
	}

	return cfg, nil
}
