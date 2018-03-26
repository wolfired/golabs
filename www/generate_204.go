package www

import (
	"net/http"
)

/*
Generate204 生成一个空内容的页面，用于Google手机的在线验证
*/
func Generate204(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNoContent)
}
