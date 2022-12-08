package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Port            string        `mapstructure:"PORT"`
	DBURL           string        `mapstructure:"DB_URL"`
	HashAddition    string        `mapstructure:"HASH_ADDITION"`
	AccessTokenTTL  time.Duration `mapstructure:"ACCESS_TOKEN_TTL"`
	RefreshTokenTTL time.Duration `mapstructure:"REFRESH_TOKEN_TTL"`
	SigningKey      string        `mapstructure:"SIGNING_KEY"`
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
