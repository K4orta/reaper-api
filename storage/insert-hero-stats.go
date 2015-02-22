package storage

import (
	"github.com/gocql/gocql"
	"github.com/k4orta/reaper-api/stats"
	"log"
	"time"
)

func serializeStats(session *gocql.Session, hero *stats.Hero, h *stats.HeroStats) error {
	now := time.Unix(hero.LastUpdate, 0)
	return session.Query(`INSERT INTO hero_stats (hero_id, life, damage, toughness, healing, attack_speed, armor, strength, dexterity, vitality, intelligence, physical_resist, fire_resist, cold_resist, lightning_resist, poison_resist, arcane_resist, crit_damage, block_chance, block_amount_min, block_amount_max, damage_increase, crit_chance, damage_reduction, thorns, life_steal, life_per_kill, gold_find, magic_find, life_on_hit, primary_resource, secondary_resource, last_updated) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		hero.Id,
		h.Life,
		h.Damage,
		h.Toughness,
		h.Healing,
		h.AttackSpeed,
		h.Armor,
		h.Strength,
		h.Dexterity,
		h.Vitality,
		h.Intelligence,
		h.PhysicalResist,
		h.FireResist,
		h.ColdResist,
		h.LightningResist,
		h.PoisonResist,
		h.ArcaneResist,
		h.CritDamage,
		h.BlockChance,
		h.BlockAmountMin,
		h.BlockAmountMax,
		h.DamageIncrease,
		h.CritChance,
		h.DamageReduction,
		h.Thorns,
		h.LifeSteal,
		h.LifePerKill,
		h.GoldFind,
		h.MagicFind,
		h.LifeOnHit,
		h.PrimaryResource,
		h.SecondaryResource,
		now,
	).Exec()
}

func InsertHeroStats(hero *stats.Hero, data *stats.HeroStats) error {
	s, err := NewSession()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer s.Close()
	err = serializeStats(s, hero, data)
	if err != nil {
		log.Println(err)
	}
	return err
}
