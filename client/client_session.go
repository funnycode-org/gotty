package client

import (
	"github.com/funnycode-org/gotty/client/lisener"
	"reflect"
	"sync/atomic"
)

var number int64

type Session struct {
	Using         bool
	sessionId     int64
	l             listener.Listener
	receivedBytes []byte
}

func (s *Session) Send(bytes []byte) error {
	panic("implement me")
}

func (s *Session) GetRegistryProtocol() reflect.Type {
	panic("implement me")
}

func (s *Session) GetSendChannel() <-chan []byte {
	panic("implement me")
}

func NewSession(l listener.Listener) *Session {
	return &Session{
		sessionId:     atomic.AddInt64(&number, 1),
		l:             l,
		receivedBytes: make([]byte, 1024),
	}
}

func (s Session) Close() error {
	panic("implement me")
}

func (s Session) Open() error {
	panic("implement me")
}

func (s Session) SessionId() int {
	return 0
}
