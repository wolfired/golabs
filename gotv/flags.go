package gotv

import (
	"flag"
	"os"
)

var flags = struct {
	help  bool
	title string
	addr  string
	host  string
	port  string
	ws    string
	js    string

	spritesheet string

	width  int
	height int
}{}

func parse() {
	flag.BoolVar(&flags.help, "help", false, "")

	flag.StringVar(&flags.title, "title", "GoTv", "")

	flag.StringVar(&flags.addr, "addr", "0.0.0.0", "")
	flag.StringVar(&flags.host, "host", "127.0.0.1", "")
	flag.StringVar(&flags.port, "port", "9999", "")

	flag.StringVar(&flags.ws, "ws", "/gotv", "")
	flag.StringVar(&flags.js, "js", "./gotv.js", "")

	flag.StringVar(&flags.spritesheet, "spritesheet", "./spritesheet.png", "")

	flag.IntVar(&flags.width, "width", 160, "The Game Boy width is 160")
	flag.IntVar(&flags.height, "height", 144, "The Game Boy height is 144")

	flag.Parse()

	if flags.help {
		flag.Usage()
		os.Exit(0)
	}
}
