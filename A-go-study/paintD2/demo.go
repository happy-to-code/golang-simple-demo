package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dx)
	for i := range pic {
		pic[i] = make([]uint8, dy)
		for j := range pic[i] {
			pic[i][j] = uint8(i * j)
		}
	}
	return pic
}

func main() {
	Show(Pic)
}

func Show(f func(int, int) [][]uint8) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy) // 图片坐标点的颜色二维数组。
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	ShowImage(m)
}

func ShowImage(m image.Image) {

	// 需要保存的文件
	imgcounter := 1234
	imgfile, _ := os.Create(fmt.Sprintf("%03d.png", imgcounter))
	defer imgfile.Close()

	// 以PNG格式保存文件
	err := png.Encode(imgfile, m)
	if err != nil {
		log.Fatal(err)
	}

}
