package main

import (
	"math/rand"
	"time"

	"github.com/lucasb-eyer/go-colorful"
)

type Color struct {
	R, G, B float64
}

type Palette struct {
	colors []colorful.Color
}

func newColor(r, g, b float64) Color {
	return Color{
		R: r,
		G: g,
		B: b,
	}
}

func randomColor() Color {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return Color{
		R: random.Float64(),
		G: random.Float64(),
		B: random.Float64(),
	}
}

func newPalette(colors int) Palette {
	brownies, _ := colorful.SoftPaletteEx(10, colorful.SoftPaletteSettings{isSimilar, 50, true})
	return Palette{
		colors: brownies,
	}
}

func isSimilar(l, a, b float64) bool {
	c := colorful.Lab(l, a, b)
	return inRange(c.R) && inRange(c.G) && inRange(c.B)
}

func inRange(color float64) bool {
	base, limit := randomRange()
	return base < color && color < limit
}

func randomRange() (base, limit float64) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	base = random.Float64()
	limit = base + random.Float64()
	if limit > 1.0 {
		limit = 1.0
	}
	return
}

func (p *Palette) getColor(index int) Color {
	color := p.colors[index]
	return Color{
		R: color.R,
		G: color.G,
		B: color.B,
	}
}
