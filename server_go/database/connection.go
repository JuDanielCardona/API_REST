package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	//host := "localhost"
	host := os.Getenv("DATABASE_URL")
	var DSN = "host=" + host + " user=admin password=12345 dbname=users_db port=5432"
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Database is connected.")
		log.Println("=> " + DSN)
	}
}
