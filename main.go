package main

import (
	"os"
	"github.com/mauriliommachado/fogaoAPI/server"
	"github.com/mauriliommachado/fogaoAPI/db"
)

var mongo_url = "mongodb://heroku_j98w5qn4:heroku_j98w5qn4@ds161913.mlab.com:61913/heroku_j98w5qn4"

func main() {
	startDb()
	server.StartUsers(server.ServerProperties{Address: "/api/users", Port: determineListenAddress()})
}

func startDb() {
	db.Start()
}

func determineListenAddress() (string) {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}