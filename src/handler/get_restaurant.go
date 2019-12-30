package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/personal-projects/postgres-play/src/dao"
	"log"
	"net/http"
	"strconv"
)

func HandleGetRestaurant(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("Error parsing id key as int", err)
	}

	restaurants, err := dao.ListRestaurants(10)
	if err != nil {
		log.Fatal("Error listing restaurants", err)
	}

	for _, r := range restaurants {
		if r.ID == key {
			json.NewEncoder(writer).Encode(r)
		}
	}
}
