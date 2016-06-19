package graw

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
)

type Rectangle struct {
	width  int
	height int
	draw   draw.Image
}

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{
		width,
		height,
		image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

func (r *Rectangle) Set(x, y, xx, yy, hex int) {
	c := rgb(hex)
	l := image.Rect(x, y, xx, yy)

	draw.Draw(r.draw, l, c, image.ZP, draw.Src)
}

func (r *Rectangle) Copy(x, y int, rect *Rectangle) {
	l := image.Rect(x, y, x+rect.width, y+rect.height)
	draw.Draw(r.draw, l, rect.draw, image.ZP, draw.Src)
}

func (r *Rectangle) Write(w io.Writer) {
	png.Encode(w, r.draw)
}

func rgb(hex int) *image.Uniform {
	r := (hex >> 16) & 0xff
	g := (hex >> 8) & 0xff
	b := hex & 0xff

	return &image.Uniform{color.RGBA{uint8(r), uint8(g), uint8(b), 255}}
}
