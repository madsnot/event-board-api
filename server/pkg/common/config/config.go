package config

import (
	"time"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Port     string `mapstructure:"PORT"`
	DBURL    string `mapstructure:"DB_URL"`
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

type Config struct {
	DataBaseConf DBConfig
	TokenConf    TokenConfig
	EmailConf    EmailConfig
}

func LoadConfig() (conf Config, err error) {
	viper.AddConfigPath("./pkg/common/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&conf)

	return
}
