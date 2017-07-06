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
	<a href="https://twitter.com/Bertoort">Twitter</a>
	<a href="https://github.com/berto/vango">Code</a>
</body>
</html>`

func index(response http.ResponseWriter, request *http.Request) {
	colors := 20
	percentFill := .50
	hex := newHexagon(1000, 1000, 20, 20)
	palette := newPalette(colors)
	hex.fill(palette, percentFill)
	err := hex.post()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(response, "%s", indexTemplate)
}
