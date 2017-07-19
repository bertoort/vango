package main

import (
	"math/rand"
	"time"

	"github.com/lucasb-eyer/go-colorful"
)

type Color interface {
	values() (int, int, int, int)
}

type RGB struct {
	R, G, B int
}

type RGBA struct {
	R, G, B, A int
}

func (r RGB) values() (int, int, int, int) {
	return r.R, r.G, r.B, 1
}

func (r RGBA) values() (int, int, int, int) {
	return r.R, r.G, r.B, r.A
}

type Palette struct {
	colors []colorful.Color
}

func newRGB(r, g, b int) RGB {
	return RGB{
		R: r,
		G: g,
		B: b,
	}
}

func newRGBA(r, g, b, a int) RGBA {
	return RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

func randomRGB() RGB {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return RGB{
		R: random.Int(),
		G: random.Int(),
		B: random.Int(),
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

func convertRGBA255(r, g, b, a uint32) (int, int, int, int) {
	fa := int(a/257) / 255
	if a == 0xffff {
		fa = 1
	}
	if a == 0 {
		return 0, 0, 0, 100
	}
	r = (r * 0xffff) / a
	g = (g * 0xffff) / a
	b = (b * 0xffff) / a
	return 255 - int(r/257), 255 - int(g/257), 255 - int(b/257), fa
}

func (p *Palette) getRGB(index int) RGB {
	color := p.colors[index]
	return RGB{
		R: int(color.R * 255),
		G: int(color.G * 255),
		B: int(color.B * 255),
	}
}
