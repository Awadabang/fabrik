package cmd

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Awadabang/fabrik/internal/handlers"
	"github.com/Awadabang/fabrik/internal/services"
	"github.com/Awadabang/fabrik/pkg/signal"
)

var GlobalChan = make(chan struct{})

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

	go func() {
		<-signal.SigChan

		timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer timeoutCancel()

		if err := httpServer.Shutdown(timeoutCtx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Println("Http服务暴力停止，一般是达到超时context的时间当前还有尚未完结的http请求：" + err.Error())
		} else {
			log.Println("Http服务优雅停止")
		}

		close(GlobalChan)
	}()

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		svcContext.LogService.Write("Http服务异常 " + err.Error())
		log.Println("Http服务异常", err.Error())
	}
}
