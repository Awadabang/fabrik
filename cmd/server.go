package cmd

import (
	"github.com/Awadabang/fabrik/internal/registry"
	"github.com/Awadabang/fabrik/internal/server"
)

func FabrikServe() {
	// Fabrik
	server := server.NewFabrikServer(
		server.WithName(),
	)
	go server.Start()

	// Registry
	registry := registry.NewRegistry()
	registry.Start()
}
