package main

import (
	"os"

	"github.com/wolfired/golabs/go2v8"
	"github.com/wolfired/golabs/namesilo"
	"github.com/wolfired/golabs/openshift"
)

func init() {
	openshift.Passwd = os.Getenv("PASSWD")
	openshift.SiloCli = &namesilo.SiloClient{1, "xml", os.Getenv("NS_TOKEN")}
	openshift.Hight = 5
	openshift.Low = 1
}

func main() {
	// http.HandleFunc("/", openshift.Index)
	// http.HandleFunc("/set_ip", openshift.SetIP)
	// http.HandleFunc("/get_ip", openshift.GetIP)
	// http.HandleFunc("/show_nts", openshift.ShowNextTimestamp)
	// http.HandleFunc("/generate_204", openshift.Generate204)
	// http.HandleFunc("/gen_noes", openshift.GenNoes)

	// bind := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	// fmt.Printf("listening on %s...", bind)
	// err := http.ListenAndServe(bind, nil)
	// if err != nil {
	// 	panic(err)
	// }

	go2v8.Test()
}
