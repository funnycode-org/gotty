package base

type Session interface {
	Close() error
	SessionId() int
	Send(bytes []byte) error
}
