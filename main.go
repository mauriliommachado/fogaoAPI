package main

import (
	"os"
	"github.com/mauriliommachado/fogaoAPI/server"
	"github.com/mauriliommachado/fogaoAPI/db"
	"github.com/bmizerany/pat"
	"github.com/rs/cors"
	"fmt"
	"log"
	"net/http"
)


func main() {
	startDb()
	m := pat.New()
	handler := cors.AllowAll().Handler(m)
	properties := server.ServerProperties{Address: "/api", Port: determineListenAddress()}
	server.StartUsers(m)
	http.Handle(properties.Address, handler)
	fmt.Println("servidor iniciado no endereço localhost:" + properties.Port + properties.Address)
	err := http.ListenAndServe(":"+properties.Port, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
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