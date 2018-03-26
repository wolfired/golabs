package www

import (
	"strconv"
	"fmt"
	"net/http"
	"os"

	"github.com/wolfired/golabs/namesilo"
)

func init() {
	Passwd = os.Getenv("AUTH_PASSWD")
	
	namesiloVersion, _ := strconv.Atoi(os.Getenv("NAMESILO_VERSION"))
	SiloCli = &namesilo.SiloClient{Version: uint(namesiloVersion), Type: os.Getenv("NAMESILO_TYPE"), Key: os.Getenv("NAMESILO_KEY")}
	
	hight, _ := strconv.Atoi(os.Getenv("GEN_NOES_HIGHT"))
	Hight = uint(hight)

	low, _ := strconv.Atoi(os.Getenv("GEN_NOES_LOW"))
	Low = uint(low)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "%s", "<a href='set_ip?passwd=&key=&domain=&host=' target='_blank'>记录IP</a><br/>")
	fmt.Fprintf(res, "%s", "<a href='get_ip?key=' target='_blank'>获取IP</a><br/>")
	fmt.Fprintf(res, "%s", "<a href='show_nts?key=' target='_blank'>下次更新IP的时间</a><br/>")
	fmt.Fprintf(res, "%s", "<a href='gen_noes?noes=' target='_blank'>GenNoes</a>")
}

/*
Serve Web服务
*/
func Serve() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set_ip", SetIP)
	http.HandleFunc("/get_ip", GetIP)
	http.HandleFunc("/show_nts", ShowNextTimestamp)
	http.HandleFunc("/generate_204", Generate204)
	http.HandleFunc("/gen_noes", GenNoes)

	bind := fmt.Sprintf("%s:%s", os.Getenv("WEB_HOST"), os.Getenv("WEB_PORT"))
	fmt.Printf("listening on %s...", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
}
