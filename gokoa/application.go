package gokoa

import (
	"net/http"

	"github.com/wolfired/golabs/middleware"
)

/*
Application 应用
*/
type Application struct {
	middleware.Middleware
}

func (a *Application) Listen() {
	http.Handle("/index", a)
	http.ListenAndServe(":8889", nil)
}

func (a *Application) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

}

type Session struct {
	Req  *http.Request
	Resp http.ResponseWriter
}

func (s *Session) Next() {

}
