package main

func main() {
	colors := 20
	percentFill := .50
	hex := newHexagon(1000, 1000, 20, 20)
	palette := newPalette(colors)
	hex.fill(palette, percentFill)
	hex.write("test.png")
}
