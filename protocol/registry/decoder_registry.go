package registry

import (
	"errors"
	"fmt"
	"github.com/funnycode-org/gotty/protocol"
	"sync"
)

var (
	registry map[string]protocol.ProtocolDecoder
	locker   sync.RWMutex
)

func init() {
	registry = make(map[string]protocol.ProtocolDecoder)
	locker = sync.RWMutex{}
}

func AddProtocol(name string, pd protocol.ProtocolDecoder) error {
	locker.Lock()
	defer locker.Unlock()
	if _, exist := registry[name]; exist {
		return errors.New(fmt.Sprintf("ProtocolDecoder %s had been registered!", name))
	}
	registry[name] = pd
	return nil
}

func FindProtocol(name string) (protocol.ProtocolDecoder, error) {
	locker.RLock()
	defer locker.Unlock()
	if pd, exist := registry[name]; !exist {
		return nil, errors.New(fmt.Sprintf("ProtocolDecoder %s hadn't  been registered!", name))
	} else {
		return pd, nil
	}
}
