package storage

import (
	"github.com/k4orta/reaper-api/stats"
	"log"
)

func InsertHeroSummary(hero *stats.CharacterSummary) error {
	c, err := CreateConnection()
	if err != nil {
		log.Fatal("Error connecting to DB in InsertHeroSummary", err)
	}

	defer c.Close()
	props := []string{
		"id",
		"name",
		"dead",
		"owner",
		"class",
		"level",
		"last_updated",
	}

	_, err = c.NamedExec(`UPDATE heroes SET (`+serializeProps("", props)+`) = (`+serializeProps(":", props)+`) WHERE id=:id;`, hero)
	_, err = c.NamedExec(`INSERT INTO heroes (`+serializeProps("", props)+`) 
							SELECT `+serializeProps(":", props)+`
							WHERE NOT EXISTS (SELECT 1 FROM heroes WHERE id=:id);`, hero)

	if err != nil {
		log.Fatal("Error serializing in InsertHeroSummary", err)
	}

	return nil
}
