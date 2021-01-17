package controllers

import (
	"net/http"
	"regexp"
)

type heroController struct {
	heroIDPattern *regexp.Regexp
}

func (hc heroController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World, this is a type conversion, because a string is an alias for bytes!"))
}

// a constructor function starts with keyword 'new' followed by the object we are making
// this is not voodoo, it is a meer convetion.
// note: we return pointers, so that we edit directly and avoid unnecessary copy operations.
func newHeroController() *heroController {
	return &heroController{
		heroIDPattern: regexp.MustCompile(`^/heroes/(\d+)/?`),
	}
}
