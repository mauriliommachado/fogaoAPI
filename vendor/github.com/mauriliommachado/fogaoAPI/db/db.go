package db

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"os"
)

var fSession mgo.Session

var mongo_url = "mongodb://heroku_j98w5qn4:heroku_j98w5qn4@ds161913.mlab.com:61913/heroku_j98w5qn4"

func Start() {

	uri := os.Getenv(mongo_url)
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

func GetCollection() (c *mgo.Collection) {
	s := fSession.Copy()
	c = s.DB("fogao").C("users")
	return c
}
