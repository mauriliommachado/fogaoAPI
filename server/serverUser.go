package server

import (
	"github.com/bmizerany/pat"
	"net/http"
	"encoding/json"
	"github.com/mauriliommachado/fogaoAPI/db"
	"gopkg.in/mgo.v2/bson"
	"encoding/base64"
)

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	if !validAuthHeader(req) {
		unauthorized(w)
		return
	}
	var user db.User
	id := req.URL.Query().Get(":id")
	err := user.FindById(db.GetUsersCollection(), bson.ObjectIdHex(id))
	user.Remove(db.GetUsersCollection())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ResponseWithJSON(w, nil, http.StatusNoContent)
}

func InsertUser(w http.ResponseWriter, req *http.Request) {
	if !validAuthHeader(req) {
		unauthorized(w)
		return
	}
	var user db.User
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&user)
	if err != nil {
		badRequest(w, err)
		return
	}
	if len(user.Id.Hex()) > 0 {
		badRequest(w, nil)
		return
	}
	user.Admin = false
	user.Token = base64.StdEncoding.EncodeToString([]byte(user.Email + ":" + user.Pwd))
	err = user.Persist(db.GetUsersCollection())
	if err != nil {
		badRequest(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", req.URL.Path+"/"+user.Id.Hex())
	w.WriteHeader(http.StatusCreated)
}

func UpdateUser(w http.ResponseWriter, req *http.Request) {
	if !validAuthHeader(req) {
		unauthorized(w)
		return
	}
	var user db.User
	var userUp db.User

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&user)
	userUp.FindById(db.GetUsersCollection(), user.Id)
	if len(userUp.Id.Hex()) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		badRequest(w, err)
		return
	}
	user.Token = base64.StdEncoding.EncodeToString([]byte(user.Email + ":" + user.Pwd))
	err = user.Merge(db.GetUsersCollection())
	if err != nil {
		badRequest(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func FindAllUsers(w http.ResponseWriter, req *http.Request) {
	if !validAuthHeader(req) {
		unauthorized(w)
		return
	}
	var users db.Users
	users, err := users.FindAll(db.GetUsersCollection())
	if err != nil {
		badRequest(w, err)
		return
	}
	for i,_ := range users{
		users[i].Pwd=""
	}
	resp, _ := json.Marshal(users)
	ResponseWithJSON(w, resp, http.StatusOK)
}

func FindUserById(w http.ResponseWriter, req *http.Request) {
	var user db.User
	id := req.URL.Query().Get(":id")
	err := user.FindById(db.GetUsersCollection(), bson.ObjectIdHex(id))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp, _ := json.Marshal(user)
	ResponseWithJSON(w, resp, http.StatusOK)
}


func StartUsers(properties ServerProperties, m *pat.PatternServeMux) {
	mapEndpointsUser(*m, properties)
}
func mapEndpointsUser(m pat.PatternServeMux, properties ServerProperties) {
	m.Post(properties.Address, http.HandlerFunc(InsertUser))
	m.Put(properties.Address, http.HandlerFunc(UpdateUser))
	m.Del(properties.Address+"/:id", http.HandlerFunc(DeleteUser))
	m.Get(properties.Address, http.HandlerFunc(FindAllUsers))
	m.Get(properties.Address+"/:id", http.HandlerFunc(FindUserById))
}


