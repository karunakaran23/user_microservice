package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
}
type data map[string]interface{}

func JSONWriter(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/home" {
		fmt.Fprintf(w, "Welcome!")
		return
	}
}

func (h *Handler) MicroserviceName(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		JSONWriter(w, data{
			"Error": "Method Not Allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	JSONWriter(w, data{"Response": "User Microservice"}, 200)
}
