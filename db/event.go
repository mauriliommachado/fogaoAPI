package db

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2"
	"github.com/mauriliommachado/gomodels/dbutil"
	"log"
)

type Event struct {
	Id    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title  string `json:"title"`
	Text string `json:"text"`
	CreatedIn   time.Time `json:"createdIn"`
	CreatedBy bson.ObjectId `json:"createdBy"`
	Date  time.Time `json:"date"`
	Range int `json:"range"`
	Recipes []bson.ObjectId `json:"recipes"`
}


type Events []Event

func (event *Event) Persist(c *mgo.Collection) error {
	var err error
	defer dbutil.CloseSession(c)
	event.Id = bson.NewObjectId()
	event.CreatedIn = time.Now()
	err = c.Insert(event)
	log.Println("Evento ", event.Title, "inserido")
	if err != nil {
		return err
	}
	return nil
}

func (event *Event) Merge(c *mgo.Collection) error {
	var err error
	defer dbutil.CloseSession(c)
	err = c.Update(bson.M{"_id": event.Id}, &event)
	log.Println("Evento ", event.Title, "atualizado")
	if err != nil {
		return err
	}
	return nil
}

func (event *Event) Remove(c *mgo.Collection) error {
	var err error
	defer dbutil.CloseSession(c)
	err = c.Remove(bson.M{"_id": event.Id})
	log.Println("Evento ", event.Title, "removido")
	if err != nil {
		return err
	}
	return nil
}

func (event *Event) FindById(c *mgo.Collection, id bson.ObjectId) error {
	defer dbutil.CloseSession(c)
	err := c.Find(bson.M{"_id": id}).One(&event)
	if err != nil {
		return err
	}
	return nil
}

func (events Events) FindAll(c *mgo.Collection) (Events, error) {
	defer dbutil.CloseSession(c)
	err := c.Find(bson.M{}).All(&events)
	if err != nil {
		return events, err
	}
	return events, nil
}