package lib

import (
	"encoding/json"
	"net/http"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

func SendResponseMessage(w http.ResponseWriter, msg string, status int) {
	message := ResponseMessage{
		Message: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}
