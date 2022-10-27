package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateResponse(w http.ResponseWriter, input string) {
	res := json.RawMessage(input)

	response, _ := json.MarshalIndent(res, "", "    ")
	fmt.Fprint(w, string(response))
}
