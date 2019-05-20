// Package logging wraps logrus imported logger for quick replacement.
package logging

import (
	"github.com/sirupsen/logrus"
)

// Configuration is the container for this packages configuration.
// It is used within the pkg/configuration package.
type Configuration struct {
	DescribeCaller bool `json:"describe-caller"`
}

var (
	config = Configuration{}
)

// ApplyConfiguration applies specific configuration to this package.
func ApplyConfiguration(thisConfig Configuration) error {
	config = thisConfig

	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	logger.SetFormatter(formatter)

	Info("Applied configuration to logging package.")
	return nil
}

// DefaultConfiguration for this packages configuration.
func DefaultConfiguration() Configuration {
	return Configuration{
			DescribeCaller: true,
	}
}
