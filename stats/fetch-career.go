package stats

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	careerUrl = "https://us.api.battle.net/d3/profile"
)

type Profile struct {
	BattleTag string              `json:"battleTag"`
	Heroes    []*CharacterSummary `json:"heroes"`
}

type CharacterSummary struct {
	Name        string
	Id          int64
	Class       string
	Level       int
	LastUpdated int64 `json:"last-updated"`
}

func FetchCareer(battleTag string) (*Profile, error) {
	resp, err := http.Get(careerUrl + "/" + battleTag + "/?locale=en_US&apikey=")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cd Profile

	json.Unmarshal([]byte(b), &cd)
	return &cd, nil
}
