package handler

import (
	"encoding/json"
	"net/http"
)

func Executor(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
		defer func() {
			if r := recover(); r != nil {
				RespondWithError(w, r.(error))
			}
		}()
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
		byteArr, _ := json.Marshal(map[string]string{"message": "Internal Server Error"})
		RespondwithJSON(w, http.StatusBadRequest, byteArr)
	}
}
