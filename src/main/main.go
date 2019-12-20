package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/personal-projects/postgres-play/src/handler"
)

func main() {
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler.HandleRoot)
	router.HandleFunc("/restaurants/list", handler.HandleListRestaurants)
	router.HandleFunc("/restaurant/{id}", handler.HandleGetRestaurant)

	err := http.ListenAndServe(GetPort(), router)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
