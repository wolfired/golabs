package openshift

import (
	"fmt"
	"github.com/wolfired/golabs/utils"
	"net/http"
	"strconv"
)

var (
	Hight uint
	Low   uint
)

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
