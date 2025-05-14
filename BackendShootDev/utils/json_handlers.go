package utils

import (
    "net/http"
    "log"
    "encoding/json"
)

func RespondWithError(w http.ResponseWriter, code int, msg string, err error) {
    if err != nil {
        log.Println(err)
    }
    if code > 499{
        log.Printf("Responding with %d error %s", code, msg)
    }
    type errorResponse struct {
        Error string `json:"error"`
    }
    RespondWithJson(w, code, errorResponse{
        Error: msg, 
    })

}

func RespondWithJson(w http.ResponseWriter, code int, payload any) {
    w.Header().Set("Content-Type", "application/json")
    data, err := json.Marshal(payload)
    if err != nil {
	log.Printf("Error marshalling JSON: %s", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.WriteHeader(code)
    w.Write(data)
}

