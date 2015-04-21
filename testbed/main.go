package main

import "archive/zip"

import "bytes"
import "os"

import "fmt"

// import "github.com/wolfired/golabs/server"
// import "time"

func main() {
	b := bytes.Buffer{}

	zip_file_buf := zip.NewWriter(&b)

	zip_item, err := zip_file_buf.Create("phone.txt")

	if nil == err {
		zip_item.Write([]byte("18601011241"))
		zip_file_buf.Close()

		zip_file, err := os.Create("C:/Users/zelda/Desktop/res/phones.zip")
		if nil == err {
			zip_file.Write(b.Bytes())
			zip_file.Close()
		} else {
			fmt.Println(err.Error())
		}

	} else {
		fmt.Println(err.Error())
	}

	// gs := server.GateServer{server.Server{"tcp", ":9090"}}
	// go gs.Run()

	// time.Sleep(60 * time.Second)

	// fmt.Println("exit")
}
