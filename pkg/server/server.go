package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/logging"
)

type SimpleServer struct {
	sync.WaitGroup
	http.Server
	shutdownReq chan bool
	reqCount    uint32
	hostString  string
}

// newServer will create a new instance of a server.
func NewServer() *SimpleServer {
	//create server
	s := &SimpleServer{
		WaitGroup: sync.WaitGroup{},
		Server: http.Server{
			Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		shutdownReq: make(chan bool),
	}
	router := NewRouter()
	s.Handler = router
	return s
}

// waitShutdown will kill a running server.
func (s *SimpleServer) waitShutdown() {
	<-s.shutdownReq
	logging.Info("Stoping http server ...")

	//Create shutdown context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//shutdown the server
	err := s.Shutdown(ctx)
	if err != nil {
		logging.Errors(fmt.Sprintf("Shutdown request error: %v", err))
	}
	s.WaitGroup.Done()
}

// Stop will stop a running HTTP server.
func (s *SimpleServer) Stop() {
	if s != nil {
		s.shutdownReq <- true
	}
}

// Start will create and start a new HTTP Server.
func (s *SimpleServer) Start() {
	s.WaitGroup.Add(1)
	go func() {
		go s.waitShutdown()
		logging.Info(fmt.Sprintf("Server is starting on: %s", fmt.Sprintf("http://%s:%d", config.Host, config.Port)))
		err := s.ListenAndServe()
		if err != nil {
			logging.Errors(fmt.Sprintf("Listen and serve: %v", err))
		}
	}()
}

// Await will block the calling goroutine until the server exits.
func (s *SimpleServer) Await() {
	if s != nil {
		s.WaitGroup.Wait()
	}
}
