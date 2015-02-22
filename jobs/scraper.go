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
		time.Sleep(time.Minute)
	}
}

var (
	battleTags = []string{
		"brackets-1829",
	}
)

func scrape() {
	log.Println("Scraping...")
	for _, battleTag := range battleTags {
		log.Println("Fetching ", battleTag)
		profile, err := stats.FetchCareer(battleTag)
		if err != nil {
			log.Fatal("Error in scape1: ", err)
		}
		for _, hero := range profile.Heroes {
			heroData, err := stats.FetchHero(battleTag, hero.Id)
			if err != nil {
				log.Fatal("Error in scape2: ", err)
			}

			err = storage.InsertHeroStats(heroData, heroData.Stats)
			if err != nil {
				log.Fatal("Error in scape3: ", err)
			}
		}

	}
}
