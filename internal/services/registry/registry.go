package registry

import "sync"

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

func (r *Registry) Start() {

}

func (r *Registry) Add(name, url string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	r.Registrations = append(r.Registrations, Registration{
		ServiceName: name,
		ServiceURL:  url,
	})
}
