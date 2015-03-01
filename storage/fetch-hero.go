package storage

import (
	"github.com/k4orta/reaper-api/stats"
)

func FetchHeroStats(heroId int64) ([]*stats.HeroStats, error) {
	s, err := CreateConnection()
	if err != nil {
		return nil, err
	}
	defer s.Close()
	heroStats := []*stats.HeroStats{}
	rows, err := s.Queryx(`SELECT * FROM hero_stats WHERE hero_id = $1;`, heroId)
	hs := stats.HeroStats{}
	for rows.Next() {
		rows.StructScan(&hs)
		heroStats = append(heroStats, &stats.HeroStats{
			hs.HeroId,
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
