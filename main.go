package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"myevents/persistence"
	"net/http"
	"strings"
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

	http.ListenAndServe(":8181", r)
}

//A costructor func to build the event service handler with a db handler injected
func newEventHandler(databaseHandler persistence.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler{dbhandler: databaseHandler}
}

type eventServiceHandler struct {
	dbhandler persistence.DatabaseHandler
}

func (eh *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No search criteria found, you can either search by id via /id/4
                   to search by name via /name/coldplayconcert}`)
		return
	}

	searchkey, ok := vars["search"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No search keys found, you can either search by id via /id/4
                   to search by name via /name/coldplayconcert}`)
		return
	}

	var event persistence.Event
	var err error
	switch strings.ToLower(criteria) {
	case "name":
		event, err = eh.dbhandler.FindEventByName(searchkey)
	case "id":
		id, err := hex.DecodeString(searchkey)
		if err == nil {
			event, err = eh.dbhandler.FindEvent(id)
		}
	}
	if err != nil {
		fmt.Fprintf(w, "{error %s}", err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&event)
}

func (eh *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {

}

func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {

}
