package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "calhounio_demo"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"dbname=%s sslmode=disable", host, port, user, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	listRestaurants(db, 5)
}

type Restaurant struct {
	ID int
	Name string
	Area string
	ImageURL string
	Cuisine string
}

func createRestaurant(db *sql.DB) {
	sqlStatement := `
INSERT INTO restaurants (name, area, image_url, cuisine)
VALUES ($1, $2, $3, $4)
RETURNING id`

	var restaurant Restaurant
	err := db.QueryRow(sqlStatement, "Kebab Man", "Holloway", "https://www.toconoco.com/wp-content/uploads/Toconoco-inside.jpg", "Korean").Scan(&restaurant.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println("New record:", restaurant)
}

func readRestaurant(db *sql.DB, restaurantID int) {
	sqlStatement := `
SELECT *
FROM restaurants
WHERE id = $1;`

	var restaurant Restaurant
	row := db.QueryRow(sqlStatement, restaurantID)
	switch err := row.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Area, &restaurant.ImageURL, &restaurant.Cuisine); err {
	case sql.ErrNoRows:
		fmt.Println("No rows returned!")
	case nil:
		fmt.Println(restaurant)
	default:
		panic(err)
	}
}

func listRestaurants(db *sql.DB, limit int) {
	sqlStatement := `
SELECT *
FROM restaurants
LIMIT $1;`

	rows, err := db.Query(sqlStatement, limit)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var restaurant Restaurant
		err = rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Area, &restaurant.ImageURL, &restaurant.Cuisine)
		if err != nil {
			panic(err)
		}
		fmt.Println(restaurant)
	}

	err = rows.Err()
	if err != nil {
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
