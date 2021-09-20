package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	handler := &eventServiceHandler{}
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()

	//Retrieves events by search criteria by GET
	eventsrouter.Methods("GET").
		Path("/{SearchCriteria}/{search}").
		HandlerFunc(handler.findEventHandler)

	//Retrieves all events by GET
	eventsrouter.Methods("GET").
		Path("").
		HandlerFunc(handler.allEventHandler)

	//Creates a new Event from a POST request
	eventsrouter.Methods("POST").
		Path("").
		HandlerFunc(handler.newEventHandler)

	http.ListenAndServe(":8181",r)
}

type eventServiceHandler struct {

}

func (eh *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request)  {

}

func (eh *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request)  {

}

func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request)  {

}
