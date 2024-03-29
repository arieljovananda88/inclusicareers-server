package user_controller

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/m/src/internal/database"
	"example.com/m/src/internal/lib"
	"example.com/m/src/internal/models"
)

type GetAllUserResponse struct {
	lib.ResponseMessage
	Data []models.Users
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	var users []models.Users

	// Find the user in the database
	db := database.GetInstance()

	result := db.Find(&users)
	if result.Error != nil {
		log.Println(result.Error.Error())

		resMsg := "invalid user id"
		lib.SendResponseMessage(w, resMsg, http.StatusNotFound)
		return
	}

	msg := GetAllUserResponse{
		ResponseMessage: lib.ResponseMessage{
			Message: "success",
		},
		Data: users,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}
