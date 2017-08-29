package main

import "gopkg.in/mgo.v2/bson"

type Ingredients struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	Quantity    int `json: "quantity"`
	Observation string `json:obs`
}