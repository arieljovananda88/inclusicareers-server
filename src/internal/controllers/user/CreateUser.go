package user_controller

import (
	"net/http"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))

	// resMsg := "success"
	// lib.SendResponseMessage(w, resMsg, http.StatusCreated)
}
