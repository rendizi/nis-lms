package handler

import (
	"encoding/json"
	"fmt"
	db "lms/db/students"
	"lms/internal/decoder"
	"lms/internal/reqs"
	"lms/internal/server"
	"net/http"
	"strconv"
)

type registerRequest struct {
	Login          string `json:"login"`
	Password       string `json:"password"`
	Username       string `json:"username"`
	SystemPassword string `json:"system_password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var creds registerRequest

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		server.Error(map[string]interface{}{"message": "invalid request"}, w)
		return
	}

	token, err := reqs.Login(creds.Login, creds.Password)
	if err != nil {
		server.Error(map[string]interface{}{"message": "login failed"}, w)
		return
	}
	info, err := decoder.DecodeJWT(token)
	if err != nil {
		server.Error(map[string]interface{}{"message": "invalid token"}, w)
	}
	fmt.Println(info["Email"].(string))
	email := info["Email"].(string)
	klass, school, err := reqs.AdditionalInfo(token)
	if err != nil {
		server.Error(map[string]interface{}{"message": "invalid token"}, w)
		return
	}
	parallel, err := strconv.Atoi(klass[:len(klass)-1])
	if err != nil {
		server.Error(map[string]interface{}{"message": "invalid token"}, w)
		return
	}

	var student db.Student
	student.Login = creds.Username
	student.Password = creds.SystemPassword
	student.Email = email
	student.Klass = klass
	student.Parallel = parallel
	student.School = school
	err = student.Insert()
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}

	server.Ok(map[string]interface{}{"message": "success", "email": email, "klass": klass, "school": school, "parallel": parallel}, w)
}
