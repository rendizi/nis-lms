package handler

import (
	"encoding/json"
	"fmt"
	"lms/db/tasks"
	"lms/internal/server"
	"net/http"
	"strconv"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	task, err := tasks.GetTask(id)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}

	fmt.Println(task.SecondExample)

	server.Ok(map[string]interface{}{"title": task.Title, "id": task.Id, "author": task.Author, "description": task.Description,
		"difficulty": task.Difficulty, "example": []string{task.FirstExample, task.SecondExample, task.ThirdExample},
		"image": task.Image}, w)
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	var task tasks.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	id, err := task.Post()
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	server.Ok(map[string]interface{}{"message": "success", "id": id}, w)
}

func SearchTask(w http.ResponseWriter, r *http.Request) {
	difficulty := r.URL.Query().Get("difficulty")
	title := r.URL.Query().Get("title")
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
	tasks, err := tasks.GetTasksWithPagination(intPage, intPageSize, difficulty, title)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
	}
	response, err := json.Marshal(tasks)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}

	// Set the response header content-type to application/json
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func GetTests(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	tests, err := tasks.GetTests(intId)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	response, err := json.Marshal(tests)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
