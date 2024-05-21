package main

import (
	"fmt"
	"lms/db"
	"lms/internal/server/handler"
	"net/http"
	"os"

	"github.com/MadAppGang/httplog"
)

func main() {
	db.Init()
	mux := http.NewServeMux()

	mux.Handle("/health", httplog.Logger(http.HandlerFunc(handler.CheckHealth)))

	//Auth
	mux.Handle("/register", httplog.Logger(http.HandlerFunc(handler.Register)))
	mux.Handle("/login", httplog.Logger(http.HandlerFunc(handler.Login)))
	mux.Handle("GET /u/{username}", httplog.Logger(http.HandlerFunc(handler.GetProfile)))

	//Compile
	mux.Handle("/compile/cpp", httplog.Logger(http.HandlerFunc(handler.CompileCpp)))
	mux.Handle("/compile/python", httplog.Logger(http.HandlerFunc(handler.CompilePython)))

	//Tasks
	mux.Handle("GET /task/{id}", httplog.Logger(http.HandlerFunc(handler.GetTask)))
	mux.Handle("GET /search/task", httplog.Logger(http.HandlerFunc(handler.SearchTask)))
	mux.Handle("POST /task", httplog.Logger(http.HandlerFunc(handler.PostTask)))
	mux.Handle("GET /tests/{id}", httplog.Logger(http.HandlerFunc(handler.GetTests)))

	//Solutions
	mux.Handle("POST /task/{id}", httplog.Logger(http.HandlerFunc(handler.SubmitTask)))
	mux.Handle("GET /u/{username}/solutions", httplog.Logger(http.HandlerFunc(handler.SolvedByUser)))
	mux.Handle("GET /task/{id}/solutions", httplog.Logger(http.HandlerFunc(handler.TaskSolutions)))
	mux.Handle("GET /u/{username}/{id}", httplog.Logger(http.HandlerFunc(handler.DidHeSolvedIt)))

	//Classword
	mux.Handle("POST /classwork", httplog.Logger(http.HandlerFunc(handler.NewClassWork)))
	mux.Handle("GET /classwork/{id}", httplog.Logger(http.HandlerFunc(handler.GetStudents)))
	mux.Handle("GET /t/{username}/classwork", httplog.Logger(http.HandlerFunc(handler.ByTeacher)))
	mux.Handle("GET /u/{username}/classwork", httplog.Logger(http.HandlerFunc(handler.ForStudents)))

	//Search student
	mux.Handle("GET /search/user", httplog.Logger(http.HandlerFunc(handler.Search)))

	//Stats: teachers count, students count
	//Get all teachers and students

	//Delete teacher and student
	mux.Handle("DELETE /t/{username}", httplog.Logger(http.HandlerFunc(handler.DeleteTeacher)))
	mux.Handle("DELETE /u/{username}", httplog.Logger(http.HandlerFunc(handler.DeleteStudent)))

	//Rating: top 10 and your location.
	mux.Handle("GET /top", httplog.Logger(http.HandlerFunc(handler.Top)))
	mux.Handle("GET /top/{username}", httplog.Logger(http.HandlerFunc(handler.Place)))

	corsHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	err := http.ListenAndServe(":8080", corsHandler(mux))
	if err != nil {
		if err == http.ErrServerClosed {
			fmt.Println("server closed")
		} else {
			fmt.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	}
}
