package length_field

import "github.com/funnycode-org/gotty/protocol"

type LengthFieldBasedFrame struct {
	protocol.Protocol
	MaxFrameLength      uint // 发送的数据包最大长度
	LengthFieldOffset   uint // 长度域偏移量
	LengthFieldLength   uint // 长度域的自己的字节数长度
	LengthAdjustment    uint // 长度域的偏移量矫正
	InitialBytesToStrip uint // 丢弃的起始字节数
}
