package database

import (
	"fmt"

	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/logging"
)

// Configuration is the container for this packages configuration.
// It is used within the pkg/configuration package.
type Configuration struct {
	DBName      string `json:"db-name"`
	LockTimeout int64  `json:"lock-timeout"`
}

var (
	config = Configuration{} // config for this package.
)

// ApplyConfiguration applies specific configuration to this package.
func ApplyConfiguration(thisConfig Configuration) error {
	config = thisConfig
	logging.Info("Applied configuration to database package.")
	err := connectDatabase()
	if err != nil {
		return err
	}
	if database == nil {
		return fmt.Errorf("database was not populated")
	}
	return nil
}

// DefaultConfiguration for this packages configuration.
func DefaultConfiguration() Configuration {
	return Configuration{
		DBName: "application.db",
	}
}
