package app

import (
	"log"
	"net/http"
	"os"

	"github.com/c0z0/go-shrt/app/data"
	"github.com/c0z0/go-shrt/app/handlers"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Init() {
	db := data.NewDB(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DB"))
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/{id}", handlers.Index(db)).Methods("GET")
	a.Router.HandleFunc("/s", handlers.Shorten(db)).Methods("POST")
}

func (a *App) Run(host string) {
	log.Printf("App listening on %s", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}
