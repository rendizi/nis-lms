package handler

import (
	"lms/internal/server"
	"net/http"
)

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	server.Ok(map[string]interface{}{"message": "server is alive"}, w)
}
