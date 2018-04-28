package static

import (
	"github.com/wolfired/golabs/middleware"
)

func Serve(root string, opts struct{}) middleware.Middleware {
	return func(s *middleware.Context) {
		s.Next()
	}
}
