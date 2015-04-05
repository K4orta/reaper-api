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
		time.Sleep(time.Minute * 10)
	}
}

func scrape() {
	log.Println("Scraping...")

	users, _ := storage.FetchAllUsers()

	for _, user := range users {
		log.Println("Fetching", user.BattleTag)
		profile, err := stats.FetchCareer(user.BattleTag)
		if err != nil {
			log.Println("Error in scape: ", err)
			return
		}
		for _, hero := range profile.Heroes {
			// Only scrape the hero if it was updated after this account was last logged
			if hero.LastUpdated > user.LastUpdated {
				log.Println("Scraping", hero.Name)
				heroData, err := stats.FetchHero(user.BattleTag, hero.Id)

				if err != nil {
					log.Println("Error in scape: ", err)
					return
				}

				err = storage.InsertHeroSummary(hero)
				if err != nil {
					log.Println("Error inserting Hero: ", err)
					return
				}

				err = storage.InsertHeroStats(heroData, heroData.Stats)
				if err != nil {
					log.Println("Error in scape: ", err)
					return
				}
			}
		}
		user.LastUpdated = time.Now().Unix()
		_ = storage.UpdateUser(user)
	}
}
