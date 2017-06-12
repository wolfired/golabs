package swfchef

import (
	"fmt"
	"testing"
)

func TestSwfRead(t *testing.T) {
	s := SwfRead("C:\\Users\\Administrator\\Desktop\\res\\C.swf")

	fmt.Printf("%s\n", s.Signature())
	fmt.Printf("%d\n", s.Version())
	fmt.Printf("%d\n", s.Length())
	fmt.Println(s.FrameSize())
	fmt.Printf("%f\n", s.FrameRate())
	fmt.Println(s.FrameCount())

	fmt.Println(raw2sbn([]byte{0x60, 0x00, 0x28, 0x00, 0x00, 0x28, 0x00}, 41, 12))
	fmt.Println(raw2sbn([]byte{0xFF, 0xFE, 0xFD, 0xFC}, 8, 8))
}
