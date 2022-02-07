package connections

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseURI := os.Getenv("DB_URI")

	db, err := gorm.Open(mysql.Open(databaseURI), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database")
	}

	return db
}
