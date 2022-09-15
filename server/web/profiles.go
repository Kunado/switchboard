package web

import (
	"encoding/json"
	"net/http"
	"switchboard/db"
)

func HandleProfiles(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = listProfiles(w, r)
	case "POST":
		err = createProfile(w, r)
	case "DELETE":
		err = deleteProfile(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func listProfiles(w http.ResponseWriter, r *http.Request) (err error) {
	profiles, err := db.ListProfiles()
	if err != nil {
		return
	}
	output, err := json.Marshal(&profiles)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func createProfile(w http.ResponseWriter, r *http.Request) (err error) {
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	var profileName db.ProfileName
	json.Unmarshal(body, &profileName)
	profile, err := db.CreateProfile(profileName.Name)
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

func deleteProfile(w http.ResponseWriter, r *http.Request) (err error) {
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	var profileName db.ProfileName
	json.Unmarshal(body, &profileName)
	profiles, err := db.DeleteProfile(profileName.Name)
	if err != nil {
		return
	}
	output, err := json.Marshal(&profiles)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
