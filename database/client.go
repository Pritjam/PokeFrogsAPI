package database

import (
	"crud-api/structures"
	"log"

	"github.com/jinzhu/gorm"
)

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

//Migrate create/updates database table
func Migrate(saveTable *structures.Save, otherTable *structures.OtherStorage) {
	Connector.AutoMigrate(&saveTable, &otherTable)
	log.Println("Table migrated")
}
