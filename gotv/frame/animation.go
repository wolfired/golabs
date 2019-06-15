package frame

import (
	"bytes"
	"compress/gzip"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strconv"
)

type frame []byte

// Animation 帧动画
type Animation struct {
	bitmapdatas []image.Image
}

// FromFile 从文件生成帧动画
func FromFile(file string, rowHei int, colWid int, count int) *Animation {
	f, _ := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	raw, _ := png.Decode(f)
	defer f.Close()

	pic, _ := raw.(*image.NRGBA)

	bitmapdataWid := pic.Bounds().Max.X
	bitmapdataHei := pic.Bounds().Max.Y

	blocks := bitmapdataWid / colWid * bitmapdataHei / rowHei

	if 0 < count {
		count = int(math.Min(float64(blocks), float64(count)))
	} else if blocks < -count {
		count = blocks
	} else {
		count += blocks
	}

	ani := &Animation{
		bitmapdatas: make([]image.Image, count),
	}

	index := 0

	for row := 0; row < bitmapdataHei; row += rowHei {
		for col := 0; col < bitmapdataWid && index < count; col += colWid {
			ani.bitmapdatas[index] = pic.SubImage(image.Rect(col, row, col+colWid, row+rowHei))

			index++
		}
	}

	return ani
}

// ToFiles 生成独立帧动画文件
func (a *Animation) ToFiles(file string) {
	for i, img := range a.bitmapdatas {
		f, _ := os.Create(file + strconv.Itoa(i) + ".png")
		defer f.Close()
		png.Encode(f, img)
	}
}

// ImageAt 截图
func (a *Animation) Len() int {
	return len(a.bitmapdatas)
}

// ImageAt 截图
func (a *Animation) ImageAt(i int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	offset := 0

	bs := a.FrameAt(i, false)

	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			c := color.RGBA{bs[offset+0], bs[offset+1], bs[offset+2], bs[offset+3]}

			img.Set(x, y, c)
		}

		offset += 4
	}

	return img
}

// FrameAt 生成独立帧动画文件
func (a *Animation) FrameAt(i int, z bool) []uint8 {
	pic, _ := a.bitmapdatas[i].(*image.NRGBA)

	raw := make([]byte, (pic.Rect.Max.X-pic.Rect.Min.X)*(pic.Rect.Max.Y-pic.Rect.Min.Y)*4)

	counter := 0

	for y := pic.Rect.Min.Y; y < pic.Rect.Max.Y; y++ {
		offset := y*pic.Stride + pic.Rect.Min.X

		for x := pic.Rect.Min.X; x < pic.Rect.Max.X; x++ {
			offset = (y-pic.Rect.Min.Y)*pic.Stride + (x-pic.Rect.Min.X)*4

			raw[counter+0] = pic.Pix[offset+0]
			raw[counter+1] = pic.Pix[offset+1]
			raw[counter+2] = pic.Pix[offset+2]
			raw[counter+3] = pic.Pix[offset+3]

			counter += 4
		}
	}

	if z {
		bf := bytes.NewBuffer([]byte{})

		gz := gzip.NewWriter(bf)

		gz.Write(raw)
		gz.Flush()
		gz.Close()

		return bf.Bytes()
	}

	return raw
}

type player struct {
}
