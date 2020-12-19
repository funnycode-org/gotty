package sample

import (
	"bytes"
	"encoding/binary"
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

func (c *CustomizeProtocolBasedLengthField) Encode(srcObj interface{}) ([]byte, error) {
	cpblf := srcObj.(CustomizeProtocolBasedLengthField)
	myBytes := make([]byte, 1+1+8+len(cpblf.Body))
	buf := bytes.NewBuffer(myBytes)
	buf.WriteByte(cpblf.Type)
	buf.WriteByte(cpblf.Flag)
	binary.Write(buf, binary.LittleEndian, cpblf.Length)
	buf.WriteString(cpblf.Body)
	return buf.Bytes(), nil
}

func New(maxFrameLength uint) protocol.ProtocolDecoder {
	var cpblf CustomizeProtocolBasedLengthField
	return length_field.NewLengthFieldBasedFrame(
		length_field.WithMaxFrameLength(maxFrameLength),
		length_field.WithLengthFieldOffset(uint(unsafe.Sizeof(cpblf.Type)+unsafe.Sizeof(cpblf.Flag))),
		length_field.WithLengthFieldLength(uint32(unsafe.Sizeof(cpblf.Length))),
	)
}
