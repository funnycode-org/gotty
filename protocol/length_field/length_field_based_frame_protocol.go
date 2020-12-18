package length_field

import (
	"io"
)

type LengthFieldBasedFrame struct {
	MaxFrameLength    uint // 发送的数据包最大长度
	LengthFieldOffset uint // 长度域偏移量
	LengthFieldLength uint // 长度域的自己的字节数长度
	// 暂时不考虑下面的场景
	//LengthAdjustment    uint // 长度域的偏移量矫正
	//InitialBytesToStrip uint // 丢弃的起始字节数
}

type Option func(*LengthFieldBasedFrame)

func NewLengthFieldBasedFrame(options ...Option) *LengthFieldBasedFrame {
	var lfbf LengthFieldBasedFrame
	for _, option := range options {
		option(&lfbf)
	}
	return &lfbf
}

func WithMaxFrameLength(maxFrameLength uint) Option {
	return func(lfbf *LengthFieldBasedFrame) {
		lfbf.MaxFrameLength = maxFrameLength
	}
}

func WithLengthFieldOffset(lengthFieldOffset uint) Option {
	return func(lfbf *LengthFieldBasedFrame) {
		lfbf.LengthFieldOffset = lengthFieldOffset
	}
}

func WithLengthFieldLength(lengthFieldLength uint) Option {
	return func(lfbf *LengthFieldBasedFrame) {
		lfbf.LengthFieldLength = lengthFieldLength
	}
}
func WithLengthAdjustment(lengthAdjustment uint) Option {
	return func(lfbf *LengthFieldBasedFrame) {
		lfbf.LengthAdjustment = lengthAdjustment
	}
}

func WithInitialBytesToStrip(initialBytesToStrip uint) Option {
	return func(lfbf *LengthFieldBasedFrame) {
		lfbf.InitialBytesToStrip = initialBytesToStrip
	}
}

func (l *LengthFieldBasedFrame) Decode(reader io.Reader, targetObj interface{}) error {
	panic("implement me")
}

func (l *LengthFieldBasedFrame) Encode(srcObj interface{}) ([]byte, error) {
	panic("implement me")
}
