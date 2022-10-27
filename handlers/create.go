package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/initializers"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/models"
)

type CreateUser struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

func newUserResponse(user models.User) models.UserResponse {
	return models.UserResponse{
		Id:             user.Id,
		Username:       user.Username,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		CreatedAt:      user.CreatedAt,
		ProfilePicture: user.ProfilePicture,
	}
}

func CreateOne(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("Creating User...")

	// Decode the request body into a new `User` instance
	decoder := json.NewDecoder(r.Body)
	var decodedResponse CreateUser
	err2 := decoder.Decode(&decodedResponse)
	if err2 != nil {
		panic(err2)
	}

	// Verify decoded response user data
	if decodedResponse.FirstName == "" {
		log.Println("First Name is required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if decodedResponse.LastName == "" {
		log.Println("Last Name is required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if decodedResponse.Email == "" {
		log.Println("Email is required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if decodedResponse.Password == "" {
		log.Println("Password is required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// If no username, generate from first and last name with random generated number
	if decodedResponse.Username == "" {
		decodedResponse.Username = decodedResponse.FirstName + decodedResponse.LastName + fmt.Sprintf("%02d", rand.Intn(99-1)+1)
	}

	hashedPassword, _ := models.HashPassword(decodedResponse.Password)

	userInstance := models.User{
		FirstName: decodedResponse.FirstName,
		LastName:  decodedResponse.LastName,
		Username:  decodedResponse.Username,
		Email:     decodedResponse.Email,
		Password:  hashedPassword,
	}

	result := initializers.DB.Create(&userInstance)
	if result.Error != nil {
		log.Println("Error creating User:", result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		// create json string with message and error
		jsonString, _ := json.Marshal(map[string]string{"message": "Error creating User", "details": result.Error.Error()})
		CreateResponse(w, string(jsonString))
		return
	}

	w.WriteHeader(http.StatusCreated)

	rsp := newUserResponse(userInstance)

	jsonUser, _ := json.Marshal(rsp)
	jsonString := `{"message": "Created User with success", "details": ` + string(jsonUser) + `}`

	CreateResponse(w, jsonString)
}
