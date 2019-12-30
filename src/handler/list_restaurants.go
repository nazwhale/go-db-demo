package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/personal-projects/postgres-play/src/dao"
)


func HandleListRestaurants(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "about to list...")

	restaurants, err := dao.ListRestaurants(10)

	_, _ = fmt.Fprintf(writer, "listed!")

	if err != nil {
		log.Fatal("Error listing restaurants", err)
	}

	json.NewEncoder(writer).Encode(restaurants)
}

