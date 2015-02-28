package storage

// CREATE TABLE reaper.heroes(id int, name text, class text, dead boolean, owner text, level int, last_updated timestamp, PRIMARY KEY(id));
// CREATE TABLE reaper.hero_stats(hero_id int, life float, damage float, toughness float, healing float, attack_speed float, armor float, strength float, dexterity float, vitality float, intelligence float, physical_resist float, fire_resist float, cold_resist float, lightning_resist float, poison_resist float, arcane_resist float, crit_damage float, block_chance float, block_amount_min float, block_amount_max float, damage_increase float, crit_chance float, damage_reduction float, thorns float, life_steal float, life_per_kill float, gold_find float, magic_find float, life_on_hit float, primary_resource float, secondary_resource float, last_updated timestamp, PRIMARY KEY (hero_id, last_updated));
import (
	"github.com/k4orta/reaper-api/stats"
	"github.com/relops/cqlr"
)

func FetchHeroes(battleTag string) ([]*stats.CharacterSummary, error) {
	s, err := NewSession()
	if err != nil {
		return nil, err
	}
	defer s.Close()
	row := s.Query(`SELECT * FROM heroes WHERE owner = ?`, battleTag)
	heroes := []*stats.CharacterSummary{}
	hero := stats.CharacterSummary{}

	b := cqlr.BindQuery(row)

	for b.Scan(&hero) {
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
