package main

import (
	"log"
	"net/http"
	"os"
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
}

func testHexagon(w, h, hexSizePer float64) {
	colors := 20
	backgroundFill := 1.0
	path := "test.png"
	hexagon := newHexagon(w, h, hexSizePer)
	palette := newPalette(colors)
	hexagon.fill(palette, backgroundFill)
	hexagon.board.SavePNG(path)
}
