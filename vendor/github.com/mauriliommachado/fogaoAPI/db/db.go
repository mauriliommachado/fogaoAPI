package db

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"os"
)

var fSession mgo.Session

func Start() {
	fmt.Printf(os.Getenv("MONGOLAB_URL"))
	session, err := mgo.Dial(os.Getenv("MONGOLAB_URL"))
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
