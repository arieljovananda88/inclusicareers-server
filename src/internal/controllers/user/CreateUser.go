package user_controller

import (
	"encoding/json"
	"net/http"

	userdata_controller "example.com/m/src/internal/controllers/userData"
	"example.com/m/src/internal/database"
	"example.com/m/src/internal/lib"
	"example.com/m/src/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type NewUser struct {
	FullName string
	Username string
	Password string
	Email    string
}

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
	newUser := models.Users{FullName: userReq.FullName, Username: userReq.Username, Password: string(passwordHash), Email: userReq.Email}
	result := db.Create(&newUser)

	if result.Error != nil {
		resMsg := "Error while creating user"
		lib.SendResponseMessage(w, resMsg, http.StatusBadRequest)
		return
	}

	res := userdata_controller.CreateUserData(string(newUser.UserID))

	if res != nil {
		resMsg := "Error while creating user data"
		lib.SendResponseMessage(w, resMsg, http.StatusBadRequest)
		return
	}

	resMsg := "success"
	lib.SendResponseMessage(w, resMsg, http.StatusCreated)
}
