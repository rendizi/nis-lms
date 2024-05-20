package handler

import (
	"encoding/json"
	"lms/internal/compile"
	"lms/internal/server"
	"net/http"
)

type compileRequest struct {
	Code string `json:"code"`
}

func CompileCpp(w http.ResponseWriter, r *http.Request) {
	var code compileRequest
	err := json.NewDecoder(r.Body).Decode(&code)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	result, err := compile.ExecuteCppCode(code.Code)
	server.Ok(map[string]interface{}{"message": result, "error": err.Error()}, w)
}

func CompilePython(w http.ResponseWriter, r *http.Request) {
	var code compileRequest
	err := json.NewDecoder(r.Body).Decode(&code)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	result, err := compile.ExecutePythonCode(code.Code)
	server.Ok(map[string]interface{}{"message": result}, w)
}
