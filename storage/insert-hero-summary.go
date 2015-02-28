package storage

import (
	"github.com/gocql/gocql"
	"github.com/k4orta/reaper-api/stats"
	"github.com/relops/cqlr"
	"log"
)

func serializeHeroSummary(session *gocql.Session, hero *stats.CharacterSummary) error {
	b := cqlr.Bind(`INSERT INTO heroes (id, name, dead, owner, class, level, last_updated) VALUES (?,?,?,?,?,?,?)`, hero)
	if err := b.Exec(session); err != nil {
		log.Fatal("Error inserting Hero", err)
	}

	return nil
}

func InsertHeroSummary(hero *stats.CharacterSummary) error {
	s, err := NewSession()
	if err != nil {
		log.Fatal("Error connecting to DB in InsertHeroSummary", err)
	}
	defer s.Close()
	err = serializeHeroSummary(s, hero)
	if err != nil {
		log.Fatal("Error serializing in InsertHeroSummary", err)
	}

	return nil
}
