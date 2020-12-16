package base

type Listener interface {
	OnOpen() error
	OnClose() error
	OnSend(pks []byte) error
	OnReceive() ([]byte, error)
	OnError() error
}
