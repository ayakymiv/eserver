package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	//log.Fatal(http.ListenAndServe(":8000", a.Router))
	log.Fatal(http.ListenAndServeTLS(":8443", "server.crt", "server.key", a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/listener", a.addEvent).Methods("POST")
	a.Router.HandleFunc("/listener/{event}", a.removeEvent).Methods("DELETE")
	a.Router.HandleFunc("/publish/{event}", a.publishEvent).Methods("POST")
}

func (a *App) addEvent(w http.ResponseWriter, r *http.Request) {
	var e Event
	var h Handler
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		log.Fatal("Error:", err)
	}
	err = json.NewDecoder(r.Body).Decode(&h)
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println(w)
}

func (a *App) removeEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w)
}

func (a *App) publishEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w)
}