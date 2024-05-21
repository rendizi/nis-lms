package handler

import (
	"encoding/json"
	"lms/db/classwork_students"
	"lms/internal/server"
	"net/http"
	"strconv"
)

func NewClassWork(w http.ResponseWriter, r *http.Request) {
	login, role := server.GetLogin(w, r)
	if login == "" {
		return
	}
	if role != "teacher" {
		server.Error(map[string]interface{}{"message": "you dont have access to it"}, w)
		return
	}
	var ClassWordData classwork_students.ClassWork
	err := json.NewDecoder(r.Body).Decode(&ClassWordData)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	id, err := ClassWordData.Insert(login)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	server.Ok(map[string]interface{}{"message": "added successfully", "id": id}, w)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	students, err := classwork_students.GetStudentInfo(intId)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	resp, err := json.Marshal(students)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func ForStudents(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
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
	classwork, err := classwork_students.ForStudent(username, intPage, intPageSize)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	resp, err := json.Marshal(classwork)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func ByTeacher(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
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
	classword, err := classwork_students.GetByTeacher(username, intPage, intPageSize)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	resp, err := json.Marshal(classword)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
