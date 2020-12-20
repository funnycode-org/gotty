package base

import "reflect"

type Session interface {
	Close() error
	SessionId() int
	Send(bytes []byte) error
	GetRegistryProtocol() reflect.Type
	GetSendChannel() <-chan []byte
}
