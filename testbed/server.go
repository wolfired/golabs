package main

import (
	"time"

	"github.com/wolfired/golabs/server"
)

func main() {
	var logic_sev server.LogicServer = server.LogicServer{"tcp", ":8080"}
	go logic_sev.Run()

	var gate_sev server.GateServer = server.GateServer{"tcp", ":8081"}
	go gate_sev.Run()

	time.Sleep(60 * time.Minute)
}
