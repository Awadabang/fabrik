package signal

import (
	"os"
	"os/signal"
	"syscall"
)

var SigChan = make(chan os.Signal, 2)

func Catch() {
	signal.Notify(SigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	<-SigChan
	signal.Stop(SigChan)
	close(SigChan)
}
