package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateResponse(w http.ResponseWriter, status string, statusCode int, input string) {
	w.Header().Set("Content-Type", "application/json")

	res := json.RawMessage(input)

	response, _ := json.MarshalIndent(res, "", "    ")
	fmt.Fprint(w, string(response))
}
