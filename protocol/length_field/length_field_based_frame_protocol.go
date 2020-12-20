package length_field

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/funnycode-org/gotty/util"
	"log"
)

type LengthFieldBasedFrame struct {
	MaxFrameLength    int64 // 发送的数据包最大长度
	LengthFieldOffset int32 // 长度域偏移量
	LengthFieldLength int32 // 长度域的自己的字节数长度
	// 暂时不考虑下面的场景
	//LengthAdjustment    uint // 长度域的偏移量矫正
	//InitialBytesToStrip uint // 丢弃的起始字节数
	skipBytes              int64
	leftContentFrameLength int64
	readingPkgs            bool
}

const LengthFieldBasedFrameName = "length_field"

func init() {
}

type Option func(*LengthFieldBasedFrame)

func NewLengthFieldBasedFrame(options ...Option) *LengthFieldBasedFrame {
	var lfbf LengthFieldBasedFrame
	for _, option := range options {
		option(&lfbf)
	}
	return &lfbf
}

func WithMaxFrameLength(maxFrameLength int64) Option {
	if maxFrameLength < 1 {
		log.Panicf("无效的maxFrameLength:%d", maxFrameLength)
	}
	return func(lfbf *LengthFieldBasedFrame) {
		lfbf.MaxFrameLength = maxFrameLength
	}
}

func WithLengthFieldOffset(lengthFieldOffset int32) Option {
	return func(lfbf *LengthFieldBasedFrame) {
		lfbf.LengthFieldOffset = lengthFieldOffset
	}
}

func WithLengthFieldLength(lengthFieldLength int32) Option {
	return func(lfbf *LengthFieldBasedFrame) {
		lfbf.LengthFieldLength = lengthFieldLength
	}
}

//func WithLengthAdjustment(lengthAdjustment uint) Option {
//	return func(lfbf *LengthFieldBasedFrame) {
//		lfbf.LengthAdjustment = lengthAdjustment
//	}
//}
//
//func WithInitialBytesToStrip(initialBytesToStrip uint) Option {
//	return func(lfbf *LengthFieldBasedFrame) {
//		lfbf.InitialBytesToStrip = initialBytesToStrip
//	}
//}

func (l *LengthFieldBasedFrame) Decode(reader *bytes.Buffer, writer *bytes.Buffer) (b bool, err error) {

	for ; l.skipBytes > 0; l.skipBytes-- {
		err := reader.UnreadByte()
		if err != nil {
			log.Println("目前reader的包已经被丢弃完了")
			l.skipBytes++
			break
		}
	}
	if l.skipBytes > 0 {
		log.Println("目前reader还没被丢弃完")
		return
	}

	// 需要丢弃的字节已经被丢弃完了
	// 接下来读取去读字节
	for ; l.leftContentFrameLength > 0; l.leftContentFrameLength-- {
		readByte, err := reader.ReadByte()
		if err != nil {
			log.Println("目前reader的包已经被丢弃完了")
			l.leftContentFrameLength++
			break
		}
		writer.WriteByte(readByte)
	}
	if l.leftContentFrameLength < 1 && l.readingPkgs {
		log.Println("目前已经读取完了整个包")
		b = true
		l.readingPkgs = false
		return
	}
	if l.leftContentFrameLength > 0 && l.readingPkgs {
		log.Println("继续读吧")
		return
	}

	// 读取下个包
	if uint64(reader.Len()) < uint64(l.LengthFieldOffset)+uint64(l.LengthFieldLength) {
		log.Println("字节长度不能读到长度")
		return
	}

	contentFrameLength := l.getContentFrameLength(reader)
	totalFrameLength := contentFrameLength + int64(l.LengthFieldOffset) + int64(l.LengthFieldLength)

	if totalFrameLength > l.MaxFrameLength {
		l.skipBytes = totalFrameLength
		err = errors.New("该包太长了，直接丢弃了")
		return
	}

	l.leftContentFrameLength = contentFrameLength
	// 把长度和长度前面的字节写到writer里面去
	writer.Write(reader.Next(int(l.LengthFieldOffset) + int(l.LengthFieldLength)))
	l.readingPkgs = true
	return
}

func (l *LengthFieldBasedFrame) getContentFrameLength(reader *bytes.Buffer) int64 {
	reader.Next(int(l.LengthFieldOffset))
	contentLengthBytes := reader.Next(int(l.LengthFieldLength))

	bytesBuffer := bytes.NewBuffer(contentLengthBytes)
	var contentLength int64
	binary.Read(bytesBuffer, binary.BigEndian, &contentLength)
	//
	util.UnReadBuffer(reader, uint64(l.LengthFieldOffset)+uint64(l.LengthFieldLength))
	return contentLength
}
