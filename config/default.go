package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUri    string `mapstructure:"MONGODB_LOCAL_URI"`
	RedisUri string `mapstructure:"REDIS_URL"`
	Port     string `mapstructure:"PORT"`
}

// The Golang viper package uses github.com/mitchellh/mapstructure under the hood for unmarshaling values which uses mapstructure tags by default.
//MONGODB_LOCAL_URI REDIS_URL and PORT are coming fro0m app.env

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
