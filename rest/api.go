package rest

import (
	"github.com/gorilla/mux"
	"myevents/persistence"
	"net/http"
)

//ServeAPI Initializes mux routes and serves the Events API
func ServeAPI(endpoint string, dbHandler persistence.DatabaseHandler) error {
	handler := newEventHandler(dbHandler)
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()

	//Retrieves events by search criteria by GET
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)

	//Retrieves all events by GET
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)

	//Creates a new Event from a POST request
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	return http.ListenAndServe(endpoint, r)
}

//A constructor func to build the event service handler with a db handler injected
func newEventHandler(databaseHandler persistence.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler{dbhandler: databaseHandler}
}
