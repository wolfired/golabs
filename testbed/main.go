package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
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

	"github.com/gen2brain/x264-go"
	"github.com/wolfired/golabs/go2clabs"
	"github.com/wolfired/golabs/gotv/frame"
	"github.com/wolfired/golabs/wasm/wabf"
)

func main() {
	bin := []byte{
		0x01, 0x04, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x6f,
		0x6c, 0x65, 0x06, 0x6c, 0x6f, 0x67, 0x0c, 0x53,
		0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x65, 0x78,
		0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x68,
		0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x6a, 0x73, 0x0d,
		0x00, 0x02, 0x00, 0x9e, 0x01, 0x00, 0x01, 0x00,
		0x03, 0x00, 0x00, 0x14, 0x01, 0xa0, 0x01, 0x00,
		0x00, 0x00, 0x38, 0xc4, 0x00, 0x00, 0x00, 0x42,
		0xc5, 0x00, 0x00, 0x00, 0x04, 0xc6, 0x00, 0x00,
		0x00, 0x27, 0x01, 0x00, 0xd2, 0x2b, 0x8e, 0x03,
		0x01, 0x00,
	}

	buf, _ := ioutil.ReadFile("./sample.js")

	js := go2clabs.NewJS()
	js.EvalBinary([]string{"a", "b"}, bin)
	js.EvalSource([]string{"a", "b"}, buf)

	// con0 := display.CreateContainer("Con0")

	// fmt.Printf("%p %p %p %p", con0, &con0.Object, con0.Object.IEntity, con0.Object.Child)

	// frame.FromFile("./spritesheet_magicbubbles.png", 100, 100, -3) //.ToFiles("./magicbubbles")
	// f, _ := os.Create("png.gzip")
	// defer f.Close()
	// f.Write(frame.FromFile("./spritesheet_magicbubbles.png", 100, 100, -3).FrameAt(0))
	// gotv.Boot()
	// goss.Boot()
	// ptt.Boot()
	// stage := display.CreateContainer("stage")
	// obj0 := display.CreateObject("obj0")
	// fmt.Println(obj0.Name(), stage.Name())
	// stage.Add(obj0)
	// ecs.CreateEntity("a").DryGet(comps.CCTransform).(comps.Transform).Exec()

	// {
	// 	rand.Seed(time.Now().UnixNano())
	// 	bs := make([]byte, 320*288*4)

	// 	for i := 0; i < len(bs); i++ {
	// 		bs[i] = byte(rand.Intn(256))
	// 	}

	// 	g, _ := os.Create("time.gzip")
	// 	defer g.Close()
	// 	gw, _ := gzip.NewWriterLevel(g, gzip.BestCompression)
	// 	gw.Write(bs)

	// 	img := image.NewNRGBA(image.Rect(0, 0, 320, 288))

	// 	for i := 0; i < len(bs); i++ {
	// 		img.Pix[i] = bs[i]
	// 	}

	// 	p, _ := os.Create("time.png")
	// 	defer p.Close()

	// 	png.Encode(p, img)
	// }

	// testX264()
}

func testX264() {
	buf := bytes.NewBuffer(make([]byte, 0))

	wid := 100
	hei := 100

	opts := &x264.Options{
		Width:     wid,
		Height:    hei,
		FrameRate: 24,
		Tune:      "zerolatency",
		Preset:    "veryfast",
		Profile:   "baseline",
		LogLevel:  x264.LogDebug,
	}

	enc, err := x264.NewEncoder(buf, opts)
	if err != nil {
		panic(err)
	}

	img := x264.NewYCbCr(image.Rect(0, 0, opts.Width, opts.Height))
	draw.Draw(img, img.Bounds(), image.Black, image.ZP, draw.Src)

	a := frame.FromFile("./spritesheet_magicbubbles.png", 100, 100, -3)

	for i := 0; i < 24*60; i++ {
		raw := a.FrameAt(i%a.Len(), false)

		offset := 0

		for x := 0; x < wid; x++ {
			for y := 0; y < hei; y++ {
				img.Set(x, y, color.NRGBA{raw[offset+0], raw[offset+1], raw[offset+2], raw[offset+3]})
			}
		}

		err = enc.Encode(img)
		if err != nil {
			panic(err)
		}
	}

	err = enc.Flush()
	if err != nil {
		panic(err)
	}

	err = enc.Close()
	if err != nil {
		panic(err)
	}

	f, _ := os.Create("x.264")
	defer f.Close()

	f.Write(buf.Bytes())
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
