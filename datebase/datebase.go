package datebase

import (
	"log"
	"os"

	"github.com/mergen1212/fiber-api-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	DB *gorm.DB
}

var Database DBInstance

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Faild connect DB")
		os.Exit(2)
	}

	log.Println("Connected DB")

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	// TODO: Add Migrations
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DBInstance{DB: db}

}
