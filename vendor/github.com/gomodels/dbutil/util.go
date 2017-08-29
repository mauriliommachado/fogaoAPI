package dbutil

import "gopkg.in/mgo.v2"

func CloseSession(c *mgo.Collection) {
	c.Database.Session.Close()
}
