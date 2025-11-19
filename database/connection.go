package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
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
