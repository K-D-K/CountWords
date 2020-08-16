package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Executor(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r.(error))
				RespondWithError(w, r.(error))
			}
		}()
		handler(w, r)
	}
}

// RespondwithJSON : generic handling to send response.
func RespondwithJSON(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithError : handle errors in project
func RespondWithError(w http.ResponseWriter, err error) {
	switch err.(type) {
	default:
		errResp := make(map[string]string)
		errResp["message"] = "Internal Server Error"
		errResp["error"] = err.Error()
		byteArr, _ := json.Marshal(errResp)
		RespondwithJSON(w, http.StatusBadRequest, byteArr)
	}
}
