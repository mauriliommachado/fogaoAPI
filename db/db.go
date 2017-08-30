package db

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"os"
)

var fSession mgo.Session

func Start() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri= "localhost"
	}
	session, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}
	fSession = *session
	fmt.Println("Sessão do banco criada")
}

func GetUsersCollection() (c *mgo.Collection) {
	s := fSession.Copy()
	c = s.DB("heroku_s9jcnfls").C("users")
	return c
}

func GetIngrediantsCollection() (c *mgo.Collection) {
	s := fSession.Copy()
	c = s.DB("heroku_s9jcnfls").C("ingrediants")
	return c
}

func GetRecipesCollection() (c *mgo.Collection) {
	s := fSession.Copy()
	c = s.DB("heroku_s9jcnfls").C("recipes")
	return c
}

func GetEventsCollection() (c *mgo.Collection) {
	s := fSession.Copy()
	c = s.DB("heroku_s9jcnfls").C("events")
	return c
}

