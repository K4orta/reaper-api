package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/profile/{battleTag}", api.Profile)
	n := negroni.New()
	n.Use(negroni.NewStatic(http.Dir("webapp/public")))
	n.UseHandler(router)

	go tasks.StartScrape()

	n.Run(":8048")
}
