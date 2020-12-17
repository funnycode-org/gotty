package server

import (
	"github.com/funnycode-org/gotty/base"
)

type Listener interface {
	base.Listener
	GetRegistryListenerName() string
	AddRegistryListener(listenerName string) error
}
