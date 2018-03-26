package main

import (
	"github.com/wolfired/golabs/www"
	"net/http"
	"time"
	"context"
	"fmt"
	"math/rand"
	"github.com/wolfired/golabs/gokoa"
	"github.com/wolfired/golabs/gokoa/route"
	"github.com/wolfired/golabs/gokoa/static"
	"runtime/pprof"
	"os"
)

func main() {
	www.Serve()
}

func test01(){
	koa := new(gokoa.Application)

	koa.Use(route.Get("/index", func(ctx context.Context, next gokoa.Next){
		resp := ctx.Value(gokoa.KeyResp("resp")).(http.ResponseWriter)
		
		resp.Write([]byte("Index in\n"))
		next()
		resp.Write([]byte("Index out\n"))
	}))

	koa.Use(route.Get("/about", func(ctx context.Context, next gokoa.Next){
		resp := ctx.Value(gokoa.KeyResp("resp")).(http.ResponseWriter)
		
		resp.Write([]byte("About in\n"))
		next()
		resp.Write([]byte("About out\n"))
	}))

	koa.Use(func(ctx context.Context, next gokoa.Next){
		resp := ctx.Value(gokoa.KeyResp("resp")).(http.ResponseWriter)
		
		resp.Write([]byte("Hello gokoa A\n"))
		next()
		resp.Write([]byte("Hello gokoa A\n"))
	})

	koa.Use(static.Serve("C:\\", struct{}{}))

	koa.Listen()
}

func test0() {
	koa := new(gokoa.Application)

	koa.Use(func(ctx context.Context, next gokoa.Next){
		resp := ctx.Value(gokoa.KeyResp("resp")).(http.ResponseWriter)
		
		fmt.Println(">>>>A")
		resp.Write([]byte("Hello gokoa A\n"))
		next()
		resp.Write([]byte("Hello gokoa A\n"))
		fmt.Println("<<<<A")
	})

	koa.Use(func(ctx context.Context, next gokoa.Next){
		resp := ctx.Value(gokoa.KeyResp("resp")).(http.ResponseWriter)

		fmt.Println(">>>>B")
		resp.Write([]byte("Hello gokoa B\n"))
		next()
		resp.Write([]byte("Hello gokoa B\n"))
		fmt.Println("<<<<B")
	})

	koa.Use(func(ctx context.Context, next gokoa.Next){
		resp := ctx.Value(gokoa.KeyResp("resp")).(http.ResponseWriter)

		fmt.Println(">>>>C")
		resp.Write([]byte("Hello gokoa C\n"))
		next()
		resp.Write([]byte("Hello gokoa C\n"))
		fmt.Println("<<<<C")
	})

	koa.Listen()
}

func test(){
	out := make(chan struct{})


	next := func(){
		out<-struct{}{}
		<-make(chan struct{})
	}

	makefn := func(x int) (func()){
		return func(){
			fmt.Println(">>>", x)
			if x != 24{
				next()
			}
			fmt.Println("<<<", x)
		}
	}

	fns := [32](func()){}

	for index := 0; index < len(fns); index++ {
		fns[index] = makefn(index)
	}

	go func(){
		for _, fn := range fns {
			go fn()
			<-out
		}
	}()



	time.Sleep(3 * time.Second)
}

func test2(){
	fn0 := func(x int, next func()){
		fmt.Println(">>>", x)
		next()
		fmt.Println("<<<", x)
	}

	fn1 := func(x int, next func()){
		fmt.Println(">>>", x)
		next()
		fmt.Println("<<<", x)
	}

	fn2 := func(x int, next func()){
		fmt.Println(">>>", x)
		next()
		fmt.Println("<<<", x)
	}

	fnx := [](func(x int, next func())){fn0, fn1, fn2}

	handle := func(){
		pre := make(chan struct{})
		cur := make(chan struct{})
		nxt := make(chan struct{})

		for i, fn := range fnx {
			go func(p chan struct{}, c chan struct{}) {
				next := func() {
					nxt<-struct{}{}
					<-c
				}

				fn(i, next)

				p<-struct{}{}
			}(pre, cur)

			<-nxt

			pre = cur
			cur = make(chan struct{})
		}
	}

	go handle()

	time.Sleep(3 * time.Second)
}

func test3(){
	makefn := func() (func(int, func())){
		return func(x int, next func()){
			time.Sleep(time.Duration(rand.Int63n(128)) * time.Millisecond)
			fmt.Println(">>>", x)
			if 0 != x {
				next()
			}
			fmt.Println("<<<", x)
			time.Sleep(time.Duration(rand.Int63n(128)) * time.Millisecond)
		}
	}

	fnx := [2](func(x int, next func())){}

	for i := 0; i < len(fnx); i++ {
		fnx[i] = makefn()
	}

	done := make(chan struct{})
	lock := make(chan chan struct{})

	go func (){
		for i := 0; i < len(fnx); i++ {

			go func(i int, p chan struct{}, c chan struct{}) {
				next := func() {
					lock<-c
					<-c
				}

				fnx[i](i, next)

				p<-struct{}{}
			}(i, <-lock, make(chan struct{}))

		}
		
		// <-lock<-struct{}{}
	}()

	lock<-done
	<-done
}

func test40(){
	f, _ := os.Create("E:\\workspace_go\\src\\github.com\\wolfired\\golabs\\testbed\\log.prof")
	defer f.Close()
	
	test41()

	p := pprof.Lookup("goroutine")
	p.WriteTo(f, 1)
}

func test41() {
	s := makeSession()

	go func (s *session){
		for i := 0; i < 4; i++ {
			go func(s *session, i int) {
				//->检查运行权
				select {
				case key := <-s.downKey:
					if key == i {
						break
					}
				}

				next := func(key int){
					s.downKey<-key
					select {
					case k := <-s.upKey:
						if k == key {
							break
						}
					}
				}

				func() { //工作函数
					//具体工作
					fmt.Println(">>>", i)

					next(i + 1)

					//具体工作
					fmt.Println("<<<", i)
				}()

				//<-交出运行权
				s.upKey<-i-1
			}(s, i)

		}
		
		// <-lock<-struct{}{}
	}(s)

	s.downKey<-0

	time.Sleep(6 * time.Second)
}

type session struct {
	downKey chan int
	upKey chan int
}

func makeSession()*session{
	return &session{make(chan int), make(chan int)}
}

// func (s *session)next(key int){
// 	s.downKey<-key
// 	select {
// 	case k := <-s.upKey:
// 		if k == key {
// 			break
// 		}
// 	}
// }

// func (s *session)down(key int){
// 	fmt.Println("down", key)
	// select {
	// case k := <-s.downKey:
	// 	if k == key {
	// 		break
	// 	}
	// }
// }

// func (s *session)up(key int){
// 	select {
// 	case k := <-s.upKey:
// 		if k == key {
// 			break
// 		}
// 	}
// }

