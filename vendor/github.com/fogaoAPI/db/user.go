package db

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"
)

type User struct {
	Id    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
	Token string `json:"token"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	Admin bool
}


type Users []User

const ID_MS_URL  = "http://localhost:8080/goid"

// Creates a new user
func NewUser() (*User) {
	return &User{Name: "", Email: "", Pwd: ""}
}

func (user *User) Persist(c *mgo.Collection) error {
	var err error
	defer c.Database.Session.Close()
	user.Id = bson.NewObjectId()
	err = c.Insert(user)
	log.Println("Usuário", user.Name, "inserido")
	if err != nil {
		return err
	}
	return nil
}

func (user *User) Merge(c *mgo.Collection) error {
	var err error
	defer c.Database.Session.Close()
	err = c.Update(bson.M{"_id": user.Id}, &user)
	log.Println("Usuário", user.Name, "atualizado")
	if err != nil {
		return err
	}
	return nil
}

func (user *User) Remove(c *mgo.Collection) error {
	var err error
	defer c.Database.Session.Close()
	err = c.Remove(bson.M{"_id": user.Id})
	log.Println("Usuário", user.Name, "removido")
	if err != nil {
		return err
	}
	return nil
}

func (user *User) FindById(c *mgo.Collection, id bson.ObjectId) error {
	defer c.Database.Session.Close()
	err := c.Find(bson.M{"_id": id}).One(&user)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) FindLogin(c *mgo.Collection) bool {
	defer c.Database.Session.Close()
	err := c.Find(bson.M{"email": user.Email,"pwd":user.Pwd}).One(&user)
	if err != nil {
		return false
	}
	return true
}

func (user *User) FindHash(c *mgo.Collection) bool {
	defer c.Database.Session.Close()
	err := c.Find(bson.M{"token": user.Token}).One(&user)
	if err != nil {
		log.Println(err,user.Token)
		return false
	}
	return true
}

func (users Users) FindAll(c *mgo.Collection) (Users, error) {
	defer c.Database.Session.Close()
	err := c.Find(bson.M{}).All(&users)
	if err != nil {
		return users, err
	}
	return users, nil
}
