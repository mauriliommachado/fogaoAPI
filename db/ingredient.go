package db

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/mauriliommachado/gomodels/dbutil"
	"log"
)

type Ingredient struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name       string `json:"name"`
	Text        string `json:"text"`
	Quantity    int `json:"quantity"`
	Unit 	string `json:"unit"`
	Observation string `json:"obs"`
}

type Ingredients []Ingredient


func (ingredient *Ingredient) Persist(c *mgo.Collection) error {
	var err error
	defer dbutil.CloseSession(c)
	ingredient.Id = bson.NewObjectId()
	err = c.Insert(ingredient)
	log.Println("Ingrediente", ingredient.Name, "inserido")
	if err != nil {
		return err
	}
	return nil
}

func (ingredient *Ingredient) Merge(c *mgo.Collection) error {
	var err error
	defer dbutil.CloseSession(c)
	err = c.Update(bson.M{"_id": ingredient.Id}, &ingredient)
	log.Println("Ingrediente", ingredient.Name, "atualizado")
	if err != nil {
		return err
	}
	return nil
}

func (ingredient *Ingredient) Remove(c *mgo.Collection) error {
	var err error
	defer dbutil.CloseSession(c)
	err = c.Remove(bson.M{"_id": ingredient.Id})
	log.Println("Ingrediente", ingredient.Name, "removido")
	if err != nil {
		return err
	}
	return nil
}

func (ingredient *Ingredient) FindById(c *mgo.Collection, id bson.ObjectId) error {
	defer dbutil.CloseSession(c)
	err := c.Find(bson.M{"_id": id}).One(&ingredient)
	if err != nil {
		return err
	}
	return nil
}

func (ingredients Ingredients) FindAll(c *mgo.Collection) (Ingredients, error) {
	defer dbutil.CloseSession(c)
	err := c.Find(bson.M{}).All(&ingredients)
	if err != nil {
		return ingredients, err
	}
	return ingredients, nil
}