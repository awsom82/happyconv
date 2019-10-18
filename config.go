package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Hostname     string
	Port         uint16
	RateLimit    float64
	RateLimitTTL int
}

func NewConfig(filename string) *Config {

	var conf Config

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return &conf
}
