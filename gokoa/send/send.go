package send

import (
	"os"
	"github.com/wolfired/golabs/gokoa"
	"context"
	"net/http"
)

func Send(ctx context.Context, path string, opts struct{}){
	// req := ctx.Value(gokoa.KeyReq("req")).(*http.Request)
	resp := ctx.Value(gokoa.KeyResp("resp")).(http.ResponseWriter)

	f, _ := os.Open(path)
	defer f.Close()

	bs := make([]byte, 1024)
	f.Read(bs)
	resp.Write(bs)
}