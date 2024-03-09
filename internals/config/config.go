package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port              string `mapstructure:"PORT"`
	DBUrl             string `mapstructure:"DBURL"`
	DBName            string `mapstructure:"DBNAME"`
	AdminSecurityKey  string `mapstructure:"ADMIN_TOKENKEY"`
	LaywerSecurityKey string `mapstructure:"LAYWER_TOKENKEY"`
	UserSecurityKey   string `mapstructure:"USER_TOKENKEY"`
}

func InitConfig() (*Config, error) {
	var c *Config

	viper.AddConfigPath("./")
	viper.SetConfigType("env")
	viper.SetConfigName("dev")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
