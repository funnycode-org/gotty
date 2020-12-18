package listener

import (
	"errors"
	"fmt"
	"sync"
)

type factory func() (Listener, error)

var registries = make(map[string]factory)

var lock sync.RWMutex

func RegisterListener(name string, factory factory) error {
	lock.Lock()
	defer lock.Unlock()
	registries[name] = factory
	if _, exist := registries[name]; exist {
		return errors.New(fmt.Sprintf("listener %s had been registered!", name))
	}
	registries[name] = factory
	return nil
}

func FindListener(name string) factory {
	lock.RLock()
	defer lock.Unlock()
	return registries[name]
}
