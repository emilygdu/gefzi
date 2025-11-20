package database

import (
	"log"

	sqlite "github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("gefzi.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Fehler bei der Verbindung mit der Datenbank", err)
	}
	log.Println("Alles supi die Verbindung zur Datenbank steht")
}
