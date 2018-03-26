package static

import  (
	"fmt"
	"github.com/wolfired/golabs/gokoa"
	"github.com/wolfired/golabs/gokoa/send"
	"context"
	"net/http"
)

func Serve(root string, opts struct{}) gokoa.Middleware {
	return func(ctx context.Context, next gokoa.Next){
		fmt.Println("static file serve")

		req := ctx.Value(gokoa.KeyReq("req")).(*http.Request)
		// resp := ctx.Value(gokoa.KeyResp("resp")).(http.ResponseWriter)
		send.Send(ctx, root + req.URL.Path, opts)

		next()
	}
}