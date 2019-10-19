package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	// ErrConfigNotFound
	ErrConfigNotFound = errors.New("webconv: config file not found")

	// ErrConfigDecode
	ErrConfigDecode = errors.New("webconv: unable to decode config file")
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
		log.Panic(fmt.Errorf("%v: %w", ErrConfigNotFound, err))
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Panic(fmt.Errorf("%v: %w", ErrConfigDecode, err))
	}

	return &conf
}
