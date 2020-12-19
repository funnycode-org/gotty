package protocol

import (
	"bytes"
)

type Protocol interface {
	Decode(reader bytes.Buffer, writer bytes.Buffer) (bool, error)
	Encode(srcObj interface{}) ([]byte, error)
}
