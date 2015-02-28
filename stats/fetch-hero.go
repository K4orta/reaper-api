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
	Life              float32 `json:"life" cql:"life"`
	Damage            float32 `json:"damage" cql:"damage"`
	Toughness         float32 `json:"toughness" cql:"toughness"`
	Healing           float32 `json:"healing" cql:"healing"`
	AttackSpeed       float32 `json:"attackSpeed" cql:"attack_speed"`
	Armor             float32 `json:"armor" cql:"armor"`
	Strength          float32 `json:"strength" cql:"strength"`
	Dexterity         float32 `json:"dexterity" cql:"dexterity"`
	Vitality          float32 `json:"vitality" cql:"vitality"`
	Intelligence      float32 `json:"intelligence" cql:"intelligence"`
	PhysicalResist    float32 `json:"physicalResist" cql:"physical_resist"`
	FireResist        float32 `json:"fireResist" cql:"fire_resist"`
	ColdResist        float32 `json:"coldResist" cql:"cold_resist"`
	LightningResist   float32 `json:"lightningResist" cql:"lightning_resist"`
	PoisonResist      float32 `json:"poisonResist" cql:"poison_resist"`
	ArcaneResist      float32 `json:"arcaneResist" cql:"arcane_resist"`
	CritDamage        float32 `json:"critDamage" cql:"crit_damage"`
	BlockChance       float32 `json:"blockChance" cql:"block_chance"`
	BlockAmountMin    float32 `json:"blockAmountMin" cql:"block_amount_min"`
	BlockAmountMax    float32 `json:"blockAmountMax" cql:"block_amount_max"`
	DamageIncrease    float32 `json:"damageIncrease" cql:"damage_increase"`
	CritChance        float32 `json:"critChance" cql:"crit_chance"`
	DamageReduction   float32 `json:"damageReduction" cql:"damage_reduction"`
	Thorns            float32 `json:"thorns" cql:"thorns"`
	LifeSteal         float32 `json:"lifeSteal" cql:"life_steal"`
	LifePerKill       float32 `json:"lifePerKill" cql:"life_per_kill"`
	GoldFind          float32 `json:"goldFind" cql:"gold_find"`
	MagicFind         float32 `json:"magicFind" cql:"magic_find"`
	LifeOnHit         float32 `json:"lifeOnHit" cql:"life_on_hit"`
	PrimaryResource   float32 `json:"primaryResource" cql:"primary_resource"`
	SecondaryResource float32 `json:"secondaryResource" cql:"secondary_resource"`
	LastUpdated       int64   `json:"last-updated" cql:"last_updated"`
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
	json.Unmarshal([]byte(b), &hd)
	return &hd, nil
}
