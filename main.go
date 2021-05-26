package main

import (
	"crud-api/controllers"
	"crud-api/database"
	"crud-api/entity"
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
	router.HandleFunc("/create", controllers.CreateSave).Methods("POST")
	router.HandleFunc("/get", controllers.GetSave).Methods("GET")
	// router.HandleFunc("/get{id, password}", controllers.GetSaveByID).Methods("GET")
	router.HandleFunc("/update", controllers.UpdateSave).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletSave).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "prit",
			Password:   "",
			DB:         "pokefrogs",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Save{})
}
