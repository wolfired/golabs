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

	// fmt.Print(s)
	fmt.Printf("%s\n", s.Signature())
	fmt.Printf("%d\n", s.Version())
	fmt.Printf("%d\n", s.Length())

}
