package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/initializers"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/models"
)

func GetOne(w http.ResponseWriter, requestData *http.Request, routerParams httprouter.Params) {
	// Check if admin (later)
	log.Println("Get one User as Admin...")

	// Get all users
	var rows []models.UserResponse
	initializers.DB.Raw(`SELECT * FROM users WHERE id = '` + routerParams.ByName("id") + `'`).Scan(&rows)
	jsonUser, _ := json.Marshal(rows[0])

	var message string
	if len(rows) == 0 {
		message = "User not found"
	} else {
		message = "User found"
	}

	jsonString := `{"message": "` + message + `", "userInfo": ` + string(jsonUser) + `}`

	CreateResponse(w, string(jsonString))
}
