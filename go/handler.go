package main

import (
	"encoding/json"
	"github.com/marni/goigc"
	"net/http"
	"strings"
)



func apiHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if r.Body == nil {
			http.Error(w, "POST request must have a JSON body", http.StatusBadRequest)
			return
		}

		URL := strings.Split(r.URL.Path, "/")
		if URL[3] == "track" {
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
			t.Track_src_url = url["url"]


			id := db.Add(t)
			IDs = append(IDs, id)
			json.NewEncoder(w).Encode(id)
			http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)

		} else if URL[3] == "webohook" && URL[4] == "new_track" {

		}

	case "GET":
		url := strings.Split(r.URL.Path, "/")
		if len(url) == 6 || url[4] == "latest" {
			w.Header().Set("Content-Type", "text/plain")
		} else {
			w.Header().Set("Content-Type", "application/json")
		}

		if url[2] == "api" {
			if url[3] == "track" {

				if len(url) == 5 {
					replyInfo(&w, &db, url[4])
					http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
					return

				} else if len(url) == 6 {
					replyField(&w, &db, url[4], url[5])
					http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
					return
				}

				replyAll(&w)
				http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
				return

			} else if url[3] == "ticker" {

				if url[4] == "latest" {

					replyLatest()
					http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
					return

				} else if url[4] != "" {

					replyTime()
					http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
					return
				}

				replyTicker()
				http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
				return

			} else if url[3] == "webhook" && url[4] == "new_track" {


				http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
				return
			}


			replyApi(&w)
			http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
			return

		} else {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}

	case "DELETE":
		w.Header().Set("Content-Type", "application/json")
		url := strings.Split(r.URL.Path, "/")

		if url[3] == "webhook" && url[4] == "new_track" && len(url) == 5 {



		} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}


	default:
		http.Error(w, "Not yet implemented", http.StatusNotImplemented)
		return

	}
}

func adminHandler(w http.ResponseWriter, r *http.Request) {

}