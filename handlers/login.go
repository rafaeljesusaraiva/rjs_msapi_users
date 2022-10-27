package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/initializers"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/models"
)

func LoginRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var decodedResponse models.LoginUserRequest
	err1 := json.NewDecoder(r.Body).Decode(&decodedResponse)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Aqui", r.Body)

	var req models.LoginUserRequest
	req.Username = decodedResponse.Username
	req.Password = decodedResponse.Password

	var user models.User
	request := initializers.DB.Where("username = ?", req.Username).Or("email = ?", req.Username).First(&user)
	err := request.Error
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			jsonString, _ := json.Marshal(map[string]string{"message": "Login Error - SQL", "details": err.Error()})
			CreateResponse(w, string(jsonString))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		jsonString, _ := json.Marshal(map[string]string{"message": "Login Error - Internal Server", "details": err.Error()})
		CreateResponse(w, string(jsonString))
		return
	}

	err = models.CheckPassword(req.Password, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		jsonString, _ := json.Marshal(map[string]string{"message": "Login Error - Unauthorized", "details": err.Error()})
		CreateResponse(w, string(jsonString))
		return
	}

	accessToken, err := initializers.SRV.TokenMaker.CreateToken(
		user.Username,
		initializers.SRV.Config.AccessTokenDuration,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonString, _ := json.Marshal(map[string]string{"message": "Login Error - Internal Server", "details": err.Error()})
		CreateResponse(w, string(jsonString))
		return
	}

	rsp := models.LoginUserResponse{
		AccessToken: accessToken,
		User:        models.NewUserResponse(user),
	}

	jsonRsp, _ := json.Marshal(rsp)
	w.WriteHeader(http.StatusOK)
	jsonString, _ := json.Marshal(map[string]string{"message": "Login Success", "details": string(jsonRsp)})
	CreateResponse(w, string(jsonString))
}
