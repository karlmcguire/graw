package graw

import (
	"image/color"
	"math"
)

func ipart(x float64) float64 {
	return math.Floor(x)
}

func round(x float64) float64 {
	return ipart(x + .5)
}

func fpart(x float64) float64 {
	return x - ipart(x)
}

func rfpart(x float64) float64 {
	return 1 - fpart(x)
}

func (self *Image) Line(x1, y1, x2, y2 float64, w int, col color.NRGBA) {
	dx := x2 - x1
	dy := y2 - y1
	ax := dx
	if ax < 0 {
		ax = -ax
	}
	ay := dy
	if ay < 0 {
		ay = -ay
	}

	var plot func(int, int, float64)

	if ax < ay {
		x1, y1 = y1, x1
		x2, y2 = y2, x2
		dx, dy = dy, dx
		plot = func(x, y int, c float64) {
			self.Set(y, x, color.NRGBA{col.R, col.G, col.B, uint8(255 * c)})
		}
	} else {
		plot = func(x, y int, c float64) {
			self.Set(x, y, color.NRGBA{col.R, col.G, col.B, uint8(255 * c)})
		}
	}
	if x2 < x1 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	gradient := dy / dx

	xend := round(x1)
	yend := y1 + gradient*(xend-x1)
	xgap := rfpart(x1 + .5)
	xpxl1 := int(xend)
	ypxl1 := int(ipart(yend))
	plot(xpxl1, ypxl1, rfpart(yend)*xgap)
	plot(xpxl1, ypxl1+1, fpart(yend)*xgap)
	intery := yend + gradient

	xend = round(x2)
	yend = y2 + gradient*(xend-x2)
	xgap = fpart(x2 + 0.5)
	xpxl2 := int(xend)
	ypxl2 := int(ipart(yend))
	plot(xpxl2, ypxl2, rfpart(yend)*xgap)
	plot(xpxl2, ypxl2+1, fpart(yend)*xgap)

	for x := xpxl1 + 1; x <= xpxl2-1; x++ {
		plot(x, int(ipart(intery)), rfpart(intery))
		plot(x, int(ipart(intery))+1, fpart(intery))
		intery = intery + gradient
	}
}
