package main

type Color struct {
	red   int
	green int
	blue  int
}

func newColor(r, g, b int) Color {
	return Color{
		red:   r,
		green: g,
		blue:  b,
	}
}
