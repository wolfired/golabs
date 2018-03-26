package www

import (
	"fmt"
	"github.com/wolfired/golabs/utils"
	"net/http"
	"strconv"
)

var (
	/*Hight 高位阀值*/
	Hight uint
	/*Low 低位阀值*/
	Low   uint
)

/*
GenNoes 生成数字串
*/
func GenNoes(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if nil != err {
		fmt.Fprintf(res, err.Error())
		return
	}

	noes := req.Form.Get("noes")

	kses := make([]uint, len(noes))
	var d int
	for i, v := range noes {
		d, _ = strconv.Atoi(string(v))
		kses[i] = uint(d)
	}

	fmt.Fprint(res, utils.Nogener(kses, Hight, Low))
}
