package goss

import (
	"flag"
	"os"
)

var flags = struct {
	help    bool
	network string
	host    string
	port    string
	pfPort  string
}{}

func parse() {
	flag.BoolVar(&flags.help, "help", false, "")

	flag.StringVar(&flags.network, "network", "tcp", "")
	flag.StringVar(&flags.host, "host", "0.0.0.0", "")
	flag.StringVar(&flags.port, "port", "8888", "")
	flag.StringVar(&flags.pfPort, "pfPort", "8889", "")

	flag.Parse()

	if flags.help {
		flag.Usage()
		os.Exit(0)
	}
}
