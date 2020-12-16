package line

import "github.com/funnycode-org/gotty/protocol"

type LineBasedFrameProtocol struct {
	protocol.Protocol
	MaxLength      uint
	StripDelimiter bool
}
