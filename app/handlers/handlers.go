package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/c0z0/go-shrt/app/data"
	"github.com/c0z0/go-shrt/app/model"
	"github.com/gorilla/mux"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func Index(db *data.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if url := db.Find(vars["id"]); url != nil {
			http.Redirect(w, r, url.Url, http.StatusMovedPermanently)
		} else {
			w.WriteHeader(404)
			w.Write([]byte("404 page not found"))
		}
	}
}

func Shorten(db *data.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		url := model.Url{}

		if err := decoder.Decode(&url); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad request"))
			return
		}

		url.Url = fixUrl(url.Url)

		url = db.Create(url.Url)

		sendJSON(w, 200, url)
	}
}

func fixUrl(url string) string {
	if matched, _ := regexp.Match("(^http://|https://).+", []byte(url)); matched {
		return url
	}
	return "https://" + url
}

func sendJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
