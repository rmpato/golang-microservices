package main

import (
	"github.com/gorilla/mux"
	"myevents/persistence"
	"net/http"
)

func main() {
	ServeAPI()
}

//ServeAPI Initializes mux routes and serves the Events API
func ServeAPI() {
	handler := &eventServiceHandler{}
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()

	//Retrieves events by search criteria by GET
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)

	//Retrieves all events by GET
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)

	//Creates a new Event from a POST request
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	http.ListenAndServe(":8181",r)
}

//A costructor func to build the event service handler with a db handler injected
func newEventHandler(databaseHandler persistence.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler{dbhandler: databaseHandler}
}

type eventServiceHandler struct {
	dbhandler persistence.DatabaseHandler
}

func (eh *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request)  {

}

func (eh *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request)  {

}

func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request)  {

}
