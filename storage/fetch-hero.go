package storage

import (
	"github.com/k4orta/reaper-api/stats"
	"github.com/relops/cqlr"
)

func FetchHeroStats(heroId int64) ([]*stats.HeroStats, error) {
	s, err := NewSession()
	if err != nil {
		return nil, err
	}
	defer s.Close()
	row := s.Query(`SELECT * FROM hero_stats WHERE hero_id = ?`, heroId)

	b := cqlr.BindQuery(row)
	heroStats := []*stats.HeroStats{}

	hs := stats.HeroStats{}
	for b.Scan(&hs) {
		heroStats = append(heroStats, &stats.HeroStats{
			hs.Life,
			hs.Damage,
			hs.Toughness,
			hs.Healing,
			hs.AttackSpeed,
			hs.Armor,
			hs.Strength,
			hs.Dexterity,
			hs.Vitality,
			hs.Intelligence,
			hs.PhysicalResist,
			hs.FireResist,
			hs.ColdResist,
			hs.LightningResist,
			hs.PoisonResist,
			hs.ArcaneResist,
			hs.CritChance,
			hs.BlockChance,
			hs.BlockAmountMin,
			hs.BlockAmountMax,
			hs.DamageIncrease,
			hs.CritChance,
			hs.DamageReduction,
			hs.Thorns,
			hs.LifeSteal,
			hs.LifePerKill,
			hs.GoldFind,
			hs.MagicFind,
			hs.LifeOnHit,
			hs.PrimaryResource,
			hs.SecondaryResource,
			hs.LastUpdated,
		})
	}

	return heroStats, nil
}
