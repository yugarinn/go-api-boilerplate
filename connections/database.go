package connections

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/yugarinn/go-api-boilerplate/utils"
)

const projectDirName = "go-api-boilerplate"

func Database() *gorm.DB {
	utils.LoadEnvFile(os.Getenv("BOILERPLATE_ENV"))

	databaseURI := os.Getenv("DB_URI")
	database, err := gorm.Open(mysql.Open(databaseURI), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	return database
}
