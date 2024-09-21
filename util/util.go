package util

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Msg string `json:"message"`
}

type ErrorMessage struct {
	Msg string `json:"error"`
}

func MsgToJson(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Message{Msg: msg})
}

func ErrorToJson(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&ErrorMessage{Msg: err.Error()})
}

func ErrorStringToJson(w http.ResponseWriter, statusCode int, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&ErrorMessage{Msg: err})
}
