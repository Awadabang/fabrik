package registry

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Awadabang/fabrik/internal/services/logservice"
)

type Registration struct {
	ServiceName      string
	ServiceURL       string
	ServiveAccessKey string
}

type Registry struct {
	Registrations sync.Map
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
		r.Registrations.Range(func(key, value any) bool {
			registration, ok := value.(Registration)
			if !ok {
				return true
			}
			log.Printf("AliveCheck: service name: %v, addr: %v\n", registration.ServiceName, registration.ServiceURL)

			req, err := http.NewRequest(http.MethodPost, registration.ServiceURL+"/ping", nil)
			if err != nil {
				log.Printf("service: %v NewRequest error\n", registration.ServiceName)
				r.Log.Write(registration.ServiceName + " " + err.Error())
				return true
			}

			req.Header.Set("accessKey", registration.ServiveAccessKey)

			resp, err := r.Client.Do(req)
			if err != nil {
				log.Printf("service: %v is unhealthy\n", registration.ServiceName)
				r.Log.Write(registration.ServiceName + " " + err.Error())
				return true
			}
			if resp.StatusCode == http.StatusOK {
				log.Printf("service: %v is OK\n", registration.ServiceName)
				r.Log.Write(registration.ServiceName + resp.Status)
			} else {
				log.Printf("service: %v is unhealthy\n", registration.ServiceName)
				r.Log.Write(registration.ServiceName + resp.Status)
			}
			return true
		})
	}
}

func (r *Registry) Add(name, url, accessKey string) {
	r.Registrations.Store(name, Registration{
		ServiceName:      name,
		ServiceURL:       url,
		ServiveAccessKey: accessKey,
	})
}

func (r *Registry) Delete(name string) {
	r.Registrations.Delete(name)
}
