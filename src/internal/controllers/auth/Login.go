package auth_controller

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
	"fmt"

	"example.com/m/src/internal/database"
	"example.com/m/src/internal/lib"
	"example.com/m/src/internal/models"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	lib.ResponseMessage
	Token	string
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	var user models.User

	json.NewDecoder(r.Body).Decode(&loginReq)

	// Check for invalid request body
	if loginReq.Email == "" || loginReq.Password == "" {
		resMsg := "invalid request"
		lib.SendResponseMessage(w, resMsg, http.StatusBadRequest)
		return
	}

	database.GetInstance().First(&user, "email = ?", loginReq.Email)

	if user.UserID == "" {
		lib.SendResponseMessage(w, "email not found", http.StatusNotFound)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))

	if err != nil {
		lib.SendResponseMessage(w, "wrong password", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		fmt.Print(err)
		lib.SendResponseMessage(w, "signing error", http.StatusInternalServerError)
		return
	}

	msg := LoginResponse{
		ResponseMessage: lib.ResponseMessage{
			Message: "success",
		},
		Token: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}