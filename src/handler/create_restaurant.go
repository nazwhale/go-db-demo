package handler

import (
	"encoding/json"
	"github.com/personal-projects/postgres-play/src/dao"
	"io/ioutil"
	"log"
	"net/http"
)

func HandleCreateRestaurant(writer http.ResponseWriter, request *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(request.Body)

	log.Printf("reqbody: %v", reqBody)

	var restaurant dao.Restaurant
	err := json.Unmarshal(reqBody, &restaurant)
	if err != nil {
		panic(err)
	}

	log.Printf("restaurant: %v", restaurant)
	log.Printf("restaurant name: %v", restaurant.Name)

	newRestaurantID, err := dao.CreateRestaurant(restaurant.Name)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(writer).Encode(newRestaurantID)
}