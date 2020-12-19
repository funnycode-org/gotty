package server

import (
	"github.com/funnycode-org/gotty/protocol"
	"github.com/funnycode-org/gotty/server/listener"
)

type Session struct {
	//Connection net.Conn
	l             listener.Listener
	receivedBytes []byte
	protocol      protocol.ProtocolDecoder
}

func (s *Session) Send() error {
	panic("implement me")
}

func NewSession(l listener.Listener, protocol protocol.ProtocolDecoder) *Session {
	return &Session{
		l:             l,
		receivedBytes: make([]byte, 1024),
		protocol:      protocol,
	}
}

func (s *Session) Close() error {
	s.l = nil
	s.receivedBytes = nil
	return nil
}

func (s Session) Open() error {
	panic("implement me")
}

func (s Session) SessionId() int {
	return 0
}
