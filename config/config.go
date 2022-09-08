package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		PoolMax            int    `env-required:"true" yaml:"pool_max" env:"POSTGRESQL_POOL_MAX"`
		MaxConnAttempts    int    `env-required:"true" yaml:"max_attempts" env:"POSTGRESQL_MAX_ATTEMPTS"`
		User               string `env-required:"true" env:"POSTGRESQL_USER"`
		Pass               string `env-required:"true" env:"POSTGRESQL_PASSWORD"`
		Host               string `env-required:"true" env:"POSTGRESQL_HOST"`
		Port               int    `env-required:"true" env:"POSTGRESQL_PORT"`
		Database           string `env-required:"true" env:"POSTGRESQL_DATABASE"`
		ConnSecondsTimeout int    `env-required:"true" yaml:"conn_seconds_timeout"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
