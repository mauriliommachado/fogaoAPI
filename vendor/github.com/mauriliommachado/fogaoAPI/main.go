package main

import (
	"os"
	"github.com/mauriliommachado/fogaoAPI/server"
	"github.com/mauriliommachado/fogaoAPI/db"
)

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