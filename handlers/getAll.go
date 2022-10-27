package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/initializers"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/models"
)

func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Check if admin (later)
	log.Println("Gettting all users as Admin...")

	// Get all users
	var rows []models.UserResponse
	initializers.DB.Raw("SELECT * FROM users").Scan(&rows)

	jsonUsers, _ := json.Marshal(rows)
	jsonString := `{"totalUsers": ` + strconv.Itoa(len(rows)) + `, "usersList": ` + string(jsonUsers) + `}`

	CreateResponse(w, string(jsonString))
}
