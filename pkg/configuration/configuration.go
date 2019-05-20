// Package configuration is a simple application configuration provider.
package configuration

import (
	"encoding/json"
	"io/ioutil"

	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/database"
	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/logging"
	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/server"
)

// ApplicationConfiguration contains configuration for the whole application.
type ApplicationConfiguration struct {
	LogConfig      logging.Configuration  `json:"log-configuration"`
	ServerConfig   server.Configuration   `json:"server-configuration"`
	DatabaseConfig database.Configuration `json:"database-configuration"`
}

// Default configuration for this application.
func Default() *ApplicationConfiguration {
	return &ApplicationConfiguration{
		LogConfig:      logging.DefaultConfiguration(),
		ServerConfig:   server.DefaultConfiguration(),
		DatabaseConfig: database.DefaultConfiguration(),
	}
}

// Load a configuration from file in JSON format.
func Load(path string) (*ApplicationConfiguration, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var appConfig ApplicationConfiguration

	err = json.Unmarshal([]byte(file), &appConfig)
	if err != nil {
		return nil, err
	}

	return &appConfig, nil
}

// Save a configuration to file.
func Save(appConfig *ApplicationConfiguration, path string) error {
	file, err := json.MarshalIndent(appConfig, "", "    ") // JSON indentation with four spaces.
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, file, 0644)
	return err
}
