package sample

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestUintptr(t *testing.T) {
	var cpblf CustomizeProtocolBasedLengthField
	sizeof := unsafe.Sizeof(cpblf.Body)
	fmt.Println(uint(sizeof))
}
