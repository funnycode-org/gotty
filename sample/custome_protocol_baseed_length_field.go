package sample

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/funnycode-org/gotty/base"
	"github.com/funnycode-org/gotty/protocol"
	"github.com/funnycode-org/gotty/protocol/length_field"
	"github.com/funnycode-org/gotty/protocol/registry"
	"reflect"
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

func init() {
	registry.SetProtocol(reflect.TypeOf(CustomizeProtocolBasedLengthField{}), newProtocolBasedLengthField(base.GottyConfig.Server.MaxFrameLength))
}

func (c *CustomizeProtocolBasedLengthField) Decode(myBytes []byte) (decodeRrr error) {
	defer func() {
		if err := recover(); err != nil {
			decodeRrr = errors.New(fmt.Sprintf("解析到CustomizeProtocolBasedLengthField出错：%v", err))
			return
		}
	}()
	if len(myBytes) <= int(unsafe.Sizeof(c.Type)+unsafe.Sizeof(c.Flag)+unsafe.Sizeof(c.Length)) {
		decodeRrr = errors.New("长度不对，无法解析")
		return
	}
	c.Type = myBytes[0]
	c.Flag = myBytes[1]
	buf := bytes.NewBuffer(myBytes[2 : 2+8])
	binary.Read(buf, binary.LittleEndian, c.Length)
	c.Body = string(myBytes[2+8:])
	return
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

func newProtocolBasedLengthField(maxFrameLength int64) protocol.ProtocolDecoder {
	var cpblf CustomizeProtocolBasedLengthField
	return length_field.NewLengthFieldBasedFrame(
		length_field.WithMaxFrameLength(maxFrameLength),
		length_field.WithLengthFieldOffset(int32(unsafe.Sizeof(cpblf.Type)+unsafe.Sizeof(cpblf.Flag))),
		length_field.WithLengthFieldLength(int32(unsafe.Sizeof(cpblf.Length))),
	)
}
