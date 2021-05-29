package controllers

import (
	"crud-api/database"
	"crud-api/entity"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"unicode"

	"github.com/gorilla/mux"
)

func isAlphabetic(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

//GetAllSave get all save data
//TODO: make something i'm happy with here
func GetAllSave(w http.ResponseWriter, r *http.Request) {
	var saves []entity.Save
	database.Connector.Find(&saves)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(saves)
}

//GetOther returns a string in other_storage, used for quests, market, and more
func GetOther(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	label := vars["lbl"]
	if !isAlphabetic(label) {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var other_storage entity.OtherStorage
	database.Connector.Select("content").Where("label = ?", label).First(&other_storage)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(other_storage)
}

//GetSaveByID returns save with specific ID
func GetSave(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var credentials entity.Credentials
	json.Unmarshal(requestBody, &credentials)

	if !isAlphabetic(credentials.Username) {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var save entity.Save
	database.Connector.Where("username = ?", credentials.Username).First(&save)
	if credentials.Password == save.Password {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(save)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

//CreateSave creates save
func CreateSave(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var save entity.Save
	json.Unmarshal(requestBody, &save)

	database.Connector.Select("username", "password", "save_data").Create(&save)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(save)
}

//UpdateSaveByID updates save with respective ID
func UpdateSave(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var oldSave, save entity.Save
	json.Unmarshal(requestBody, &save)
	database.Connector.Where("username = ?", save.Username).First(&oldSave)
	if oldSave.Password == save.Password {
		database.Connector.Model(&oldSave).Select("username", "password", "save_data").Updates(&save)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(save)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

//DeleteSaveByID delete's save with specific ID
func DeletSave(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var credentials entity.Credentials
	json.Unmarshal(requestBody, &credentials)
	var save entity.Save
	database.Connector.Where("username = ?", credentials.Username).First(&save)

	if credentials.Password == save.Password {
		database.Connector.Where("username = ?", save.Username).Delete(&save)
	}
	w.WriteHeader(http.StatusNoContent)
}
