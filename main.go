package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/handlers"
)

func main() {
	router := httprouter.New()

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

	router.GET("/", Hello)                     // Public
	router.POST("/create", handlers.CreateOne) // Public
	router.POST("/update", handlers.UpdateOne) // Self or Admin
	router.POST("/delete", handlers.DeleteOne) // Self or Admin
	router.GET("/self", handlers.GetSelf)      // Self only
	router.GET("/all", handlers.GetAll)        // Admin only
	router.GET("/get/:id", handlers.GetOne)    // Admin only

	log.Fatal(http.ListenAndServe(":8080", router))
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
	handlers.CreateResponse(w, "success", 200, response)
}

func RouteNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound) // StatusNotFound = 404

	data := "The requested page was not found"

	handlers.CreateResponse(w, "notFound", 404, data)
}
