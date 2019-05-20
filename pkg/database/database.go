package database

import (
	"fmt"
	"time"

	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/logging"
	"github.com/boltdb/bolt"
)

var (
	database *bolt.DB // database stores the package reference to the db instance.
)

// connectDatabase will attempt to read the database from the local disk or
// create a new one.
// If the database exists and is locked by another process, this process will wait for
// config.LockTimeout seconds before returning an error.
func connectDatabase() (err error) {
	logging.Info("Attempting to connect to database.")
	database, err = bolt.Open(
		config.DBName, 0600, &bolt.Options{Timeout: time.Duration(config.LockTimeout * time.Second.Nanoseconds())})
	if database != nil {
		logging.Info("Connection made to database.")
	}
	return err
}

// releaseDatabase unlocks the database, and must be called when the process exits.
func releaseDatabase() error {
	logging.Info("Attempting to release database.")
	if database != nil {
		return database.Close()
	}
	return nil
}

// WriteInput is a simplified proxy method for writing something to the Input bucket.
func WriteInput(key []byte, valBytes []byte) error {
	// Write the data to the database, under the hoover bucket.
	logging.Info("Writing an input value to the Database.")
	return write(key, valBytes, []byte("hoover-input"))
}

// WriteOutput is a simplified proxy method for writing something to the Output bucket.
func WriteOutput(key []byte, valBytes []byte) error {
	// Write the data to the database, under the hoover bucket.
	logging.Info("Writing an output value to the Database.")
	return write(key, valBytes, []byte("hoover-output"))
}

// write is a helper method for writing any k:v pair to any bucket.
func write(key []byte, val []byte, bucketName []byte) error {
	err := database.Update(func(tx *bolt.Tx) error {
		bk, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return fmt.Errorf("failed to create bucket: %v", err)
		}
		if err := bk.Put(key, val); err != nil {
			return fmt.Errorf("failed to insert '%s' with val '%s': %v", string(key), string(val), err)
		}
		logging.Info(fmt.Sprintf(
			"Inserted record with ID:'%s' and value: '%s' to the database bucket: '%s'",
			string(key), string(val), string(bucketName)))
		return nil
	})
	return err
}
