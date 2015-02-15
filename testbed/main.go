package main

import "archive/zip"

import "bytes"
import "os"

// import "unsafe"

import "fmt"

func exchange(x, y int, more ...string) (int, int) {
	for _, v := range more {
		fmt.Println(v)
	}
	return y, x
}

func main() {
	fmt.Println(exchange(1, 2, "3", "4"))

	b := bytes.Buffer{}

	zip_file_buf := zip.NewWriter(&b)

	zip_item, err := zip_file_buf.Create("phone.txt")

	if nil == err {
		zip_item.Write([]byte("18601011241"))
		zip_file_buf.Close()

		zip_file, err := os.Create("D:/phones.zip")
		if nil == err {
			zip_file.Write(b.Bytes())
			zip_file.Close()
		} else {
			fmt.Println(err.Error())
		}

	} else {
		fmt.Println(err.Error())
	}
}
