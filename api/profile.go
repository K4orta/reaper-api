package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/k4orta/reaper-api/storage"
	"net/http"
)

func ListHeroes(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	hs, err := storage.FetchHeroes(vars["battleTag"])

	if err == nil {
		out, _ := json.Marshal(hs)
		fmt.Fprint(w, string(out))
	} else {
		fmt.Fprint(w, `{"error": "There was an error"}`)
	}
}
