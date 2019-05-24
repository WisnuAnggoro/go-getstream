package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config struct to implement model of this service's configuration
type Config struct {
	// Setting port for gin
	Port string `envconfig:"PORT" default:""`

	// GoStream
	GoStreamAPIKey    string `envconfig:"GOSTREAM_API_KEY" default:""`
	GoStreamAPISecret string `envconfig:"GOSTREAM_API_SECRET" default:""`
	GoStreamAPIRegion string `envconfig:"GOSTREAM_API_REGION" default:""`
}

// Get to get defined configuration
func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
