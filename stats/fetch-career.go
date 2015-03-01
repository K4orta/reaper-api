package stats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	careerUrl = "https://us.api.battle.net/d3/profile"
)

type Profile struct {
	BattleTag   string              `json:"battleTag"`
	Heroes      []*CharacterSummary `json:"heroes"`
	LastUpdated string              `json:"last-updated"`
}

type CharacterSummary struct {
	Name        string `json:"name" db:"name"`
	Id          int64  `json:"id" db:"id"`
	Class       string `json:"class" db:"class"`
	Level       int    `json:"level" db:"level"`
	Dead        bool   `json:"dead" db:"dead"`
	LastUpdated int64  `json:"last-updated" db:"last_updated"`
	Owner       string `json:"owner" db:"owner"`
}

func FetchCareer(battleTag string) (*Profile, error) {
	fmt.Println(FormatUrl(careerUrl + "/" + battleTag + "/"))
	resp, err := http.Get(FormatUrl(careerUrl + "/" + battleTag + "/"))
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

	for _, h := range cd.Heroes {
		h.Owner = battleTag
	}

	return &cd, nil
}
