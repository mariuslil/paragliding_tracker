package main

import (
	"encoding/json"
	"github.com/marni/goigc"
	"net/http"
	"strings"
)



func igcHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if r.Body == nil {
			http.Error(w, "POST request must have a JSON body", http.StatusBadRequest)
			return
		}

		url := strings.Split(r.URL.Path, "/")
		if url[3] == "igc" {
			var t Track
			url := make(map[string]string)
			err := json.NewDecoder(r.Body).Decode(&url)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			track, err := igc.ParseLocation(url["url"])
			if err != nil {
				http.Error(w, "Problem reading the track info", http.StatusBadRequest)
				return
			}

			t.H_date = track.Header.Date.String()
			t.Pilot = track.Pilot
			t.Glider = track.GliderType
			t.Glider_id = track.GliderID
			t.Track_length = track.Points[0].Distance(track.Points[len(track.Points)-1])

			id := db.Add(t)
			IDs = append(IDs, id)
			json.NewEncoder(w).Encode(id)
			http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
		}

	case "GET":
		w.Header().Set("Content-Type", "application/json")
		url := strings.Split(r.URL.Path, "/")

		if url[2] == "api" && len(url) == 3 {
			replyApi(&w)
			http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
		} else if url[2] == "api" && url[3] == "igc" && len(url) == 4 {
			replyAll(&w)
			http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
		} else if url[2] == "api" && url[3] == "igc" && len(url) == 5 {
			id := url[4]
			replyInfo(&w, &db, id)
			http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
		} else if url[2] == "api" && url[3] == "igc" && len(url) == 6 {
			id := url[4]
			field := url[5]
			replyField(&w, &db, id, field)
			http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
		} else {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}

	default:
		http.Error(w, "Not yet implemented", http.StatusNotImplemented)
		return

	}
}