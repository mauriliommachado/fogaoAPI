package server

import (
	"github.com/bmizerany/pat"
	"net/http"
	"encoding/json"
	"github.com/mauriliommachado/fogaoAPI/db"
	"gopkg.in/mgo.v2/bson"
)

func DeleteRecipe(w http.ResponseWriter, req *http.Request) {
	if !validAuthHeader(req) {
		unauthorized(w)
		return
	}
	var recipe db.Recipe
	id := req.URL.Query().Get(":id")
	err := recipe.FindById(db.GetRecipesCollection(), bson.ObjectIdHex(id))
	recipe.Remove(db.GetRecipesCollection())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ResponseWithJSON(w, nil, http.StatusNoContent)
}

func InsertRecipe(w http.ResponseWriter, req *http.Request) {
	if !validAuthHeader(req) {
		unauthorized(w)
		return
	}
	var recipe db.Recipe
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&recipe)
	if err != nil {
		badRequest(w, err)
		return
	}
	if len(recipe.Id.Hex()) > 0 {
		badRequest(w, nil)
		return
	}
	err = recipe.Persist(db.GetRecipesCollection())
	if err != nil {
		badRequest(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", req.URL.Path+"/"+recipe.Id.Hex())
	w.WriteHeader(http.StatusCreated)
}

func UpdateRecipe(w http.ResponseWriter, req *http.Request) {
	if !validAuthHeader(req) {
		unauthorized(w)
		return
	}
	var recipe db.Recipe
	var recipeUp db.Recipe

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&recipe)
	recipeUp.FindById(db.GetRecipesCollection(), recipe.Id)
	if len(recipeUp.Id.Hex()) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		badRequest(w, err)
		return
	}
	err = recipe.Merge(db.GetRecipesCollection())
	if err != nil {
		badRequest(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func FindAllRecipes(w http.ResponseWriter, req *http.Request) {
	if !validAuthHeader(req) {
		unauthorized(w)
		return
	}
	var recipes db.Recipes
	recipes, err := recipes.FindAll(db.GetRecipesCollection())
	if err != nil {
		badRequest(w, err)
		return
	}
	resp, _ := json.Marshal(recipes)
	ResponseWithJSON(w, resp, http.StatusOK)
}

func FindRecipeById(w http.ResponseWriter, req *http.Request) {
	var recipe db.Recipe
	id := req.URL.Query().Get(":id")
	err := recipe.FindById(db.GetRecipesCollection(), bson.ObjectIdHex(id))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp, _ := json.Marshal(recipe)
	ResponseWithJSON(w, resp, http.StatusOK)
}


func StartRecipes(properties ServerProperties, m *pat.PatternServeMux) {
	mapEndpointsRecipe(*m, properties)
}
func mapEndpointsRecipe(m pat.PatternServeMux, properties ServerProperties) {
	m.Post(properties.Address, http.HandlerFunc(InsertRecipe))
	m.Put(properties.Address, http.HandlerFunc(UpdateRecipe))
	m.Del(properties.Address+"/:id", http.HandlerFunc(DeleteRecipe))
	m.Get(properties.Address, http.HandlerFunc(FindAllRecipes))
	m.Get(properties.Address+"/:id", http.HandlerFunc(FindRecipeById))
}


