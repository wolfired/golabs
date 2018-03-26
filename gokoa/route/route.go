package route

import  (
	"github.com/wolfired/golabs/gokoa"
	"net/http"
	"context"
)

func Get(path string, middleware gokoa.Middleware) gokoa.Middleware {
	return func(ctx context.Context, next gokoa.Next){
		req := ctx.Value(gokoa.KeyReq("req")).(*http.Request)
		// resp := ctx.Value(gokoa.KeyResp("resp")).(http.ResponseWriter)

		if http.MethodGet == req.Method && path == req.URL.Path {
			middleware(ctx, next)
		}else{
			next()
		}
	}
}

