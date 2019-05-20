package common

import (
	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/database"
	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/logging"
)

// Tidy will encourage all packages to clear up whatever they are doing.
func Tidy() {
	logging.Info("Requesting that packages tidy themselves.")
	database.Tidy()
}
