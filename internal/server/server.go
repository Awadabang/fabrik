package server

import (
	"log"
	"sync"
	"time"
)

type FabrikServer struct {
	Node sync.Map
}

type FabrikServeOption func(*FabrikServer)

func WithName() FabrikServeOption {
	return func(fs *FabrikServer) {
		fs.Node.Store("ymonitor", "localhost:19668")
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
		fs.Node.Range(func(key, value any) bool {
			log.Println(key, value)
			return true
		})
	}
}
