package db

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/mauriliommachado/gomodels/dbutil"
	"log"
)

type Recipe struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	Ingredient []bson.ObjectId `json:"ingredient"`
	Quantity    int `json:"quantity"`
	Observation string `json:obs`
}

type Recipes []Recipe

func (recipe *Recipe) Persist(c *mgo.Collection) error {
	var err error
	defer dbutil.CloseSession(c)
	recipe.Id = bson.NewObjectId()
	err = c.Insert(recipe)
	log.Println("Receita ", recipe.Title, "inserido")
	if err != nil {
		return err
	}
	return nil
}

func (recipe *Recipe) Merge(c *mgo.Collection) error {
	var err error
	defer dbutil.CloseSession(c)
	err = c.Update(bson.M{"_id": recipe.Id}, &recipe)
	log.Println("Receita ", recipe.Title, "atualizado")
	if err != nil {
		return err
	}
	return nil
}

func (recipe *Recipe) Remove(c *mgo.Collection) error {
	var err error
	defer dbutil.CloseSession(c)
	err = c.Remove(bson.M{"_id": recipe.Id})
	log.Println("Receita ", recipe.Title, "removido")
	if err != nil {
		return err
	}
	return nil
}

func (recipe *Recipe) FindById(c *mgo.Collection, id bson.ObjectId) error {
	defer dbutil.CloseSession(c)
	err := c.Find(bson.M{"_id": id}).One(&recipe)
	if err != nil {
		return err
	}
	return nil
}

func (recipes Recipes) FindAll(c *mgo.Collection) (Recipes, error) {
	defer dbutil.CloseSession(c)
	err := c.Find(bson.M{}).All(&recipes)
	if err != nil {
		return recipes, err
	}
	return recipes, nil
}