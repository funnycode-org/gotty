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
	session base.Session
}

func (d *DefaultListener) FactoryConstruct() error {
	panic("implement me")
}

func (d *DefaultListener) OnOpen() error {
	fmt.Printf("session %d is opened!", d.session.SessionId())
	return nil
}

func (d *DefaultListener) OnClose() error {
	fmt.Printf("session %d is closed!", d.session.SessionId())
	return nil
}

func (d *DefaultListener) OnSend(pks []byte) error {
	fmt.Printf("session %d is sending bytes: %x!", d.session.SessionId(), pks)
	return nil
}

func (d *DefaultListener) OnReceive() ([]byte, error) {
	panic("implement me")
}

func (d *DefaultListener) OnError() error {
	panic("implement me")
}

func (d *DefaultListener) GetRegistryListenerName() string {
	panic("implement me")
}
