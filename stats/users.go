package stats

import (
	"time"
)

type User struct {
	BattleTag   string    `json:"battleTag" db:"battletag"`
	LastUpdated int64     `json:"lastUpdated" db:"last_updated"`
	Added       time.Time `json:"added" db:"added"`
}
