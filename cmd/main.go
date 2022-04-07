package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelwintan/yoms-api-go/domain"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent domain.Event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "invalid request body")
	}

	json.Unmarshal(reqBody, &newEvent)
	domain.Events = append(domain.Events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")

	addr := ":8080"
	fmt.Printf("%+v", domain.Events)
	log.Fatal(http.ListenAndServe(addr, router))
}
