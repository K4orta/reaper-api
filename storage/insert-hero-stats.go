package storage

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/k4orta/reaper-api/stats"
	"log"
)

func serializeStats(session *sqlx.DB, hero *stats.Hero, h *stats.HeroStats) (sql.Result, error) {
	props := []string{
		"hero_id",
		"life",
		"damage",
		"toughness",
		"healing",
		"attack_speed",
		"armor",
		"strength",
		"dexterity",
		"vitality",
		"intelligence",
		"physical_resist",
		"fire_resist",
		"cold_resist",
		"lightning_resist",
		"poison_resist",
		"arcane_resist",
		"crit_damage",
		"block_chance",
		"block_amount_min",
		"block_amount_max",
		"damage_increase",
		"crit_chance",
		"damage_reduction",
		"thorns",
		"life_steal",
		"life_per_kill",
		"gold_find",
		"magic_find",
		"life_on_hit",
		"primary_resource",
		"secondary_resource",
	}

	h.LastUpdated = hero.LastUpdated

	_, err := session.NamedExec(`UPDATE hero_stats SET (`+serializeProps("", props)+`,last_updated) = (`+serializeProps(":", props)+`,to_timestamp(:last_updated))
						WHERE hero_id=:hero_id AND last_updated=to_timestamp(:last_updated)`, h)
	if err != nil {
		return nil, err
	}

	return session.NamedExec(`INSERT INTO hero_stats (`+serializeProps("", props)+`,last_updated) 
						SELECT `+serializeProps(":", props)+`,to_timestamp(:last_updated) 
						WHERE NOT EXISTS (SELECT 1 FROM hero_stats WHERE hero_id=:hero_id AND last_updated=to_timestamp(:last_updated));`, h)
}

func InsertHeroStats(hero *stats.Hero, data *stats.HeroStats) error {
	s, err := CreateConnection()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer s.Close()
	_, err = serializeStats(s, hero, data)
	if err != nil {
		log.Println(err)
	}
	return err
}
