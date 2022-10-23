package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateResponse(w http.ResponseWriter, status string, statusCode int, input string) {
	w.Header().Set("Content-Type", "application/json")

	data := map[string]interface{}{
		"status":     status,
		"statusCode": statusCode,
		"data":       input,
	}

	response, _ := json.MarshalIndent(data, "", "    ")
	fmt.Fprint(w, string(response))
}
