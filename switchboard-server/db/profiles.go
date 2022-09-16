package db

import (
	"log"
	"switchboard-server/types"
)

func ListProfiles() (profiles []types.Profile, err error) {
	profiles = make([]types.Profile, 0)
	query := `
	SELECT * FROM profiles
	`
	rows, err := Db.Query(query)
	if err != nil {
		log.Fatalf("query all profiles: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id      int
			name    string
			enabled bool
		)
		if err := rows.Scan(&id, &name, &enabled); err != nil {
			log.Fatalf("scan the profiles: %v", err)
		} else {
			profiles = append(profiles, types.Profile{
				Id:      id,
				Name:    name,
				Enabled: enabled,
			})
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Scan profiles: %v", err)
	}
	return
}

func FindProfileByName(name string) (profile types.Profile, err error) {
	query := `
	SELECT * FROM profiles WHERE profiles.name = $1
	`
	err = Db.QueryRow(query, name).Scan(&profile.Id, &profile.Name, &profile.Enabled)
	return
}

func SwitchProfile(name string) (profile types.Profile, err error) {
	updateQuery := `
	UPDATE profiles SET enabled = FALSE WHERE profiles.enabled = TRUE;
	`
	if _, err = Db.Exec(updateQuery); err != nil {
		log.Println(err)
	}
	updateQuery = `
	UPDATE profiles SET enabled = TRUE WHERE profiles.name = $1;
	`
	if _, err = Db.Exec(updateQuery, name); err != nil {
		log.Println(err)
	}
	getQuery := `
	SELECT * FROM profiles WHERE profiles.enabled = TRUE;
	`
	err = Db.QueryRow(getQuery).Scan(&profile.Id, &profile.Name, &profile.Enabled)
	return
}

func CreateProfile(name string) (profile types.Profile, err error) {
	createQuery := `
	INSERT INTO profiles (name) VALUES ($1);
	`
	if _, err = Db.Exec(createQuery, name); err != nil {
		log.Println(err)
	}
	getQuery := `
	SELECT * FROM profiles ORDER BY id DESC LIMIT 1;
	`
	err = Db.QueryRow(getQuery).Scan(&profile.Id, &profile.Name, &profile.Enabled)
	if err != nil {
		log.Println(err)
	}
	return
}

func DeleteProfile(name string) (profiles []types.Profile, err error) {
	deleteQuery := `
	DELETE FROM profiles WHERE name = $1;
	`
	if _, err = Db.Exec(deleteQuery, name); err != nil {
		log.Println(err)
	}
	profiles, err = ListProfiles()
	return
}
