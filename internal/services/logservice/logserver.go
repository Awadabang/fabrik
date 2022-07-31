package logservice

import (
	stlog "log"
	"os"
)

type LogServer struct {
	log *stlog.Logger
}

func NewLogServer(destination string) *LogServer {
	return &LogServer{
		log: stlog.New(FileLog(destination), "Fabrik: ", stlog.LstdFlags),
	}
}

type FileLog string

func (fl FileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return f.Write(data)
}

func (l *LogServer) Write(message string) {
	l.log.Printf("%v\n", message)
}
