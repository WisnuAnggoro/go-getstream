package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config struct to implement model of this service's configuration
type Config struct {
	// Setting port for gin
	Port string `envconfig:"PORT" default:""`

	// GoStream
	GoStreamAPIKey    string `envconfig:"GOSTREAM_API_KEY" default:"h2rkj5b7hd2r"`
	GoStreamAPISecret string `envconfig:"GOSTREAM_API_SECRET" default:"6sexn67m88skfp2p7e8m34h3uvrr6589z3cpadneskf8mpvpas8suj7y57j9j5bx"`
	GoStreamAPIRegion string `envconfig:"GOSTREAM_API_REGION" default:"singapore"`
}

// Get to get defined configuration
func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
