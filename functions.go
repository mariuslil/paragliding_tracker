package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)


func replyApi (w *http.ResponseWriter){
	var endTime = time.Now()
	var api = API{
		Uptime: fmt.Sprintf("P%dY%dM%dDT%dH%dM%dS",
			endTime.Year() - startTime.Year(), endTime.Month() - startTime.Month(), endTime.Day() - startTime.Day(),
			endTime.Hour() - startTime.Hour(), endTime.Minute() - startTime.Minute(), endTime.Second() - startTime.Second()),
		Info: "Service for IGC tracks.",
		Version: "V1",
	}
	json.NewEncoder(*w).Encode(api)

}


func replyAll (w *http.ResponseWriter) {
	if db.tracks == nil {
		http.Error(*w, "No tracks available", http.StatusBadRequest)
		return
	}

	json.NewEncoder(*w).Encode(IDs)
}


func replyInfo (w *http.ResponseWriter, db *TrackDB, id string){
	intID, err := strconv.Atoi(id)
	if err != nil{
		http.Error(*w, err.Error(), http.StatusBadRequest)
		return
	}

	t, ok := db.Get(intID)
	if !ok {
		http.Error(*w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	json.NewEncoder(*w).Encode(t)
}


func replyField (w *http.ResponseWriter, db *TrackDB, id string, field string){
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusBadRequest)
		return
	}

	t, ok := db.Get(intID)
	if !ok {
		http.Error(*w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if field == "Track_length" {
		fmt.Fprint(*w, t.Track_length)
		fmt.Fprintln(*w,)
	} else if field == "H_date" {
		fieldValue := db.fieldExist(intID, field)
		fmt.Fprint(*w, fieldValue)
		fmt.Println(*w,)
	}
}

func getPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	return ":" + port
}







