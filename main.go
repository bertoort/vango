package main

import (
	"image"
	"log"
	"net/http"
	"os"

	_ "image/jpeg"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	h := http.NewServeMux()
	h.HandleFunc("/", index)
	h.HandleFunc("/tweet", tweet)
	err := http.ListenAndServe(":"+port, h)
	if err != nil {
		log.Fatal(err)
	}
	err := tweetImage(hexagon.board.Image())
	if err != nil {
		log.Fatal(err)
	}

	// TODO: run the following scripts from CLI flags
	// testDraw(4, "./examples/chromie-original.jpg", "./examples/chromie-hexagons.jpg")
	// testHexagon(1000, 1000, 10)
}

func getImage(path string) (image.Image, error) {
	reader, _ := os.Open(path)
	defer reader.Close()
	im, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return im, err
}

func testDraw(sizePercentage float64, inPath, outPath string) {
	im, err := getImage(inPath)
	if err != nil {
		log.Fatal(err)
	}
	bounds := im.Bounds()
	y := float64(bounds.Max.Y)
	x := float64(bounds.Max.X)
	hexagon := newHexagon(y, x, sizePercentage)
	hexagon.draw(im)
	hexagon.board.SavePNG(outPath)
}

func testHexagon(w, h, hexSizePer float64) {
	colors := 20
	backgroundFill := 1.0
	path := "./examples/hexagon.png"
	hexagon := newHexagon(w, h, hexSizePer)
	palette := newPalette(colors)
	hexagon.fill(palette, backgroundFill)
	hexagon.board.SavePNG(path)
}
