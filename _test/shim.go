package blackbox

import (
	"fmt"
	"time"

	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/common"
	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/configuration"
	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/server"
)

// ServerShim is a struct that contains an implementation of the core application.
type ServerShim struct {
	config *configuration.ApplicationConfiguration // Stores the configuration for this test.
	srv    *server.SimpleServer                    // Sever location, not populated until targeted.
}

// LoadConfig provides a helper method for loading application configuration within
// blackbox tests.
func (a *ServerShim) LoadConfig() (err error) {
	a.config, err = configuration.Load("../config.json")
	if err != nil {
		return err
	}
	a.config.DatabaseConfig.DBName = fmt.Sprintf("../%s", a.config.DatabaseConfig.DBName)
	err = common.ApplyConfiguration(a.config)
	return err
}

// StartServer provides a helper method for starting the server from the blackbox test
// environment.
func (a *ServerShim) StartServer() {
	a.srv = server.NewServer()
	go a.srv.Start()
	time.Sleep(time.Duration(1) * time.Millisecond)
}

// StopServer provides a helper method for stopping the server from the blackbox test
// environment.
func (a *ServerShim) StopServer() {
	common.Tidy()
	go a.srv.Stop()
	a.AwaitServer()
	time.Sleep(time.Duration(1) * time.Millisecond)
	shim = nil
}

// AwaitServer does not return until the server exits.
func (a *ServerShim) AwaitServer() {
	a.srv.Await()
}

// Host returns a formatted string version of the host name and port information.
func (a *ServerShim) Host() string {
	return fmt.Sprintf("http://%s:%d", a.config.ServerConfig.Host, a.config.ServerConfig.Port)
}

// serverShimFactory produces new shims for consumption.
func serverShimFactory() (*ServerShim, error) {
	srv := ServerShim{}
	err := srv.LoadConfig()
	return &srv, err
}

// ------------------------------
// Shim state management
//

var (
	shim *ServerShim // Stores the shim that is currently in use.
)

// GetShim will return the current testing application struct, or create a new one.
func GetShim() (*ServerShim, error) {
	var err error
	if shim == nil {
		shim, err = serverShimFactory()
	}
	return shim, err
}
