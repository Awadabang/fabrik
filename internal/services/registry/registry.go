package registry

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Awadabang/fabrik/internal/services/logservice"
)

type Registration struct {
	ServiceName string
	ServiceURL  string
}

type Registry struct {
	Registrations []Registration
	Client        *http.Client
	Log           *logservice.LogServer

	Mutex sync.RWMutex
}

type RegistryOption func(*Registry)

func WithClient(client *http.Client) RegistryOption {
	return func(r *Registry) {
		r.Client = client
	}
}

func WithLog(log *logservice.LogServer) RegistryOption {
	return func(r *Registry) {
		r.Log = log
	}
}

func NewRegistry(opts ...RegistryOption) *Registry {
	registry := &Registry{}
	for _, appply := range opts {
		appply(registry)
	}

	return registry
}

func (r *Registry) AliveCheck() {
	timer := time.NewTicker(10 * time.Second)
	for {
		<-timer.C
		r.Mutex.RLock()
		for _, registration := range r.Registrations {
			log.Printf("AliveCheck: service name: %v, addr: %v\n", registration.ServiceName, registration.ServiceURL)
			resp, err := r.Client.Get(registration.ServiceURL + "/ping")
			if err != nil {
				log.Printf("service: %v is unhealthy\n", registration.ServiceName)
				r.Log.Write(registration.ServiceName + " " + err.Error())
				continue
			}
			if resp.StatusCode == http.StatusOK {
				log.Printf("service: %v is OK\n", registration.ServiceName)
				r.Log.Write(registration.ServiceName + resp.Status)
			} else {
				log.Printf("service: %v is unhealthy\n", registration.ServiceName)
				r.Log.Write(registration.ServiceName + resp.Status)
			}
		}
		r.Mutex.RUnlock()
	}
}

func (r *Registry) Add(name, url string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	r.Registrations = append(r.Registrations, Registration{
		ServiceName: name,
		ServiceURL:  url,
	})
}
