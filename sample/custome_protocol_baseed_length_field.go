package sample

import (
	"github.com/funnycode-org/gotty/protocol"
	"github.com/funnycode-org/gotty/protocol/length_field"
	"unsafe"
)

// CustomizeProtocolBasedLengthField 自定义了一个编解码协议
type CustomizeProtocolBasedLengthField struct {
	// 	类型数值以后扩展1,2,3......
	Type byte
	// 信息标志  1:表示心跳包 2:业务信息包
	Flag byte
	// 请求消息内容长度
	Length uint64
	// 请求消息体
	Body string
}

func New(maxFrameLength uint) protocol.Protocol {
	var cpblf CustomizeProtocolBasedLengthField
	return length_field.NewLengthFieldBasedFrame(
		length_field.WithMaxFrameLength(maxFrameLength),
		length_field.WithLengthFieldOffset(uint(unsafe.Sizeof(cpblf.Type)+unsafe.Sizeof(cpblf.Flag))),
		length_field.WithLengthFieldLength(uint(unsafe.Sizeof(cpblf.Length))),
	)
}
