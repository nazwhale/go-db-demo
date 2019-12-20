package dao

import (
	"database/sql"
	"os"
)

type Restaurant struct {
	ID       int
	Name     string
}

func ListRestaurants(limit int) ([]Restaurant, error) {
	connectionURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return listRestaurants(db, 10), nil
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