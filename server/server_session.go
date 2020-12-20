package server

import (
	"errors"
	"github.com/funnycode-org/gotty/base"
	"github.com/funnycode-org/gotty/protocol"
	"github.com/funnycode-org/gotty/server/listener"
	"reflect"
	"sync/atomic"
	"time"
)

var number int64

type Session struct {
	l                listener.Listener
	receivedBytes    []byte
	protocol         protocol.ProtocolDecoder
	registryProtocol reflect.Type
	sessionId        int64
	send             chan []byte
}

func (s *Session) Send(bytes []byte) (err error) {
	timeout := time.After(time.Millisecond * time.Duration(base.GottyConfig.Server.SendTimeout))
	select {
	case <-timeout:
		err = errors.New("发送超时")
		break
	case s.send <- bytes:
		break
	}
	return
}

func NewSession(l listener.Listener) *Session {
	return &Session{
		sessionId:     atomic.AddInt64(&number, 1),
		l:             l,
		receivedBytes: make([]byte, 1024),
	}
}

func (s *Session) Close() error {
	s.l = nil
	s.receivedBytes = nil
	return nil
}

func (s Session) SessionId() int {
	return 0
}
