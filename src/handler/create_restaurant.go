package handler

import (
	"encoding/json"
	"fmt"
	"github.com/personal-projects/postgres-play/src/dao"
	"log"
	"net/http"
	"strings"
)

func HandleCreateRestaurant(w http.ResponseWriter, r *http.Request) {
	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the
	// response body. A r body larger than that will now result in
	// Decode() returning a "http: r body too large" error.
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	// Setup the decoder and call the DisallowUnknownFields() method on it.
	// This will cause Decode() to return a "json: unknown field ..." error
	// if it encounters any extra unexpected fields in the JSON. Strictly
	// speaking, it returns an error for "keys which do not match any
	// non-ignored, exported fields in the destination".
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var restaurant dao.Restaurant
	err := decoder.Decode(&restaurant)
	if err != nil {
		switch {
		// Catch the error caused by extra unexpected fields in the r
		// body. We extract the field name from the error message and
		// interpolate it in our custom error message. There is an open
		// issue at https://github.com/golang/go/issues/29035 regarding
		// turning this into a sentinel error.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			http.Error(w, msg, http.StatusBadRequest)

		// Catch the error caused by the r body being too large. Again
		// there is an open issue regarding turning this into a sentinel
		// error at https://github.com/golang/go/issues/30715.
		case err.Error() == "http: r body too large":
			msg := "Request body must not be larger than 1MB"
			http.Error(w, msg, http.StatusRequestEntityTooLarge)

		// Otherwise default to logging the error and sending a 500 Internal
		// Server Error response.
		default:
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	// Check that the r body only contained a single JSON object.
	if decoder.More() {
		msg := "Request body must only contain a single JSON object"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "restaurant: %+v", restaurant)
	fmt.Fprintf(w, "restaurant without field names: %v", restaurant)
	fmt.Fprintf(w, "restaurant name: %v", restaurant.Name)

	newRestaurantID, err := dao.CreateRestaurant(restaurant.Name)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(newRestaurantID)
}
