package main

import (
	"fmt"
)

type IBook interface {
	read()
}

type Book struct {
	byte
}

func (b Book) read() {
	fmt.Printf("%p\n", &b)
}

func (b *Book) name() {
	fmt.Printf("%p\n", b)
}

func test(ib IBook) {

}

func main() {
	var b Book
	fmt.Printf("%p\n", &b)

	b.read()
	(&b).name()

	fmt.Println("----------")

	var bp *Book = new(Book)
	fmt.Printf("%p\n", bp)

	(*bp).read()
	bp.name()

	fmt.Println("----------")

	var ib IBook

	ib = b
	fmt.Printf("%p\n", &ib)
	ib.read()

	fmt.Println("----------")

	ib = bp
	fmt.Printf("%p\n", ib)
	ib.read()

	var ch chan int
	ch <- 1
	var i int = <-ch
}
