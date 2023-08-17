package initializers

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST_PRODUCTION"`
	DBUserName     string `mapstructure:"POSTGRES_USER_PRODUCTION"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD_PRODUCTION"`
	DBName         string `mapstructure:"POSTGRES_DB_PRODUCTION"`
	DBPort         string `mapstructure:"POSTGRES_PORT_PRODUCTION"`
	ServerPort     string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

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
