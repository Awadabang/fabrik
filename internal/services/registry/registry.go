package registry

import (
	"sync"
)

type Registry struct {
	ServiceNode sync.Map
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

func (r *Registry) Add(key, value any) {
	r.ServiceNode.Store(key, value)
}
