package gotv

import (
	"flag"
	"os"
)

var flags = struct {
	help   bool
	title  string
	addr   string
	host   string
	ws     string
	js     string
	width  int
	height int
}{}

func parse() {
	flag.BoolVar(&flags.help, "help", false, "")

	flag.StringVar(&flags.title, "title", "GoTv", "")

	flag.StringVar(&flags.addr, "addr", "0.0.0.0:9999", "")
	flag.StringVar(&flags.host, "host", "127.0.0.1:9999", "")

	flag.StringVar(&flags.ws, "ws", "/gotv", "")
	flag.StringVar(&flags.js, "js", "./gotv.js", "")

	flag.IntVar(&flags.width, "width", 160, "")
	flag.IntVar(&flags.height, "height", 144, "")

	flag.Parse()

	if flags.help {
		flag.Usage()
		os.Exit(0)
	}
}
