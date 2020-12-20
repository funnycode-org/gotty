package listener

import (
	"github.com/funnycode-org/gotty/base"
)

type Listener interface {
	base.Listener
	GetRegistryListenerName() string
	FactoryConstruct() error
}
