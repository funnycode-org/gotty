package fixed_length

import "github.com/funnycode-org/gotty/protocol"

type FixedLengthFrameProtocol struct {
	protocol.Protocol
	FixedLength uint // 固定长度
}
