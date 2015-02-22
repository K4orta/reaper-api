package stats

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseCareer(t *testing.T) {
	fakeServer := makeFakeServer(stubCareerData())
	careerUrl = fakeServer.URL + "/"
	cd, err := FetchCareer("brackets-1829")

	if err != nil {
		t.Error("FetchCareer function had an error ", err)
	}

	if cd.BattleTag != "Brackets#1829" {
		t.Error("Failed to unmarshal BattleTag")
	}

	if len(cd.Heroes) != 2 {
		t.Error("Failed to unmarshal Heroes")
	}
}

func makeFakeServer(resp string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, resp)
	}))
}

func stubCareerData() string {
	return `
		{
			"battleTag": "Brackets#1829",
			"paragonLevel": 72,
			"paragonLevelHardcore": 0,
			"paragonLevelSeason": 0,
			"paragonLevelSeasonHardcore": 25,
			"heroes": [{
				"paragonLevel": 72,
				"seasonal": false,
				"name": "CaptanBadass",
				"id": 34876764,
				"level": 63,
				"hardcore": false,
				"gender": 0,
				"dead": false,
				"class": "barbarian",
				"last-updated": 1395813790
			}, 
			{
				"paragonLevel": 25,
				"seasonal": true,
				"name": "YiiIII",
				"id": 57645654,
				"level": 70,
				"hardcore": true,
				"gender": 0,
				"dead": false,
				"class": "barbarian",
				"last-updated": 1424627763
			}],
			"lastHeroPlayed": 57645654,
			"lastUpdated": 1424627763,
			"kills": {
				"monsters": 167077,
				"elites": 7539,
				"hardcoreMonsters": 71632
			},
			"highestHardcoreLevel": 70,
			"timePlayed": {
				"barbarian": 1.0,
				"crusader": 0.413,
				"demon-hunter": 0.042,
				"monk": 0.133,
				"witch-doctor": 0.0,
				"wizard": 0.274
			},
			"progression": {
				"act1": true,
				"act2": true,
				"act3": true,
				"act4": true,
				"act5": true
			},
			"fallenHeroes": [],
			"seasonalProfiles": {}
		}
	`
}
