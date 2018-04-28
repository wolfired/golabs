package route

import (
	"github.com/wolfired/golabs/middleware"
)

func Get(path string, h middleware.Middleware) middleware.Middleware {
	return func(s *middleware.Context) {
		s.Next()
	}
}
