package nmapreceiver

import (
	"time"

	"go.opentelemetry.io/collector/component"
)

// Config represents the receiver config settings within the collector's config.yaml
type Config struct {
	Interval         string `mapstructure:"interval"`
	ScannerArguments string `mapstructure:"scanner_arguments"`
	Fake             bool   `mapstructure:"fake"`
}

// Validate checks if the receiver configuration is valid
func (cfg *Config) Validate() error {
	// TODO: Add validation for interval.
	//TODO: Add validation for scanner arguments.

	return nil
}

func createDefaultConfig() component.Config {
	return &Config{
		Interval: (time.Minute * 1).String(),
	}
}
