package length_field

import (
	"fmt"
	"testing"
	"unsafe"
)

const INT_SIZE = int(unsafe.Sizeof(0))

func TestSystemEndian(t *testing.T) {
	systemEndian()
}
func systemEndian() {
	var i int = 0x1234
	bs := (*[INT_SIZE]byte)(unsafe.Pointer(&i))
	fmt.Println(*bs)
	fmt.Println(&bs[0])
	fmt.Println(&bs[1])
	fmt.Println(&bs[2])
	fmt.Println("...")
	if bs[0] == 52 {
		fmt.Println("system endian is little endian")
	} else {
		fmt.Println("system endian is big endian")
	}
}
