package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var (
	instance Config
	once     sync.Once
)

func initConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&instance)
	if err != nil {
		panic(fmt.Errorf("config file can not unmarshal: %w", err))
	}

}

func GetReadonlyConfig() Config {
	once.Do(initConfig)

	return instance
}
