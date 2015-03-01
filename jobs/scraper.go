package jobs

import (
	"github.com/k4orta/reaper-api/stats"
	"github.com/k4orta/reaper-api/storage"
	"log"
	"time"
)

func StartScrape() {
	for {
		scrape()
		time.Sleep(time.Minute * 5)
	}
}

var (
	battleTags = []string{
		"brackets-1829",
		"dubs-1649",
		"obadiah-1183",
		"darkblob32-1791",
	}
)

func scrape() {
	log.Println("Scraping...")
	for _, battleTag := range battleTags {
		log.Println("Fetching ", battleTag)
		profile, err := stats.FetchCareer(battleTag)
		if err != nil {
			log.Fatal("Error in scape: ", err)
		}
		for _, hero := range profile.Heroes {
			heroData, err := stats.FetchHero(battleTag, hero.Id)
			if err != nil {
				log.Fatal("Error in scape: ", err)
			}
			err = storage.InsertHeroSummary(hero)
			if err != nil {
				log.Fatal("Error inserting Hero: ", err)
			}
			err = storage.InsertHeroStats(heroData, heroData.Stats)
			if err != nil {
				log.Fatal("Error in scape: ", err)
			}
		}

	}
}
