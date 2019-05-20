package server

import "github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/logging"

// Configuration
type Configuration struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

var (
	config Configuration // config for this package.
)

// ApplyConfiguration will configure this component with its given configuration.
func ApplyConfiguration(thisConfig Configuration) error {
	config = thisConfig
	logging.Info("Applied configuration to server package.")
	return nil
}

// DefaultConfiguration for this packages configuration.
func DefaultConfiguration() Configuration {
	return Configuration{
		Host: "localhost",
		Port: 8080,
	}
}
