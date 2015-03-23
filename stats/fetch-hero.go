package stats

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Hero struct {
	Id          int        `json:"id"`
	Level       int        `json:"level"`
	Paragon     int        `json:"paragon"`
	Stats       *HeroStats `json:"stats"`
	Name        string     `json:"name"`
	LastUpdated int64      `json:"last-updated"`
}

type HeroStats struct {
	HeroId            int64   `json:"heroId" db:"hero_id"`
	Life              float32 `json:"life" db:"life"`
	Damage            float32 `json:"damage" db:"damage"`
	Toughness         float32 `json:"toughness" db:"toughness"`
	Healing           float32 `json:"healing" db:"healing"`
	AttackSpeed       float32 `json:"attackSpeed" db:"attack_speed"`
	Armor             float32 `json:"armor" db:"armor"`
	Strength          float32 `json:"strength" db:"strength"`
	Dexterity         float32 `json:"dexterity" db:"dexterity"`
	Vitality          float32 `json:"vitality" db:"vitality"`
	Intelligence      float32 `json:"intelligence" db:"intelligence"`
	PhysicalResist    float32 `json:"physicalResist" db:"physical_resist"`
	FireResist        float32 `json:"fireResist" db:"fire_resist"`
	ColdResist        float32 `json:"coldResist" db:"cold_resist"`
	LightningResist   float32 `json:"lightningResist" db:"lightning_resist"`
	PoisonResist      float32 `json:"poisonResist" db:"poison_resist"`
	ArcaneResist      float32 `json:"arcaneResist" db:"arcane_resist"`
	CritDamage        float32 `json:"critDamage" db:"crit_damage"`
	BlockChance       float32 `json:"blockChance" db:"block_chance"`
	BlockAmountMin    float32 `json:"blockAmountMin" db:"block_amount_min"`
	BlockAmountMax    float32 `json:"blockAmountMax" db:"block_amount_max"`
	DamageIncrease    float32 `json:"damageIncrease" db:"damage_increase"`
	CritChance        float32 `json:"critChance" db:"crit_chance"`
	DamageReduction   float32 `json:"damageReduction" db:"damage_reduction"`
	Thorns            float32 `json:"thorns" db:"thorns"`
	LifeSteal         float32 `json:"lifeSteal" db:"life_steal"`
	LifePerKill       float32 `json:"lifePerKill" db:"life_per_kill"`
	GoldFind          float32 `json:"goldFind" db:"gold_find"`
	MagicFind         float32 `json:"magicFind" db:"magic_find"`
	LifeOnHit         float32 `json:"lifeOnHit" db:"life_on_hit"`
	PrimaryResource   float32 `json:"primaryResource" db:"primary_resource"`
	SecondaryResource float32 `json:"secondaryResource" db:"secondary_resource"`
	LastUpdated       int64   `json:"last-updated" db:"last_updated"`
}

var (
	heroUrl = "https://us.api.battle.net/d3/profile"
)

func FetchHero(battleTag string, heroId int64) (*Hero, error) {
	hid := strconv.FormatInt(heroId, 10)
	resp, err := http.Get(FormatUrl(careerUrl + "/" + battleTag + "/hero/" + hid))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var hd Hero
	err = json.Unmarshal([]byte(b), &hd)
	if err != nil || hd.Stats == nil {
		return nil, err
	}
	hd.Stats.HeroId = heroId

	return &hd, nil
}
