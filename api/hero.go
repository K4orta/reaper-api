package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/k4orta/reaper-api/storage"
	"net/http"
	"strconv"
)

func FetchHeroStats(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	i, _ := strconv.ParseInt(vars["id"], 10, 64)
	hs, err := storage.FetchHeroStats(i)

	if err == nil {
		out, _ := json.Marshal(hs)
		fmt.Fprint(w, string(out))
	} else {
		fmt.Fprint(w, `{"error": "There was an error"}`)
	}
}
