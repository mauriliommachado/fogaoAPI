package db

import (
	"gopkg.in/mgo.v2/bson"
)

type Recipe struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	Ingredients []bson.ObjectId `json:"ingredients"`
	Quantity    int `json: "quantity"`
	Observation string `json:obs`
}