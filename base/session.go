package base

type Session interface {
	Close() error
	Open() error
}
