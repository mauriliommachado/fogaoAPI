package db

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

var fSession mgo.Session

func Start() {
	fmt.Println("entrei")
	session, err := mgo.Dial("mongodb://heroku_s9jcnfls:ekhui87a29vj9fdr88i0kf5t3v@ds161443.mlab.com:61443/heroku_s9jcnfls")
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
