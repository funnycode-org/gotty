package listener

import (
	"fmt"
	"github.com/funnycode-org/gotty/base"
)

const DefaultListenerName = "default"

func init() {
	RegisterListener(DefaultListenerName, func() (Listener, error) {
		return &DefaultListener{}, nil
	})
}

type DefaultListener struct {
}

func (d *DefaultListener) FactoryConstruct() error {
	panic("implement me")
}

// 用户在这里能得到session，然后去发送数据
func (d *DefaultListener) OnOpen(session base.Session) error {
	fmt.Printf("session %d is opened!", session.SessionId())
	return nil
}

func (d *DefaultListener) OnClose(session base.Session) error {
	fmt.Printf("session %d is closed!", session.SessionId())
	return nil
}

func (d *DefaultListener) OnSend(session base.Session) error {
	fmt.Printf("session %d is sending bytes: %x!", session.SessionId(), session)
	return nil
}

func (d *DefaultListener) OnReceive(session base.Session, bytes []byte) ([]byte, error) {
	panic("implement me")
}

func (d *DefaultListener) OnError(session base.Session) error {
	panic("implement me")
}

func (d *DefaultListener) GetRegistryListenerName() string {
	return DefaultListenerName
}
