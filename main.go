package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/k4orta/reaper-api/api"
	"github.com/k4orta/reaper-api/jobs"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", api.CreateUser).Methods("POST")
	router.HandleFunc("/users", api.ListUsers)
	router.HandleFunc("/stats/{id}", api.FetchHeroStats)
	router.HandleFunc("/heroes/{battleTag}", api.ListHeroes)
	router.HandleFunc("/recent", api.RecentActivity)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	n := negroni.New()
	n.Use(c)
	n.Use(negroni.NewStatic(http.Dir("webapp/public")))
	n.UseHandler(router)

	go jobs.StartScrape()

	n.Run(":8048")
}
