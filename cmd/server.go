package cmd

import "github.com/Awadabang/fabrik/internal/server"

func FabrikServe() {
	server := server.NewFabrikServer(
		server.WithName(),
	)
	server.Start()
}
