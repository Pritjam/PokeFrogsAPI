package controllers

import (
	"crud-api/database"
	"crud-api/entity"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAllSave get all save data
//TODO: make something i'm happy with here
func GetAllSave(w http.ResponseWriter, r *http.Request) {
	var saves []entity.Save
	database.Connector.Find(&saves)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(saves)
}

//GetSaveByID returns save with specific ID
func GetSave(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var idpass entity.IDPass
	json.Unmarshal(requestBody, &idpass)

	var save entity.Save
	database.Connector.First(&save, idpass.ID)
	if idpass.Password == save.Password {
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

	database.Connector.Create(save)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(save)
}

//UpdateSaveByID updates save with respective ID
func UpdateSave(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var save, oldSave entity.Save
	json.Unmarshal(requestBody, &save)
	database.Connector.First(&oldSave, save.ID)
	if oldSave.Password == save.Password {
		database.Connector.Save(&save)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(save)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

//DeletSaveByID delete's save with specific ID
func DeletSaveByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ownerid"]

	var save entity.Save
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("ownerid = ?", id).Delete(&save)
	w.WriteHeader(http.StatusNoContent)
}
