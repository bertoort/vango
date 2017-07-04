package main

func main() {
	hex := newHexagon(1000, 1000, 20, 20)
	green := newColor(12, 232, 226)
	for i := 0; i < 388; i++ {
		hex.drawHexagon(green)
	}
	hex.write("hexagon.png")
}
