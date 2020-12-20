package client

import (
	"reflect"
)

type WrappedSession struct {
	session *Session
}

func (w *WrappedSession) GetWrappedSession() interface{} {
	panic("can't be called!")
}

func newWrappedSession(session *Session) *WrappedSession {
	return &WrappedSession{
		session: session,
	}
}

func (w *WrappedSession) Close() error {
	return w.session.Close()
}

func (w *WrappedSession) SessionId() int {
	return w.session.SessionId()
}

func (w *WrappedSession) Send(bytes []byte) error {
	return w.session.Send(bytes)
}

func (w *WrappedSession) GetRegistryProtocol() reflect.Type {
	return w.session.GetRegistryProtocol()
}

func (w *WrappedSession) GetSendChannel() <-chan []byte {
	panic("can't be called!")
}
