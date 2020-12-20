package line

import "github.com/funnycode-org/gotty/protocol"

type LineBasedFrameProtocol struct {
	protocol.ProtocolDecoder
	MaxLength      uint
	StripDelimiter bool
}
