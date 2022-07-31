package services

import (
	"github.com/Awadabang/fabrik/internal/services/logservice"
	"github.com/Awadabang/fabrik/internal/services/registry"
	"github.com/Awadabang/fabrik/internal/services/server"
)

type Svc struct {
	FabrikServer *server.FabrikServer
	Registry     *registry.Registry
	LogService   *logservice.LogServer
}

func GenerateSrevices() *Svc {
	// Logserver
	logServer := logservice.NewLogServer()

	// Registry
	registry := registry.NewRegistry()
	go registry.Start()

	// Fabrik
	fabrikServer := server.NewFabrikServer(
		server.WithRegistyNode(&registry.ServiceNode),
	)
	fabrikServer.Start()

	return &Svc{
		FabrikServer: fabrikServer,
		Registry:     registry,
		LogService:   logServer,
	}
}
