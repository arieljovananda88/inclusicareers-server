package routes

import (
	"log"

	test_controller "example.com/m/src/internal/controllers/test"
	user_controller "example.com/m/src/internal/controllers/user"
	auth_controller "example.com/m/src/internal/controllers/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func InitializeRoutes() *chi.Mux {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// Semua Routes bakal ditambah disini
	// Ditambah dicontroller dulu method nya apa terus tambah sini

	//Test Route
	r.Get("/", test_controller.Test)

	// User routes
	r.Post("/user", user_controller.CreateUser)
	r.Get("/user", user_controller.GetAllUser)

	// Auth routes
	r.Post("/login", auth_controller.Login)

	return r

}
