package storage

import (
	"github.com/k4orta/reaper-api/stats"
	"log"
)

func Recent() ([]*stats.CharacterSummary, error) {
	db, err := CreateConnection()
	heroes := []*stats.CharacterSummary{}
	if err != nil {
		log.Fatal(err)
		return heroes, err
	}
	defer db.Close()
	hero := stats.CharacterSummary{}
	rows, err := db.Queryx(`SELECT * FROM heroes ORDER BY last_updated DESC LIMIT 5;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.StructScan(&hero)
		heroes = append(heroes, &stats.CharacterSummary{
			hero.Name,
			hero.Id,
			hero.Class,
			hero.Level,
			hero.Dead,
			hero.LastUpdated,
			hero.Owner,
		})
	}

	return heroes, nil
}
