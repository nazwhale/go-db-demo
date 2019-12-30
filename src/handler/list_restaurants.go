package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/personal-projects/postgres-play/src/dao"
)


func HandleListRestaurants(writer http.ResponseWriter, request *http.Request) {
	restaurants, err := dao.ListRestaurants(10)

	if err != nil {
		log.Fatal("Error listing restaurants", err)
	}

	json.NewEncoder(writer).Encode(restaurants)
}

