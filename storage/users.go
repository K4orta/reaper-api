package storage

//create table users (battletag text NOT NULL PRIMARY KEY, last_updated bigint NOT NULL, added timestamp NOT NULL);
import (
	"github.com/k4orta/reaper-api/stats"
	"log"
	"time"
)

func InsertUser(battleTag string) error {
	s, err := CreateConnection()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer s.Close()

	props := []string{
		"battletag",
		"last_updated",
		"added",
	}

	h := stats.User{
		battleTag,
		0,
		time.Now(),
	}

	_, err = s.NamedExec(`INSERT INTO users (`+serializeProps("", props)+`) 
						SELECT `+serializeProps(":", props)+` 
						WHERE NOT EXISTS (SELECT 1 FROM users WHERE battletag=:battletag);`, h)
	return err
}

func FetchAllUsers() ([]*stats.User, error) {
	s, err := CreateConnection()
	users := []*stats.User{}
	if err != nil {
		log.Fatal(err)
		return users, err
	}
	defer s.Close()
	rows, err := s.Queryx(`SELECT * FROM users;`)
	usr := stats.User{}
	for rows.Next() {
		rows.StructScan(&usr)
		users = append(users, &stats.User{
			usr.BattleTag,
			usr.LastUpdated,
			usr.Added,
		})
	}
	return users, nil
}

func UpdateUser(user *stats.User) error {
	s, err := CreateConnection()
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = s.NamedExec(`UPDATE users SET (last_updated) = (:last_updated) WHERE battletag=:battletag;`, user)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
