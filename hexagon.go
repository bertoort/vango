package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

type Hexagon struct {
	board    *gg.Context
	currentX int
	currentY int
	width    float64
	height   float64
	size     int
	rows     int
}

func newHexagon(w, h, hexSizePer float64) Hexagon {
	width, height := calculateDimensions(h, hexSizePer)
	rowCount := (w - width) / width
	rows := h / height * 1.25
	size := calculateSize(rows, rowCount)
	return Hexagon{
		board:    gg.NewContext(int(w), int(h)),
		rows:     int(rowCount),
		width:    width,
		height:   height,
		size:     size,
		currentX: 1,
		currentY: 1,
	}
}

func calculateDimensions(h, percentage float64) (float64, float64) {
	height := h * percentage / 100
	r := height / 2
	width := math.Sqrt((r*r)-((r/2)*(r/2))) * 2
	return width, height
}

func calculateSize(rows, rowCount float64) int {
	formatRowCount := int(math.Floor(rowCount))
	hexSkipped := int(rows / 2)
	if int(rows)%2 != 0 {
		hexSkipped++
	}
	return formatRowCount*int(rows) - hexSkipped
}

func (h *Hexagon) drawHexagon(c Color) {
	sides := 6
	rotation := 100.0
	shift := h.rowShift()
	x := (float64(h.currentX) * h.width) + shift
	y := float64(h.currentY) * h.height * .75
	radius := h.height / 2
	h.board.DrawRegularPolygon(sides, x, y, radius, rotation)
	h.board.SetRGB(c.R, c.G, c.B)
	h.board.Fill()
	h.setNext(shift)
}

func (h *Hexagon) setNext(shift float64) {
	edge := 0
	if shift != 0 {
		edge = -1
	}
	if h.currentX == h.rows+edge {
		h.currentX = 1
		h.currentY++
	} else {
		h.currentX++
	}
}

func (h *Hexagon) rowShift() (shift float64) {
	if h.currentY%2 != 0 {
		shift = h.width / 2
	}
	return
}

func (h *Hexagon) fill(palette Palette, percentFill float64) {
	numOfColors := len(palette.colors)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < h.size; i++ {
		randNumber := random.Intn(numOfColors)
		if randNumber <= int(float64(numOfColors)*percentFill) {
			h.drawHexagon(palette.getColor(randNumber))
		} else {
			h.setNext(h.rowShift())
		}
	}
}
