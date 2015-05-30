package main

import (
	"fmt"
	"github.com/wolfired/golabs/swfchef"
)

func main() {
	testSwfchef()
}

func testSwfchef() {
	s := swfchef.ReadSwf("C:\\Users\\zelda\\Desktop\\res\\Z.swf")

	fmt.Printf("%d", s.Size())
}
