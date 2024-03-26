package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	host := os.Getenv("DATABASE_URL")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")

	var DSN = "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Database is connected.")
		log.Println("=> " + DSN)
	}
}
