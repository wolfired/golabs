package main

import (
	"fmt"
	"github.com/wolfired/golabs/swfchef"
)

func main() {
	testSwfchef()
}

func testSwfchef() {
	s := swfchef.ReadSwf("C:\\Users\\zelda\\Desktop\\res\\C.swf")

	fmt.Printf("%s\n", s.Signature())
	fmt.Printf("%d\n", s.Version())
	fmt.Printf("%d\n", s.Length())
	fmt.Println(s.FrameSize())
	fmt.Printf("%f\n", s.FrameRate())
	fmt.Println(s.FrameCount())

	// fmt.Println(swfchef.Raw2sbn([]byte{0x60, 0x00, 0x28, 0x00, 0x00, 0x28, 0x00}, 41, 12))
	// fmt.Println(swfchef.Raw2sbn([]byte{0xFF, 0xFE, 0xFD, 0xFC}, 8, 8))
}
