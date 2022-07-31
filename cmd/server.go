package cmd

import (
	"log"
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
	mux.HandleFunc("/log", handlers.LogHandler(svcContext))

	httpServer := &http.Server{
		Addr:    "0.0.0.0:8111",
		Handler: mux,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		svcContext.LogService.Write(err.Error())
		log.Println(err.Error())
	}
}
