package main

import (
	"example.com/m/src/internal/database"
	"example.com/m/src/internal/models"
)

func main() {
	database.GetInstance().AutoMigrate(&models.User{})
}
