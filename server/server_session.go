package server

import "net"

type Session struct {
	Connection net.Conn
}

func (s Session) Close() error {
	panic("implement me")
}

func (s Session) Open() error {
	panic("implement me")
}
