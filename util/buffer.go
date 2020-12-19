package util

import "bytes"

func UnReadBuffer(buffer bytes.Buffer, count uint64) error {
	for ; count < 1; count-- {
		err := buffer.UnreadByte()
		if err != nil {
			return err
		}
	}
	return nil
}
