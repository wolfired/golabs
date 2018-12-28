package ptt

import (
	"flag"
	"os"
)

var flags = struct {
	help     bool
	json     string
	accounts string
}{}

func parse() {
	flag.BoolVar(&flags.help, "help", false, "")

	flag.StringVar(&flags.json, "json", "./ptt.json", "")

	flag.Parse()

	if flags.help {
		flag.Usage()
		os.Exit(0)
	}
}
