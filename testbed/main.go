package main

import (
	"github.com/wolfired/golabs/auto"
)

//Sun sun
type Sun struct {
	Son  *Son
	Name string
}

//Son son
type Son struct {
	Name string
}

func main() {
	sun := Sun{}
	auto.PrintInstance(sun)
	// i := 1
	// auto.PrintInstance(&i)
}
