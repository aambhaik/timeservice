package main

import (
	"fmt"
	"net/http"
	"time"

	"encoding/json"
	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	router.GET("/now", GetTimeHandler)
	router.GET("/ping", PingHandler)

	fmt.Println("Starting Time service on port", 9980)
	http.ListenAndServe(":9980", router)
}

func PingHandler(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	writer.WriteHeader(200)
}

func GetTimeHandler(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	now := time.Now().Format(time.RFC3339)
	//fmt.Println("Received new request for time at", now)

	v := Time{
		Time: now,
	}
	writeJSON(writer, http.StatusOK, v)
}

func writeJSON(w http.ResponseWriter, code int, value interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(value)
}

type Time struct {
	Time string `json:"time"`
}
