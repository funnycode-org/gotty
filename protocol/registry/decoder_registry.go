package registry

import (
	"errors"
	"fmt"
	"github.com/funnycode-org/gotty/protocol"
	"reflect"
	"sync"
)

var (
	registryProtocolKind protocol.ProtocolDecoder
	registryProtocol     reflect.Type
	locker               sync.RWMutex
)

func init() {
	locker = sync.RWMutex{}
}

func SetProtocol(rp reflect.Type, pd protocol.ProtocolDecoder) error {
	locker.Lock()
	defer locker.Unlock()
	if registryProtocolKind != nil {
		return errors.New(fmt.Sprintf("ProtocolDecoder %s had been registered!", registryProtocol.Name()))
	}
	registryProtocolKind = pd
	registryProtocol = rp
	return nil
}

func GetProtocol() (protocol.ProtocolDecoder, reflect.Type, error) {
	locker.RLock()
	defer locker.Unlock()
	if registryProtocolKind != nil {
		return nil, nil, errors.New("没有注册过协议！")
	} else {
		return registryProtocolKind, registryProtocol, nil
	}
}
