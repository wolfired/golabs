package logic

import (
	"fmt"
)

type Handler func()

var Saver [128]Handler

func init() {
	Saver[0] = auth
}

func auth() {
	fmt.Println("auth")
}
