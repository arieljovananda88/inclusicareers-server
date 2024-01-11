package test_controller

import (
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello test!"))

	// resMsg := "success"
	// lib.SendResponseMessage(w, resMsg, http.StatusCreated)
}
