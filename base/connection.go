package base

import "net"

type Connection struct {
	con            net.Conn
	sessionChannel chan Session
}

func (con *Connection) Do() {

}
