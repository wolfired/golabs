package main

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"os"

	"github.com/gen2brain/x264-go"
	"github.com/wolfired/golabs/gotv/frame"
)

func main() {
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
