package cmd

import (
	"net/http"

	"github.com/Awadabang/fabrik/internal/handlers"
	"github.com/Awadabang/fabrik/internal/services"
)

func FabrikServe() {
	svcContext := services.GenerateSrevices()

	//Http Server
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/register", handlers.RegisterHandler(svcContext))

	httpServer := &http.Server{
		Addr:    "0.0.0.0:8111",
		Handler: mux,
	}

	go httpServer.ListenAndServe()

}
