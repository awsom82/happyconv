package webconv

import (
	"time"
)

var (
	// ErrConfigNotFound
	ErrConfigNotFound = errors.New("webconv: config file not found")

	// ErrConfigDecode
	ErrConfigDecode = errors.New("webconv: unable to decode config file")
)

type Config struct {
	Hostname     string
	Port         uint // pflag doesnt support uint16 out of the box
	RateLimit    float64
	RateLimitTTL time.Duration
	KeepAlive    bool
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewConfig() *Config {
	return &Config{"localhost", 8080, 2e5, 5e9, false, 5e9, 10e9}
}
