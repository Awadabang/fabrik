package server

import (
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
