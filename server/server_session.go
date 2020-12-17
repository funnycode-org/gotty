package server

type Session struct {
	//Connection net.Conn
	l Listener
}

func NewSession(l Listener) *Session {
	return &Session{
		l: l,
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
