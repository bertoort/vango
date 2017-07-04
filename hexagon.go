package main

import "github.com/fogleman/gg"

type Hexagon struct {
	board    *gg.Context
	currentX int
	currentY int
	size     float64
	rows     float64
	columns  float64
}

func newHexagon(w, h int, r, c float64) Hexagon {
	size := float64(w) / r
	return Hexagon{
		board:    gg.NewContext(w, h),
		rows:     r,
		columns:  c,
		size:     size,
		currentX: 1,
		currentY: 1,
	}
}

func (h *Hexagon) drawHexagon(c Color) {
	sides := 6
	rotation := 100.00
	shift := h.rowShift()
	x := float64(h.currentX)*h.size + shift
	y := float64(h.currentY) * h.size * .9
	h.board.DrawRegularPolygon(sides, x, y, h.size/2, rotation)
	h.board.SetRGB255(c.red, c.green, c.blue)
	h.board.Fill()
	h.setNext(shift)
}

func (h *Hexagon) setNext(shift float64) {
	edge := 1
	if shift != 0 {
		edge = 2
	}
	if h.currentX == int(h.rows)-edge {
		h.currentX = 1
		h.currentY++
	} else {
		h.currentX++
	}
}

func (h *Hexagon) rowShift() (shift float64) {
	if h.currentY%2 != 0 {
		shift = h.size / 2
	}
	return
}

func (h *Hexagon) write(path string) {
	h.board.SavePNG(path)
}
