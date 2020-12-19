package protocol

import (
	"bytes"
)

type ProtocolDecoder interface {
	Decode(reader bytes.Buffer, writer bytes.Buffer) (bool, error)
}

type ProtocolEncoder interface {
	Encode(srcObj interface{}) ([]byte, error)
}
