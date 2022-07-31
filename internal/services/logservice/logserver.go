package logservice

import (
	stlog "log"
	"os"
)

type LogServer struct {
}

func NewLogServer() *LogServer {
	return &LogServer{}
}

var log *stlog.Logger

type FileLog string

func (fl FileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return f.Write(data)
}

func Run(destination string) {
	log = stlog.New(FileLog(destination), "go", stlog.LstdFlags)
}

func (l *LogServer) Write(message string) {
	log.Printf("%v\n", message)
}
