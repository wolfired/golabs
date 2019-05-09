package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"strings"
	"time"

	"net/http"
	"net/http/cookiejar"
	_ "net/http/pprof"

	"github.com/wolfired/golabs/wasm/wabf"
)

func main() {
	// gotv.Boot()
	// goss.Boot()
	// ptt.Boot()
}

func testPt() {
	c := http.Client{}
	c.Jar, _ = cookiejar.New(nil)

	c.Post("https://www.hyperay.org/takelogin.php", "application/x-www-form-urlencoded", strings.NewReader("username=wolfired&password=knil81hyperay&authcode=&trackerssl=yes"))

	r, _ := c.Get("https://www.hyperay.org/index.php")

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	fmt.Printf("%s\n", buf.String())
}

func testWasm() {
	cutWasm("./main.wasm", 0x1dbd+12)

	raw, _ := ioutil.ReadFile("./main.wasm_")
	wabf.Decode(raw)
}

func cutWasm(fileName string, len uint) {
	raw, _ := ioutil.ReadFile(fileName)
	ioutil.WriteFile(fileName+"_", raw[:len], os.ModePerm)
}

// 生成 CPU 报告
func cpuProfile() {
	f, err := os.OpenFile("cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.Println("CPU Profile started")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	fmt.Println("CPU Profile stopped")
}

// 生成堆内存报告
func heapProfile() {
	f, err := os.OpenFile("heap.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	pprof.WriteHeapProfile(f)
	fmt.Println("Heap Profile generated")
}

// 生成追踪报告
func traceProfile() {
	f, err := os.OpenFile("trace.out", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.Println("Trace started")
	trace.Start(f)
	defer trace.Stop()

	time.Sleep(60 * time.Second)
	fmt.Println("Trace stopped")
}
