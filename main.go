package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/handlers"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/initializers"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/token"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot create token maker: %w", err)
		return
	}

	router := httprouter.New()
	initializers.InitializeServer(config, tokenMaker, router)
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	router := initializers.GetSRV().Router

	// Set CORS
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	router.NotFound = http.HandlerFunc(RouteNotFound)

	router.GET("/", Hello)                         // Public
	router.GET("/healthchecker", HealthChecker)    // Public
	router.POST("/login", handlers.LoginRoute)     // Public
	router.POST("/create", handlers.CreateOne)     // Public
	router.POST("/update/:id", handlers.UpdateOne) // Self or Admin
	router.POST("/delete/:id", handlers.DeleteOne) // Self or Admin
	router.GET("/self", handlers.GetSelf)          // Self only
	router.GET("/all", handlers.GetAll)            // Admin only
	router.GET("/get/:id", handlers.GetOne)        // Admin only

	log.Fatal(http.ListenAndServe(":"+config.ServerPort, router))
}

func Hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := `{` +
		`"info": "REST API Microservice for user handling",` +
		`"paths": {` +
		`"POST /create": "Create new User",` +
		`"POST /update": "Update User",` +
		`"POST /delete": "Delete User",` +
		`"GET /self": "Get Self User",` +
		`"GET /all": "Get All Users (Admin)",` +
		`"GET /get/:id": "Get User by ID (Admin)"` +
		`}` +
		`}`
	handlers.CreateResponse(w, response)
}

func HealthChecker(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := `{` +
		`"status": "success",` +
		`"message": "Welcome to Users Microservice health check!"` +
		`}`
	handlers.CreateResponse(w, response)
}

func RouteNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound) // StatusNotFound = 404

	data := `{"info": "The requested route was not found"}`

	handlers.CreateResponse(w, data)
}
