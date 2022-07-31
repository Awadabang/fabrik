package registry

import (
	"log"
	"sync"
	"time"
)

type Registration struct {
	ServiceName string
	ServiceURL  string
}

type Registry struct {
	Registrations []Registration

	Mutex sync.RWMutex
}

type RegistryOption func(*Registry)

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
