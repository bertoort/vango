package main

import (
	"fmt"
	"log"
	"net/http"
)

const indexTemplate = `<html>
<head>
<title>Vango Bot</title>
</head>
<body>
	<img width="300" height="300" src="https://raw.githubusercontent.com/berto/vango/master/hexagon.png">
	<br>
	<form action="/tweet">
		<button type="submit">Post New Image</button>
	</form>
	<br>
	<a href="https://twitter.com/Bertoort">Twitter</a>
	<a href="https://github.com/berto/vango">Code</a>
</body>
</html>`

func index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "%s", indexTemplate)
}

func tweet(response http.ResponseWriter, request *http.Request) {
	colors := 20
	percentFill := .50
	hexagon := newHexagon(1000, 1000, 2)
	palette := newPalette(colors)
	hexagon.fill(palette, percentFill)
	err := tweetImage(hexagon.board.Image())
	if err != nil {
		log.Fatal(err)
	}
	url := "https://twitter.com/Bertoort"
	http.Redirect(response, request, url, http.StatusTemporaryRedirect)
}
