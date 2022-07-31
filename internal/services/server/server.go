package server

import (
	"log"
	"sync"
	"time"
)

type FabrikServer struct {
	ServiceNode *sync.Map
}

type FabrikServeOption func(*FabrikServer)

func WithRegistyNode(serviceNode *sync.Map) FabrikServeOption {
	return func(fs *FabrikServer) {
		fs.ServiceNode = serviceNode
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
		fs.ServiceNode.Range(func(key, value any) bool {
			log.Printf("Service Name: %v, Addr: %v\n", key, value)
			return true
		})
	}
}
