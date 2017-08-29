package db

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"os"
)

var fSession mgo.Session

func Start() {
	session, err := mgo.Dial(os.Getenv("MONGODB_URI"))
	if err != nil {
		panic(err)
	}
	fSession = *session
	fmt.Println("Sessão do banco criada")
}

func GetCollection() (c *mgo.Collection) {
	s := fSession.Copy()
	c = s.DB("heroku_s9jcnfls").C("users")
	return c
}
