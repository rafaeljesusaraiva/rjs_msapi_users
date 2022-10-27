package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/initializers"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/models"
)

func UpdateOne(w http.ResponseWriter, r *http.Request, routerParams httprouter.Params) {
	// Check if admin (later)
	log.Println("Update User as Admin or self...")

	// Sanitize input data
	decoder := json.NewDecoder(r.Body)
	var decodedResponse models.User
	err2 := decoder.Decode(&decodedResponse)
	if err2 != nil {
		panic(err2)
	}

	userInput := map[string]interface{}{}
	if decodedResponse.FirstName != "" {
		userInput["first_name"] = decodedResponse.FirstName
	}
	if decodedResponse.LastName != "" {
		userInput["last_name"] = decodedResponse.LastName
	}
	if decodedResponse.Username != "" {
		userInput["username"] = decodedResponse.Username
	}
	if decodedResponse.Email != "" {
		userInput["email"] = decodedResponse.Email
	}
	if decodedResponse.Password != "" {
		hashedPassword, _ := models.HashPassword(decodedResponse.Password)
		userInput["password"] = hashedPassword
	}
	if decodedResponse.ProfilePicture != "" {
		userInput["profile_picture"] = decodedResponse.ProfilePicture
	}

	// if Admin
	// if (r.FormValue("blockedAccount")) != "" {
	// 	blockedBool, _ := strconv.ParseBool(r.FormValue("blockedAccount"))
	// 	userInput["BlockedAccount"] = blockedBool
	// }
	// if (r.FormValue("confirmedAccount")) != "" {
	// 	confirmedBool, _ := strconv.ParseBool(r.FormValue("confirmedAccount"))
	// 	userInput["ConfirmedAccount"] = confirmedBool
	// }

	// Update user
	userParsedUUID, _ := uuid.Parse(routerParams.ByName("id"))

	// save new data to user with uuid
	result := initializers.DB.Model(&models.User{}).Where("id = ?", userParsedUUID).Updates(userInput)

	var jsonString string
	if result.RowsAffected == 0 {
		jsonString = `{"message": "No User updated"}`
	} else {
		jsonString = `{"message": "Updated User with success"}`
	}
	CreateResponse(w, jsonString)

}
