package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Event struct {
	Id    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title  string `json:"title"`
	Text string `json:"text"`
	CreatedIn   time.Time `json:"createdIn"`
	CreatedBy bson.ObjectId `json:"createdBy"`
	Date  time.Time `json:"date"`
	Range int `json:range`
	Recipes []bson.ObjectId `json:"recipes"`
	User bson.ObjectId `json:"user"`
}
