package db

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

var fSession mgo.Session

var mongo_url = "mongodb://fogaoAdmin:fogaoAdmin@ds161913.mlab.com:61913/heroku_j98w5qn4"

func Start() {
	session, err := mgo.Dial(mongo_url)
	if err != nil {
		panic(err)
	}
	fSession = *session
	fmt.Println("Sessão do banco criada")
}

func GetCollection() (c *mgo.Collection) {
	s := fSession.Copy()
	c = s.DB("fogao").C("users")
	return c
}
