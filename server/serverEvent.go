package server

import (
	"github.com/bmizerany/pat"
	"net/http"
	"encoding/json"
	"github.com/mauriliommachado/fogaoAPI/db"
	"gopkg.in/mgo.v2/bson"
)

func DeleteEvent(w http.ResponseWriter, req *http.Request) {
	if !validAuthHeader(req) {
		unauthorized(w)
		return
	}
	var event db.Event
	id := req.URL.Query().Get(":id")
	err := event.FindById(db.GetEventsCollection(), bson.ObjectIdHex(id))
	event.Remove(db.GetEventsCollection())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ResponseWithJSON(w, nil, http.StatusNoContent)
}

func InsertEvent(w http.ResponseWriter, req *http.Request) {
	if !validAuthHeader(req) {
		unauthorized(w)
		return
	}
	var event db.Event
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&event)
	if err != nil {
		badRequest(w, err)
		return
	}
	if len(event.Id.Hex()) > 0 {
		badRequest(w, nil)
		return
	}
	err = event.Persist(db.GetEventsCollection())
	if err != nil {
		badRequest(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", req.URL.Path+"/"+event.Id.Hex())
	w.WriteHeader(http.StatusCreated)
}

func UpdateEvent(w http.ResponseWriter, req *http.Request) {
	if !validAuthHeader(req) {
		unauthorized(w)
		return
	}
	var event db.Event
	var eventUp db.Event

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&event)
	eventUp.FindById(db.GetEventsCollection(), event.Id)
	if len(eventUp.Id.Hex()) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		badRequest(w, err)
		return
	}
	err = event.Merge(db.GetEventsCollection())
	if err != nil {
		badRequest(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func FindAllEvents(w http.ResponseWriter, req *http.Request) {
	if !validAuthHeader(req) {
		unauthorized(w)
		return
	}
	var events db.Events
	events, err := events.FindAll(db.GetEventsCollection())
	if err != nil {
		badRequest(w, err)
		return
	}
	resp, _ := json.Marshal(events)
	ResponseWithJSON(w, resp, http.StatusOK)
}

func FindEventById(w http.ResponseWriter, req *http.Request) {
	var event db.Event
	id := req.URL.Query().Get(":id")
	err := event.FindById(db.GetEventsCollection(), bson.ObjectIdHex(id))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp, _ := json.Marshal(event)
	ResponseWithJSON(w, resp, http.StatusOK)
}


func StartEvents(properties ServerProperties, m *pat.PatternServeMux) {
	mapEndpointsEvent(*m, properties)
}
func mapEndpointsEvent(m pat.PatternServeMux, properties ServerProperties) {
	m.Post(properties.Address, http.HandlerFunc(InsertEvent))
	m.Put(properties.Address, http.HandlerFunc(UpdateEvent))
	m.Del(properties.Address+"/:id", http.HandlerFunc(DeleteEvent))
	m.Get(properties.Address, http.HandlerFunc(FindAllEvents))
	m.Get(properties.Address+"/:id", http.HandlerFunc(FindEventById))
}


