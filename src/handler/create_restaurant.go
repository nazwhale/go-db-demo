package handler

import (
	"encoding/json"
	"github.com/personal-projects/postgres-play/src/dao"
	"io/ioutil"
	"net/http"
)

func HandleCreateRestaurant(writer http.ResponseWriter, request *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(request.Body)

	var restaurant dao.Restaurant
	json.Unmarshal(reqBody, &restaurant)

	newRestaurant, err := dao.CreateRestaurant(restaurant.Name)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(writer).Encode(newRestaurant)
}