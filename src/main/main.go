package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/personal-projects/postgres-play/src/handler"
)

func main() {

	http.HandleFunc("/", handler.HandleRoot)
	http.HandleFunc("/restaurants/list", handler.HandleListRestaurants)

	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}


func DbInit() *sql.DB {
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

	listRestaurants(db,10)

	fmt.Println("Successfully connected!")
	return db
}


// ------------------------


// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}


type Restaurant struct {
	ID       int
	Name     string
}

func createRestaurant(db *sql.DB, name string) {
	sqlStatement := `
INSERT INTO restaurants (name)
VALUES ($1)
RETURNING id`

	var restaurant Restaurant
	err := db.QueryRow(sqlStatement, name).Scan(&restaurant.ID)
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
	switch err := row.Scan(&restaurant.ID, &restaurant.Name); err {
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

	fmt.Println("rows", rows)

	for rows.Next() {
		var restaurant Restaurant
		err = rows.Scan(&restaurant.ID, &restaurant.Name)
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
