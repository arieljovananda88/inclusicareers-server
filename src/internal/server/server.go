package server

import (
	"net/http"

	"example.com/m/src/internal/routes"
)

func Run() {

	// Gausah di otak atik yang ini
	// Nanti semua routing seharusnya ke initialize di Routes.go
	// File ini cuman buat dipanggil di cmd/server
	r := routes.InitializeRoutes()
	http.ListenAndServe(":8080", r)
}
