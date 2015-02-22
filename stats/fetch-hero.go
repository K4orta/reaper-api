package stats

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Hero struct {
	Id         int        `json:"id"`
	Level      int        `json:"level"`
	Paragon    int        `json:"paragon"`
	Stats      *HeroStats `json:"stats"`
	Name       string     `json:"name"`
	LastUpdate int64      `json:"last-update"`
}

type HeroStats struct {
	Life              float32 `json:"life"`
	Damage            float32 `json:"damage"`
	Toughness         float32 `json:"toughness"`
	Healing           float32 `json:"healing"`
	AttackSpeed       float32 `json:"attackSpeed"`
	Armor             float32 `json:"armor"`
	Strength          float32 `json:"strength"`
	Dexterity         float32 `json:"dexterity"`
	Vitality          float32 `json:"vitality"`
	Intelligence      float32 `json:"intelligence"`
	PhysicalResist    float32 `json:"physicalResist"`
	FireResist        float32 `json:"fireResist"`
	ColdResist        float32 `json:"coldResist"`
	LightningResist   float32 `json:"lightningResist"`
	PoisonResist      float32 `json:"poisonResist"`
	ArcaneResist      float32 `json:"arcaneResist"`
	CritDamage        float32 `json:"critDamage"`
	BlockChance       float32 `json:"blockChance"`
	BlockAmountMin    float32 `json:"blockAmountMin"`
	BlockAmountMax    float32 `json:"blockAmountMax"`
	DamageIncrease    float32 `json:"damageIncrease"`
	CritChance        float32 `json:"critChance"`
	DamageReduction   float32 `json:"thorns"`
	Thorns            float32 `json:"damageReduction"`
	LifeSteal         float32 `json:"lifeSteal"`
	LifePerKill       float32 `json:"lifePerKill"`
	GoldFind          float32 `json:"goldFind"`
	MagicFind         float32 `json:"magicFind"`
	LifeOnHit         float32 `json:"lifeOnHit"`
	PrimaryResource   float32 `json:"primaryResource"`
	SecondaryResource float32 `json:"secondaryResource"`
	LastUpdate        int64   `json:"last-update"`
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
