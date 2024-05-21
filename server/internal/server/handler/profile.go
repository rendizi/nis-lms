package handler

import (
	"encoding/json"
	"lms/db/teachers"
	"net/http"
	"strconv"

	db "lms/db/students"
	"lms/internal/server"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username") // Assuming you get the username from the query params
	if username == "" {
		server.Error(map[string]interface{}{"message": "username is required"}, w)
		return
	}

	var user db.Student
	user.Login = username
	err := user.GetInfo()
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}

	rank, err := db.Place(username)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}

	user.Stats.Rank = rank

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
}

func Search(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	page := r.URL.Query().Get("page")
	pagesize := r.URL.Query().Get("pagesize")
	var intPage, intPageSize int
	var err error
	if page == "" {
		intPage = 1
	} else {
		intPage, err = strconv.Atoi(page)
		if err != nil {
			server.Error(map[string]interface{}{"message": "page is not int"}, w)
			return
		}
	}
	if pagesize == "" {
		intPageSize = 10
	} else {
		intPageSize, err = strconv.Atoi(pagesize)
		if err != nil {
			server.Error(map[string]interface{}{"message": "page size is not int"}, w)
			return
		}
	}
	users, err := db.Search(username, intPage, intPageSize)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
}

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	_, role := server.GetLogin(w, r)
	if role == "" {
		return
	}
	if role != "admin" {
		server.Error(map[string]interface{}{"message": "you can't delete"}, w)
		return
	}
	err := teachers.Delete(username)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	server.Ok(map[string]interface{}{"message": "success"}, w)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	_, role := server.GetLogin(w, r)
	if role == "" {
		return
	}
	if role != "admin" {
		server.Error(map[string]interface{}{"message": "you can't delete"}, w)
		return
	}
	err := db.DeleteStudent(username)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	server.Ok(map[string]interface{}{"message": "success"}, w)
}

func Top(w http.ResponseWriter, r *http.Request) {
	top, err := db.GetTop10ByRating()
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	resp, err := json.Marshal(top)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func Place(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	rank, err := db.Place(username)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	server.Ok(map[string]interface{}{"rank": rank}, w)
}
