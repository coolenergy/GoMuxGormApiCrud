package controllers

import (
	"encoding/json"
	"net/http"
)

func sendAsJson(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			println("Error writing response:", err.Error())
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write([]byte(response))
	if err != nil {
		println("Error writing response:", err.Error())
	}
}

func sendAsJson2(w http.ResponseWriter, status int, payload interface{}) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := encoder.Encode(payload)
	if err != nil {
		println("Error writing response:", err.Error())
	}
}
