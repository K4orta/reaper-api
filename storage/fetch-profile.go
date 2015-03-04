package storage

// create table heroes (id bigint NOT NULL PRIMARY KEY, name text NOT NULL, class text NOT NULL, dead boolean NOT NULL, owner text NOT NULL, level int NOT NULL, last_updated bigint NOT NULL);

// CREATE TABLE hero_stats(hero_id bigint NOT NULL, life double precision NOT NULL, damage double precision NOT NULL, toughness double precision NOT NULL, healing double precision NOT NULL, attack_speed double precision NOT NULL, armor double precision NOT NULL, strength double precision NOT NULL, dexterity double precision NOT NULL, vitality double precision NOT NULL, intelligence double precision NOT NULL, physical_resist double precision NOT NULL, fire_resist double precision NOT NULL, cold_resist double precision NOT NULL, lightning_resist double precision NOT NULL, poison_resist double precision NOT NULL, arcane_resist double precision NOT NULL, crit_damage double precision NOT NULL, block_chance double precision NOT NULL, block_amount_min double precision NOT NULL, block_amount_max double precision NOT NULL, damage_increase double precision NOT NULL, crit_chance double precision NOT NULL, damage_reduction double precision NOT NULL, thorns double precision NOT NULL, life_steal double precision NOT NULL, life_per_kill double precision NOT NULL, gold_find double precision NOT NULL, magic_find double precision NOT NULL, life_on_hit double precision NOT NULL, primary_resource double precision NOT NULL, secondary_resource double precision NOT NULL, last_updated bigint NOT NULL);

import (
	"github.com/k4orta/reaper-api/stats"
)

func FetchHeroes(battleTag string) ([]*stats.CharacterSummary, error) {
	db, err := CreateConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	heroes := []*stats.CharacterSummary{}
	hero := stats.CharacterSummary{}
	rows, err := db.Queryx(`SELECT * FROM heroes WHERE owner = $1 ORDER BY last_updated DESC;`, battleTag)
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
