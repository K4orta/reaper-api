package api

import (
	"encoding/json"
	"fmt"
	"github.com/k4orta/reaper-api/storage"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var data map[string]string
	d, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprint(w, `{"error": "There was an error reading post body"}`)
		return
	}
	_ = json.Unmarshal(d, &data)

	// TODO: Validation logic
	if val, ok := data["battleTag"]; ok && ValidateBattleTag(val) {
		storage.InsertUser(data["battleTag"])
		fmt.Fprintln(w, `{"created": "`+data["battleTag"]+`"}`)
	} else {
		sendError(w, "No valid BattleTag found in response")
	}
}

func ValidateBattleTag(input string) bool {
	return true
}

// Dev only
func ListUsers(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := storage.FetchAllUsers()
	if err != nil {
		sendError(w, "Error")
		return
	}

	out, err := json.Marshal(users)
	if err == nil {
		fmt.Fprint(w, string(out))
	} else {
		sendError(w, "Error")
	}
}
