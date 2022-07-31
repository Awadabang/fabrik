package main

import (
	"log"

	"github.com/Awadabang/fabrik/cmd"
	"github.com/Awadabang/fabrik/pkg/signal"
)

func main() {
	go signal.Catch()
	cmd.FabrikServe()

	<-cmd.GlobalChan
	log.Println("进程退出：服务已关闭")
}
