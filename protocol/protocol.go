package protocol

import "io"

type Protocol interface {
	Decode(reader io.Reader, targetObj interface{}) error
	Encode(srcObj interface{}) ([]byte, error)
}
