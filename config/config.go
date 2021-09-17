package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

const configFilePath = "./config.yaml"

type Config struct {
	Port int `yaml:"server.port"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()             // read value ENV variable

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	err = viper.Unmarshal(&config)
	return config, err
}

var (
	config *Config
)

func InitConfig() {
	fmt.Print("Loading -> Config...")
	config_, err := LoadConfig()
	if err != nil {
		log.Fatal("Cannot load loadConfig: ", err)
	}

	config = &config_
	fmt.Println("DONE")
}

func GetConfig() *Config {
	return config
}