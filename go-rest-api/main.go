package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Welcome Home!")
}

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvent []event

var events = allEvent{
	{
		ID:          "1",
		Title:       "go ",
		Description: "Test",
	},
	{
		ID:          "12",
		Title:       "second ",
		Description: "#test",
	},
}

func creatEvent(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}
func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", creatEvent)
	router.HandleFunc("/event/{id}", getOneEvent)
	router.HandleFunc("/", getAll)
	http.ListenAndServe(":8080", router)
}
