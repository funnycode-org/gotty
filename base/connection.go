package base

import "net"

type Connection struct {
	con     net.Conn
	session Session
	//sessionMap
	close chan struct{}
}

//type sessionMap struct {
//	sessionIds map[int]struct{}
//	sessions   []Session
//}

func NewConnection(con net.Conn, sessionCount uint) *Connection {
	return &Connection{
		con: con,
		//sessionMap: sessionMap{
		//	sessionIds: make(map[int]struct{}, sessionCount),
		//	sessions:   make([]Session, sessionCount),
		//},
	}
}

func (con *Connection) Do() {

	con.session = NewSession()
	for {
		select {
		case <-con.close:
			//todo
			break
		}
	}
}
