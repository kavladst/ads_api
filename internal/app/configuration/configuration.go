package configuration

import (
	"github.com/spf13/viper"
	"os"
)

type Configuration struct {
	AppHost string `mapstructure:"app_host"`
	AppPort string `mapstructure:"app_port"`

	DBHost     string `mapstructure:"db_host"`
	DBPort     string `mapstructure:"db_port"`
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBName     string `mapstructure:"db_name"`
}

func New() (*Configuration, error) {
	configFileName, isExists := os.LookupEnv("CONFIG_FILE")
	if !isExists {
		configFileName = "config"
	}
	configFilePath, isExists := os.LookupEnv("CONFIG_FILE_PATH")
	if !isExists {
		configFilePath = "config/"
	}
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Configuration{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
