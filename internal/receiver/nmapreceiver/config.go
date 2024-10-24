package nmapreceiver

import (
	"time"

	"go.opentelemetry.io/collector/component"
)

// Config represents the receiver config settings within the collector's config.yaml
type Config struct {
	Interval         string `mapstructure:"interval"`
	ScannerArguments string `mapstructure:"scanner_arguments"`
}

// Validate checks if the receiver configuration is valid
func (cfg *Config) Validate() error {
	interval, _ := time.ParseDuration(cfg.Interval)
	// if interval.Minutes() < 1 {
	// 	return errors.New("when defined, the interval has to be set to at least 1 minute (1m)")
	// }
	_ = interval

	//TODO: Add validation for scanner arguments.

	return nil
}

func createDefaultConfig() component.Config {
	return &Config{
		Interval: (time.Minute * 1).String(),
	}
}
