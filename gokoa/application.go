package gokoa

import (
	"context"
	"net/http"
)

type KeyResp string
type KeyReq string

type Next func()
type Middleware func(ctx context.Context, next Next)

type Application struct {
	middleware []Middleware
}

func (a *Application)Use(middleware Middleware) (*Application) {
	a.middleware = append(a.middleware, middleware)
	return a
}

func (a *Application)Listen() {
	http.Handle("/", a)
	http.ListenAndServe(":8889", nil)
}

func (a *Application)ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := context.WithValue(context.Background(), KeyResp("resp"), resp)
	ctx = context.WithValue(ctx, KeyReq("req"), req)

	if(0 == len(a.middleware)){
		resp.WriteHeader(404)
		return
	}

	done := make(chan struct{})
	lock := make(chan chan struct{})

	go func(ctx context.Context){
		for i := 0; i < len(a.middleware); i++ {
			go func(ctx context.Context, i int, p chan struct{}, c chan struct{}) {
				next := func() {
					lock<-c
					<-c
				}

				a.middleware[i](ctx, next)

				p<-struct{}{}
			}(ctx, i, <-lock, make(chan struct{}))
		}

		<-lock<-struct{}{}
	}(ctx)

	lock<-done
	<-done
}
