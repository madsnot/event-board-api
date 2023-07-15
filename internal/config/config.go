package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DataBaseCfg DBConfig
	TokenCfg    TokenConfig
	EmailCfg    EmailConfig
	HashCfg     HashConfig
}

type DBConfig struct {
	DSN string `mapstructure:"DB_URL"`
}

type HashConfig struct {
	HashSalt string `mapstructure:"HASH_SALT"`
}

type TokenConfig struct {
	AccessTokenTTL  time.Duration `mapstructure:"ACCESS_TOKEN_TTL"`
	RefreshTokenTTL time.Duration `mapstructure:"REFRESH_TOKEN_TTL"`
	SigningKey      string        `mapstructure:"SIGNING_KEY"`
}

type EmailConfig struct {
	Address  string `mapstructure:"EMAIL_ADDR"`
	Password string `mapstructure:"EMAIL_PASS"`
	Host     string `mapstructure:"EMAIL_HOST"`
	Port     string `mapstructure:"EMAIL_PORT"`
}

func LoadConfig() (Config, error) {
	var cfg Config

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
