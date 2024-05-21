package handler

import (
	"encoding/json"
	db "lms/db/students"
	"lms/db/tasks"
	"lms/internal/execute"
	"lms/internal/server"
	"net/http"
	"strconv"
)

func SubmitTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	login, _ := server.GetLogin(w, r)
	if login == "" {
		return
	}
	tests, err := tasks.GetTests(intId)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	var code compileRequest
	err = json.NewDecoder(r.Body).Decode(&code)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	passed, total, err := execute.RunTests(tests, code.Code, 1)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	if passed == total {
		err = tasks.Submit(intId, login, code.Code)
		if err != nil {
			server.Error(map[string]interface{}{"message": err.Error()}, w)
			return
		}
		task, err := tasks.GetTask(id)
		if err != nil {
			server.Error(map[string]interface{}{"message": err.Error()}, w)
			return
		}
		var student db.Student
		student.Login = login
		err = student.Solved(task.Difficulty)
		if err != nil {
			server.Error(map[string]interface{}{"message": err.Error()}, w)
			return
		}
	}
	server.Ok(map[string]interface{}{"message": "success", "id": id, "total": total, "passed": passed}, w)

}

func SolvedByUser(w http.ResponseWriter, r *http.Request) {
	login := r.PathValue("username")
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
	solutions, err := tasks.Solutions(login, intPage, intPageSize)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	resp, err := json.Marshal(solutions)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func TaskSolutions(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	page := r.URL.Query().Get("page")
	pagesize := r.URL.Query().Get("pagesize")
	var intPage, intPageSize int
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
	solutions, err := tasks.WhoSolvedIt(intId, intPage, intPageSize)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	resp, err := json.Marshal(solutions)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func DidHeSolvedIt(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	login := r.PathValue("username")
	did, err := tasks.DidSolvedIt(login, intId)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	server.Ok(map[string]interface{}{"message": "success", "did": did}, w)
}
