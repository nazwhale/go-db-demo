package handler

import (
	"fmt"
	"github.com/personal-projects/postgres-play/src/dao"
	"log"
	"net/http"
)


func HandleListRestaurants(writer http.ResponseWriter, request *http.Request) {
	restaurants, err := dao.ListRestaurants(10)
	if err != nil {
		log.Fatal("Error listing restaurants", err)
		panic(err)
	}

	for _, restaurant := range restaurants {
		_, _ = fmt.Fprintf(writer, "%v\n", restaurant.Name)
	}

	_, _ = fmt.Fprintf(writer, "Restaurants listed!")
}

