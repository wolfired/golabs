package main

import (
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	help  bool
	mode  string
	ifile string
	ofile string
)

func main() {
	flag.BoolVar(&help, "help", false, "帮助")
	flag.StringVar(&mode, "mode", "decode", "处理模式, decode=解码, encode=编码")
	flag.StringVar(&ifile, "ifile", "", "输入数据文件路径, 不填测使用标准输入, 文件使用简单文本格式, 如: BD05BD05")
	flag.StringVar(&ofile, "ofile", "", "输出数据文件路径, 不填测使用标准输出, 文本使用简单文本格式, 如: 701 701")
	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	var buf []byte

	if "encode" == mode {
		if "" == ifile {
			buf = readFromStdin(make([]byte, 1024*1024*4))
		} else {
			buf, _ = ioutil.ReadFile(ifile)
		}
		values := spaceCutter(buf)

		buf = make([]byte, 1024*1024*4)
		n := 0
		for _, value := range values {
			n = n + encode(buf[n:], value)
		}
		if "" == ofile {
			fmt.Print(hexEncode(buf[:n]))
		} else {
			ioutil.WriteFile(ofile, []byte(hexEncode(buf[:n])), os.ModePerm)
		}
	} else {
		if "" == ifile {
			buf = readFromStdin(make([]byte, 1024*1024*4))
		} else {
			buf, _ = ioutil.ReadFile(ifile)
		}
		buf = spaceFilter(buf)
		buf = hexDecode(buf)
		values := make([]string, 0)
		value := uint64(0)
		n := 0
		for n < len(buf) {
			n = n + decode(buf, &value)
			values = append(values, strconv.FormatUint(value, 10))
		}
		if "" == ofile {
			fmt.Print(strings.Join(values, " "))
		} else {
			ioutil.WriteFile(ofile, []byte(strings.Join(values, " ")), os.ModePerm)
		}
	}
}

func readFromStdin(buf []byte) []byte {
	n, _ := os.Stdin.Read(buf)
	return buf[:n]
}

func spaceFilter(buf []byte) []byte {
	reg, _ := regexp.Compile("\\s+")
	return reg.ReplaceAll(buf, []byte{})
}

func spaceCutter(buf []byte) []uint64 {
	values := make([]uint64, 0)
	reg, _ := regexp.Compile("[0-9]+")
	strs := reg.FindAllString(string(buf), -1)
	for _, str := range strs {
		value, _ := strconv.ParseUint(str, 10, 64)
		values = append(values, value)
	}
	return values
}

func hexDecode(buf []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(buf)))
	hex.Decode(dst, buf)
	return dst
}

func hexEncode(buf []byte) string {
	return hex.EncodeToString(buf)
}

func encode(buf []byte, value uint64) (n int) {
	n = binary.PutUvarint(buf, value)
	return
}

func decode(buf []byte, value *uint64) (n int) {
	*value, n = binary.Uvarint(buf)
	return
}
