package webconv

import (
	"time"
)

type Config struct {
	Hostname     string
	Port         uint // pflag doesnt support uint16 out of the box
	RateLimit    float64
	RateLimitTTL time.Duration
}

func NewConfig() *Config {
	ttl, _ := time.ParseDuration("5s")
	return &Config{"localhost", 8080, 2e5, ttl}
}
