package base

type Package interface {
	Send() error
	Read() error
}
