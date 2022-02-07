package main

import (
	"crud-api/api"
	"crud-api/database"
	"crud-api/structures"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", api.CreateSave).Methods("POST")
	router.HandleFunc("/get", api.GetSave).Methods("GET")
	router.HandleFunc("/get/all", api.GetAllSave).Methods("GET")
	router.HandleFunc("/get/{lbl}", api.GetOther).Methods("GET")
	router.HandleFunc("/update", api.UpdateSave).Methods("PUT")
	router.HandleFunc("/delete", api.DeleteSave).Methods("DELETE")
}

func initDB() {
	connectionString := database.GetConnectionString()
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&structures.Save{}, &structures.OtherStorage{})
}
