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
	logServer := logservice.NewLogServer("info.log")
	logServer.Write("LogServer Constructed")

	// Registry
	registry := registry.NewRegistry()
	logServer.Write("Registry Constructed")

	// Fabrik
	fabrikServer := server.NewFabrikServer(
		server.WithRegisty(registry),
	)
	logServer.Write("FabrikServer Constructed")

	logServer.Write("Fabrik Constructed Completly")
	return &Svc{
		FabrikServer: fabrikServer,
		Registry:     registry,
		LogService:   logServer,
	}
}
