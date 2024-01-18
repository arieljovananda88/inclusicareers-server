package userdata_controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/src/internal/database"
	"example.com/m/src/internal/lib"
	"example.com/m/src/internal/models"
)

type GetAllUserResponse struct {
	lib.ResponseMessage
	Data []models.UserData
}

func CreateUserData(userId string) (err error) {

	// Find the user in the database
	db := database.GetInstance()

	result := db.Create(&models.UserData{
		UserID: userId,
	})

	if result.Error != nil {
		resMsg := fmt.Errorf("Error while creating userData")
		return resMsg
	}

	return nil
}

func GetAllUserData(w http.ResponseWriter, r *http.Request) {
	var userDatas []models.UserData

	// Find the user in the database
	db := database.GetInstance()

	result := db.Find(&userDatas)
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
		Data: userDatas,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}
