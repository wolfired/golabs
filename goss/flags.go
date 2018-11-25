package goss

import (
	"flag"
	"os"
)

var flags = struct {
	help    bool
	network string
	address string
	pfr     string
}{}

func parse() {
	flag.BoolVar(&flags.help, "help", false, "")

	flag.StringVar(&flags.network, "network", "tcp", "")
	flag.StringVar(&flags.address, "address", "0.0.0.0:8888", "")
	flag.StringVar(&flags.pfr, "pfr", "0.0.0.0:843", "")

	flag.Parse()

	if flags.help {
		flag.Usage()
		os.Exit(0)
	}
}
