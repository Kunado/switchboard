package web

import (
	"encoding/json"
	"net/http"
	"switchboard-server/db"
	"switchboard-server/types"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func HandleRecords(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = listRecords(w, r)
	case "POST":
		err = createRecord(w, r)
	case "DELETE":
		err = deleteRecord(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func listRecords(w http.ResponseWriter, r *http.Request) (err error) {
	records, err := db.ListRecords()
	if err != nil {
		return
	}
	output, err := json.Marshal(&records)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func createRecord(w http.ResponseWriter, r *http.Request) (err error) {
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	var recordBuilder types.RecordBuilder
	json.Unmarshal(body, &recordBuilder)
	record, err := db.CreateRecord(recordBuilder)
	if err != nil {
		return
	}
	output, err := json.Marshal(&record)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func deleteRecord(w http.ResponseWriter, r *http.Request) (err error) {
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	var recordValue types.RecordValue
	json.Unmarshal(body, &recordValue)
	records, err := db.DeleteRecord(recordValue.Value)
	if err != nil {
		return
	}
	output, err := json.Marshal(&records)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
