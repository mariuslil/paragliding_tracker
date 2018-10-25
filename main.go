package main

import (
	"net/http"
	"time"
)



type API struct {
	Uptime string `json:"uptime"`
	Info string `json:"info"`
	Version string `json:"version"`
}


var startTime = time.Now()
var IDs []int
var db = TrackDB{}


func main() {
	db.Init()
	http.HandleFunc("/igcinfo/api/", igcHandler)
	http.ListenAndServe(getPort(), nil)
}
