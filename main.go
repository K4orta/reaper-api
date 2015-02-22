package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/k4orta/reaper-api/jobs"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	// router.HandleFunc("/profile/{battleTag}", api.Profile)
	n := negroni.New()
	n.Use(negroni.NewStatic(http.Dir("webapp/public")))
	n.UseHandler(router)

	go jobs.StartScrape()

	n.Run(":8048")
}
