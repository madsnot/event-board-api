package config

import (
	"time"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	Host          string        `env:"SERVER_HOST"`
	Port          string        `env:"SERVER_PORT"`
	ReadTimeout   time.Duration `env:"SERVER_READ_TIMEOUT"`
	WriteTimeout  time.Duration `env:"SERVER_WRITE_TIMEOUT"`
	MigrationsCfg MigrationsConfig
	PostgresCfg   PostgresConfig
	TokenCfg      TokenConfig
	HashCfg       HashConfig
}

type MigrationsConfig struct {
	AutoRun            bool   `env:"MIGRATION_AUTO_RUN" envDefault:"false"`
	DestinationVersion string `env:"MIGRATION_DESTINATION_VERSION" envDefault:"last"`
	VersionTable       string `env:"MIGRATION_VERSION_TABLE" envDefault:"public.schema_version"`
	LogLevel           string `env:"MIGRATION_LOG_LEVEL"`
	Path               string `env:"MIGRATION_PATH"`
}

type PostgresConfig struct {
	DSN string `env:"DSN"`
}

type HashConfig struct {
	HashSalt string `env:"HASH_SALT"`
}

type TokenConfig struct {
	AccessTokenTTL  time.Duration `env:"ACCESS_TOKEN_TTL"`
	RefreshTokenTTL time.Duration `env:"REFRESH_TOKEN_TTL"`
	SigningKey      string        `env:"TOKEN_SIGNING_KEY"`
}

func LoadConfig() (Config, error) {
	var cfg Config

	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
