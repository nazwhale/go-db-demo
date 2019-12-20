package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

type Restaurant struct {
	ID       int
	Name     string
}

func HandleListRestaurants(writer http.ResponseWriter, request *http.Request) {
	connectionURL := os.Getenv("DATABASE_URL")

	// see if we can pass through one big database connection string
	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	restaurants := listRestaurants(db, 10)
	for _, restaurant := range restaurants {
		fmt.Sprintf( "%v", restaurant.Name)
	}

	_, _ = fmt.Fprintf(writer, "Restaurants listed!")
}

func listRestaurants(db *sql.DB, limit int) []Restaurant {
	sqlStatement := `
SELECT *
FROM restaurants
LIMIT $1;`

	rows, err := db.Query(sqlStatement, limit)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	restaurants := []Restaurant{}
	for rows.Next() {
		var restaurant Restaurant
		err = rows.Scan(&restaurant.ID, &restaurant.Name)
		if err != nil {
			panic(err)
		}
		restaurants = append(restaurants, restaurant)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return restaurants
}