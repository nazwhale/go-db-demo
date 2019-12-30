package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
	router.HandleFunc("/restaurant/create", handler.HandleCreateRestaurant).Methods("POST")
	router.HandleFunc("/restaurant/{id}", handler.HandleGetRestaurant)

	err := http.ListenAndServe(getPort(), router)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
