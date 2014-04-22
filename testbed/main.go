package main

import "github.com/wolfired/golabs/event"
import "fmt"

func main() {
	var e event.IEvent = MyEvent{"click"}
	var er event.IEventRouter = &MyRouter{}
	er.Router(e)
}

type MyEvent struct {
	_type string
}

func (this MyEvent) Type() string {
	return this._type
}

type MyRouter struct {
}

func (this MyRouter) Router(e event.IEvent) {
	fmt.Println(e.Type())
}
