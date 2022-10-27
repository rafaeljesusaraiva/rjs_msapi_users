package initializers

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var connectionRetries = 1

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Lisbon", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	for err != nil {
		log.Fatalf("Failed to connect to database (x%d), retrying in 5 seconds...", connectionRetries)
		connectionRetries++
		time.Sleep(5 * time.Second)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			continue
		}
	}

	if err == nil {
		log.Println("Connected Successfully to the Database")
	}

}
