package models

type Response struct {
	Status     string `json:"status"`
	StatucCode int    `json:"statusCode"`
	Data       []User `json:"data"`
}
