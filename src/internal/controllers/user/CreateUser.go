package user_controller

import (
	"encoding/json"
	"net/http"

	"example.com/m/src/internal/database"
	"example.com/m/src/internal/lib"
	"example.com/m/src/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello crr!"))
	var userReq CreateUserRequest

	json.NewDecoder(r.Body).Decode(&userReq)

	// Check for invalid request body
	if userReq.Email == "" || userReq.Password == "" || userReq.Username == "" || userReq.FullName == "" {
		resMsg := "invalid request"
		lib.SendResponseMessage(w, resMsg, http.StatusBadRequest)
		return
	}

	// Hash the user's password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 14)
	if err != nil {
		lib.SendResponseMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save the newly created user
	db := database.GetInstance()
	result := db.Create(&models.User{
		FullName: userReq.FullName,
		Username: userReq.Username,
		Password: string(passwordHash),
		Email:    userReq.Email,
	})

	if result.Error != nil {
		resMsg := "Error"
		lib.SendResponseMessage(w, resMsg, http.StatusBadRequest)
		return
	}

	resMsg := "success"
	lib.SendResponseMessage(w, resMsg, http.StatusCreated)
}