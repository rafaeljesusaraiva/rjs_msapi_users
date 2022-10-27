package handlers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/initializers"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/models"
)

func DeleteOne(w http.ResponseWriter, r *http.Request, routerParams httprouter.Params) {
	// Check if admin (later)
	log.Println("Delete one User as Admin")

	// Nullify all user data
	userNulled := models.User{
		LastName:                 "--Blocked Account",
		Password:                 "-",
		ProfilePicture:           "-",
		BlockedAccount:           true,
		AccountConfirmationToken: "-",
		ResetPasswordToken:       "-",
	}

	userParsedUUID, _ := uuid.Parse(routerParams.ByName("id"))
	result := initializers.DB.Model(&models.User{}).Where("id = ?", userParsedUUID).Updates(userNulled)

	var jsonString string
	if result.RowsAffected == 0 {
		jsonString = `{"message": "No User deleted"}`
	} else {
		jsonString = `{"message": "Deleted User with success"}`
	}
	CreateResponse(w, jsonString)
}
