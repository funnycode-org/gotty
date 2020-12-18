package server

import "github.com/funnycode-org/gotty/server/listener"

type Session struct {
	//Connection net.Conn
	l             listener.Listener
	receivedBytes []byte
}

func NewSession(l listener.Listener) *Session {
	return &Session{
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
