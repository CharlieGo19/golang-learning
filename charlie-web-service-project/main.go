package main

// Note: a package is kind of like a class with respects to encapsulation.
// Package is just a directory that contains discrete units of code for some sort fo concept.

import (
	"net/http"

	"charlie.web.service/stonks/controllers"
)

func main() {
	// do go build .
	// windows: run stonks.exe and put https://localhost:3000 in browser.
	// unix: ./stonks and put https://localhost:3000 in browser.
	controllers.RegisterControllers()
	// look at front controller back controller pattern to understand the next line.
	// this spins up this application based on controllers defined (see line 13)
	http.ListenAndServe(":3000", nil)
}
