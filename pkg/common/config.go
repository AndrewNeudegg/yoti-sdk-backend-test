package common

import (
	"fmt"

	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/configuration"
	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/database"
	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/logging"
	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/server"
)

// GetConfig will load the local configuration file.
func GetConfig(path string) *configuration.ApplicationConfiguration {
	appConfig, err := configuration.Load(path)
	if appConfig == nil && err != nil {
		appConfig = configuration.Default()
		configuration.Save(appConfig, path)
		return appConfig
	}
	return appConfig
}

// ApplyConfiguration will call to each package to update their internal states with the configuration.
func ApplyConfiguration(appConfig *configuration.ApplicationConfiguration) (err error) {
	err = logging.ApplyConfiguration(appConfig.LogConfig)
	if err != nil {
		return err
	}

	err = server.ApplyConfiguration(appConfig.ServerConfig)
	if err != nil {
		return err
	}

	err = database.ApplyConfiguration(appConfig.DatabaseConfig)
	if err != nil {
		return err
	}

	logging.Info(fmt.Sprintf("Loaded Configuration: %+v .", appConfig))
	return nil
}
