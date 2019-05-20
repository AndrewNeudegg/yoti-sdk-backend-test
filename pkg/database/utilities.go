package database

import "github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/logging"

// Tidy will close any outstanding database connections.
func Tidy() {
	logging.Info("Tidying database package.")
	releaseDatabase()
	logging.Info("Done tidying database package.")
}
