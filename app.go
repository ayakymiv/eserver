package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
	ModelsInit()
}

func (a *App) Run() {
	//log.Fatal(http.ListenAndServe(":8000", a.Router))
	var port = os.Getenv("PORT_NUMBER")
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%d", port), "server.crt", "server.key", a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/listener", a.addEvent).Methods("POST")
	a.Router.HandleFunc("/listener/{event}", a.removeEvent).Methods("DELETE")
	a.Router.HandleFunc("/publish/{event}", a.publishEvent).Methods("POST")
	//a.Router.PathPrefix("/").Handler(a.HandleAll)
}

func (a *App) addEvent(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer r.Body.Close()

	var e = Event{}
	err = json.NewDecoder(bytes.NewReader(b)).Decode(&e)
	if err != nil {
		log.Fatal("Error:", err)
	}
	var h Handler
	err = json.NewDecoder(bytes.NewReader(b)).Decode(&h)
	e.Handlers = []Handler{h}
	if err != nil {
		log.Fatal("Error:", err)
	}
	Store(e)
}

func (a *App) removeEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ModelRemove(vars["event"])
	fmt.Println(w)
}

func (a *App) publishEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var body interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Fatal("Error:", err)
	}
	ModelPublish(vars["event"],body)
	fmt.Println(w)
}

func (a *App) HandleAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w)
}