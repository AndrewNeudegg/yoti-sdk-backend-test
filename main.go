package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/common"
	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/logging"
	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/server"
)

func main() {
	// Configure and Initialise all packages that require it.
	appConfig := common.GetConfig("config.json")
	err := common.ApplyConfiguration(appConfig)
	if err != nil {
		logging.Error(err)
		os.Exit(1)
	}

	// If a user tries harder than this then
	// they bear the responsibility of tidying
	// the applications state, using the log messages
	// as guidance.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	// Prepare an OnExit goroutine.
	go func() {
		s := <-sigc
		logging.Info(fmt.Sprintf("Recieved OS signal: %v", s))
		common.Tidy()
		logging.Info("The application is now exiting.")
		os.Exit(0)
	}()

	// Launch a simple server.
	srv := server.NewServer()
	srv.Start()
	// Block the main goroutine indefinitely.
	srv.Await()
}
