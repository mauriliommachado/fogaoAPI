package server

import (
	"github.com/bmizerany/pat"
	"net/http"
	"encoding/json"
	"github.com/mauriliommachado/fogaoAPI/db"
	"gopkg.in/mgo.v2/bson"
)

func DeleteIngredient(w http.ResponseWriter, req *http.Request) {
	if !validAuthHeader(req) {
		unauthorized(w)
		return
	}
	var ingredient db.Ingredient
	id := req.URL.Query().Get(":id")
	err := ingredient.FindById(db.GetIngrediantsCollection(), bson.ObjectIdHex(id))
	ingredient.Remove(db.GetIngrediantsCollection())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ResponseWithJSON(w, nil, http.StatusNoContent)
}

func InsertIngredient(w http.ResponseWriter, req *http.Request) {
	if !validAuthHeader(req) {
		unauthorized(w)
		return
	}
	var ingredient db.Ingredient
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&ingredient)
	if err != nil {
		badRequest(w, err)
		return
	}
	if len(ingredient.Id.Hex()) > 0 {
		badRequest(w, nil)
		return
	}
	err = ingredient.Persist(db.GetIngrediantsCollection())
	if err != nil {
		badRequest(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", req.URL.Path+"/"+ingredient.Id.Hex())
	w.WriteHeader(http.StatusCreated)
}

func UpdateIngredient(w http.ResponseWriter, req *http.Request) {
	if !validAuthHeader(req) {
		unauthorized(w)
		return
	}
	var ingredient db.Ingredient
	var ingredientUp db.Ingredient

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&ingredient)
	ingredientUp.FindById(db.GetIngrediantsCollection(), ingredient.Id)
	if len(ingredientUp.Id.Hex()) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		badRequest(w, err)
		return
	}
	err = ingredient.Merge(db.GetIngrediantsCollection())
	if err != nil {
		badRequest(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func FindAllIngredients(w http.ResponseWriter, req *http.Request) {
	if !validAuthHeader(req) {
		unauthorized(w)
		return
	}
	var ingredients db.Ingredients
	ingredients, err := ingredients.FindAll(db.GetIngrediantsCollection())
	if err != nil {
		badRequest(w, err)
		return
	}
	resp, _ := json.Marshal(ingredients)
	ResponseWithJSON(w, resp, http.StatusOK)
}

func FindIngredientById(w http.ResponseWriter, req *http.Request) {
	var ingredient db.Ingredient
	id := req.URL.Query().Get(":id")
	err := ingredient.FindById(db.GetIngrediantsCollection(), bson.ObjectIdHex(id))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp, _ := json.Marshal(ingredient)
	ResponseWithJSON(w, resp, http.StatusOK)
}


func StartIngredients(properties ServerProperties, m *pat.PatternServeMux) {
	mapEndpointsIngredient(*m, properties)
}
func mapEndpointsIngredient(m pat.PatternServeMux, properties ServerProperties) {
	m.Post(properties.Address, http.HandlerFunc(InsertIngredient))
	m.Put(properties.Address, http.HandlerFunc(UpdateIngredient))
	m.Del(properties.Address+"/:id", http.HandlerFunc(DeleteIngredient))
	m.Get(properties.Address, http.HandlerFunc(FindAllIngredients))
	m.Get(properties.Address+"/:id", http.HandlerFunc(FindIngredientById))
}


