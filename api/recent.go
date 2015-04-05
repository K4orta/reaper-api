package api

import (
	"encoding/json"
	"fmt"
	"github.com/k4orta/reaper-api/storage"
	"net/http"
)

func RecentActivity(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	hs, err := storage.Recent()

	if err == nil {
		out, _ := json.Marshal(hs)
		fmt.Fprint(w, string(out))
	} else {
		fmt.Fprint(w, `{"error": "There was an error"}`)
	}
}
