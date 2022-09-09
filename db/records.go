package db

import "log"

type Record struct {
	Id        int
	Host      string
	Value     string
	ProfileId int
}

type RecordBuilder struct {
	Host        string `json:"host"`
	Value       string `json:"value"`
	ProfileName string `json:"profile_name"`
}

type RecordValue struct {
	Value string `json:"value"`
}

func ListRecords() (records []Record, err error) {
	query := `
	SELECT cname_records.* FROM profiles
	LEFT JOIN cname_records ON profiles.id = cname_records.profile_id
	WHERE profiles.enabled = TRUE
	`
	rows, err := Db.Query(query)
	if err != nil {
		log.Fatalf("query all cname_records: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id, profileId int
			host, value   string
		)
		if err := rows.Scan(&id, &host, &value, &profileId); err != nil {
			log.Fatalf("scan the cname_record: %v", err)
		}
		records = append(records, Record{
			Id:        id,
			Host:      host,
			Value:     value,
			ProfileId: profileId,
		})
	}

	if err := rows.Close(); err != nil {
		log.Fatalf("rows close: %v", err)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Scan records: %v", err)
	}
	return
}

func FindRecordByValue(value string) (record Record, err error) {
	query := `
	SELECT cname_records.* FROM cname_records
	LEFT JOIN profiles ON profiles.id = cname_records.profile_id
	WHERE (
		cname_records.value = $1 AND
		profiles.enabled = TRUE
	)
	`
	err = Db.QueryRow(query, value).Scan(&record.Id, &record.Host, &record.Value, &record.ProfileId)
	return
}

func CreateRecord(recordBuilder RecordBuilder) (record Record, err error) {
	profile, err := FindProfileByName(recordBuilder.ProfileName)
	if err != nil {
		log.Println(err)
	}
	createQuery := `
	INSERT INTO cname_records (host, value, profile_id) VALUES ($1, $2, $3);
	`
	if _, err = Db.Exec(createQuery, recordBuilder.Host, recordBuilder.Value, profile.Id); err != nil {
		log.Println(err)
	}
	getQuery := `
	SELECT * FROM cname_records ORDER BY id DESC LIMIT 1;
	`
	err = Db.QueryRow(getQuery).Scan(&record.Id, &record.Host, &record.Value, &record.ProfileId)
	if err != nil {
		log.Println(err)
	}
	return
}

func DeleteRecord(name string) (records []Record, err error) {
	deleteQuery := `
	DELETE FROM cname_records WHERE value = $1;
	`
	if _, err = Db.Exec(deleteQuery, name); err != nil {
		log.Println(err)
	}
	records, err = ListRecords()
	return
}
