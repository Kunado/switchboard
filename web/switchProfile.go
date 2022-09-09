package web

import (
	"encoding/json"
	"net/http"
	"switchboard/db"
)

func HandleSwitchProfile(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "PUT":
		err = switchProfile(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func switchProfile(w http.ResponseWriter, r *http.Request) (err error) {
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	var profileName db.ProfileName
	json.Unmarshal(body, &profileName)
	profile, err := db.SwitchProfile(profileName.Name)
	if err != nil {
		return
	}
	output, err := json.Marshal(&profile)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
