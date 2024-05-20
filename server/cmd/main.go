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

	//Compile
	mux.Handle("/compile/cpp", httplog.Logger(http.HandlerFunc(handler.CompileCpp)))
	mux.Handle("/compile/python", httplog.Logger(http.HandlerFunc(handler.CompilePython)))

	//Tasks
	mux.Handle("GET /task/{id}", httplog.Logger(http.HandlerFunc(handler.GetTask)))
	mux.Handle("GET /search/task", httplog.Logger(http.HandlerFunc(handler.SearchTask)))
	mux.Handle("POST /task", httplog.Logger(http.HandlerFunc(handler.PostTask)))

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
