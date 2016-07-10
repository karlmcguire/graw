package graw

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
)

type Image struct {
	width  int
	height int
	img    draw.Image
}

func NewImage(width, height int) *Image {
	return &Image{
		width,
		height,
		image.NewNRGBA(image.Rect(0, 0, width, height)),
	}
}

func (self *Image) GetWidth() int {
	return self.width
}

func (self *Image) GetHeight() int {
	return self.height
}

func (self *Image) Rect(x0, y0, x1, y1 int, c color.NRGBA) {
	draw.Draw(
		self.img,
		image.Rect(
			x0,
			y0,
			x1,
			y1,
		),
		&image.Uniform{c},
		image.ZP,
		draw.Over,
	)
}

func (self *Image) Set(x, y int, c color.NRGBA) {
	self.Rect(x, y, x+1, y+1, c)
}

func (self *Image) Copy(x, y int, src *Image) {
	draw.Draw(
		self.img,
		image.Rect(
			x,
			y,
			x+src.width,
			y+src.height,
		),
		src.img,
		image.ZP,
		draw.Over,
	)
}

func (self *Image) Write(w io.Writer) {
	png.Encode(w, self.img)
}
