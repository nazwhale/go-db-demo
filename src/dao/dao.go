package dao

import (
	"database/sql"
	"fmt"
	"os"
)

type Restaurant struct {
	ID       int
	Name     string
}

func ListRestaurants(limit int) ([]Restaurant, error) {
	connectionURL := os.Getenv("DATABASE_URL")

	_ = fmt.Sprintf("connection url: %v", connectionURL)

	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return listRestaurants(db, limit), nil
}

func CreateRestaurant(name string) (int, error) {
	connectionURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	return createRestaurant(db, name), nil
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


func readRestaurant(db *sql.DB, restaurantID int) {
	sqlStatement := `
SELECT *
FROM restaurants
WHERE id = $1;`

	var restaurant Restaurant
	row := db.QueryRow(sqlStatement, restaurantID)
	switch err := row.Scan(&restaurant.ID, &restaurant.Name); err {
	case sql.ErrNoRows:
		fmt.Println("No rows returned!")
	case nil:
		fmt.Println(restaurant)
	default:
		panic(err)
	}
}


func updateRestaurant(db *sql.DB) {
	sqlStatement := `
UPDATE restaurants
SET name = $2, area = $3
WHERE id = $1;`

	rsp, err := db.Exec(sqlStatement, 1, "Hoopy Town", "Madrid")
	if err != nil {
		panic(err)
	}

	count, err := rsp.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("rows affected: ", count)
}

func deleteRestaurant(db *sql.DB) {
	sqlStatement := `
DELETE from restaurants
WHERE id = $1;`

	rsp, err := db.Exec(sqlStatement, 3)
	if err != nil {
		panic(err)
	}

	count, err := rsp.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("rows affected: ", count)
}

func createRestaurant(db *sql.DB, name string) int {
	sqlStatement := `
INSERT INTO restaurants (name)
VALUES ($1)
RETURNING id`
	id := 0

	err := db.QueryRow(sqlStatement, name).Scan(&id)
	if err != nil {
		panic(err)
	}

	return id
}
