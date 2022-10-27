package main

import (
	"fmt"
	"log"

	"github.com/rafaeljesusaraiva/rjs_msapi_users/initializers"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Raw("DROP EXTENSION \"uuid-ossp\";CREATE EXTENSION \"uuid-ossp\";")
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
