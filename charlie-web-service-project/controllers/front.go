package controllers

import "net/http"

func RegisterControllers() {
	// this functions handles all the routing of the service.
	var heroController *heroController = newHeroController()
	http.Handle("/heroes", heroController)
	http.Handle("/heroes/", heroController)
}
