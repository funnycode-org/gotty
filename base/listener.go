package base

type Listener interface {
	OnOpen(session Session) error
	OnClose(session Session) error
	OnSend(session Session) error
	OnReceive(session Session, bytes []byte) ([]byte, error)
	OnError(session Session) error
}
