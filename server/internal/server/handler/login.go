package handler

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	db "lms/db/students"
	"lms/internal/server"
	"net/http"
	"time"
)

var jwtSecret = []byte("nis sucks")

func Login(w http.ResponseWriter, r *http.Request) {
	var creds registerRequest

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		server.Error(map[string]interface{}{"message": "invalid request"}, w)
		return
	}

	var student db.Student
	student.Login = creds.Login
	student.Password = creds.Password

	role := "student"

	isTeacher := r.URL.Query().Get("teacher")
	if isTeacher == "1" {
		err = student.ValidateTeacher()
		if err != nil {
			server.Error(map[string]interface{}{"message": err.Error()}, w)
			return
		}
		role = "teacher"
	} else {

		err = student.Validate()
		if err != nil {
			server.Error(map[string]interface{}{"message": err.Error()}, w)
			return
		}
	}

	if student.Login == "ren" {
		role = "admin"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": creds.Login,
		"role":  role,
		"exp":   time.Now().Add(3 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		server.Error(map[string]interface{}{"message": "Error generating token", "status": 400}, w)
		return
	}

	server.Ok(map[string]interface{}{"message": "success", "token": tokenString}, w)
}
