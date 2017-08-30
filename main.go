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
	properties := server.ServerProperties{Address: "/api/users", Port: determineListenAddress()}
	initServers(handler, properties, m)
	err := http.ListenAndServe(":"+properties.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func initServers(handler http.Handler,properties server.ServerProperties, m *pat.PatternServeMux ) {
	server.StartUsers(properties, m)
	fmt.Println("servidor iniciado no endereço localhost:" + properties.Port + properties.Address)
	properties.Address = "/api/ingredients"
	server.StartIngredients(properties, m)
	fmt.Println("servidor iniciado no endereço localhost:" + properties.Port + properties.Address)
	properties.Address = "/api/recipes"
	server.StartRecipes(properties, m)
	fmt.Println("servidor iniciado no endereço localhost:" + properties.Port + properties.Address)
	http.Handle("/", handler)
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