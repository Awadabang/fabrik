package server

import (
	"log"
	"time"

	"github.com/Awadabang/fabrik/internal/services/registry"
)

type FabrikServer struct {
	Registry *registry.Registry
}

type FabrikServeOption func(*FabrikServer)

func WithRegisty(registry *registry.Registry) FabrikServeOption {
	return func(fs *FabrikServer) {
		fs.Registry = registry
	}
}

func NewFabrikServer(opts ...FabrikServeOption) *FabrikServer {
	server := &FabrikServer{}

	for _, apply := range opts {
		apply(server)
	}

	return server
}

func (fs *FabrikServer) Start() {
	heartBeat := time.NewTicker(10 * time.Second)
	for {
		<-heartBeat.C
		log.Println("FabrikServer HeartBeat...")

		fs.Registry.Mutex.RLock()
		for _, registration := range fs.Registry.Registrations {
			log.Printf("Service Name: %v, Addr: %v\n", registration.ServiceName, registration.ServiceURL)
		}
		fs.Registry.Mutex.RUnlock()
	}
}
