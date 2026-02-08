package registry

import (
	"log"
	"sync"
)

type Registry struct {
	services map[string]string
	snc      sync.RWMutex
}

func NewRegistry() *Registry {
	return &Registry{
		services: make(map[string]string),
	}
}

func (r *Registry) Lookup(service string) (string, bool) {
	r.snc.RLock()
	defer r.snc.RUnlock()

	addrIP, ok := r.services[service]
	if !ok {
		log.Printf(" SERVICIO no encontrado: %s\n", service)
		return "", false
	}
	return addrIP, ok
}

func (r *Registry) Register(service, addrIP string) {
	r.snc.Lock()
	defer r.snc.Unlock()

	r.services[service] = addrIP
}
