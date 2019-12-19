package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)


const (
host   = "localhost"
port   = 5432
user   = "naz"
password = "cabbages"
dbname = "london-restaurants"
connectionStr =	"postgres://naz:cabbages@localhost:5432/london-restaurants?sslmode=disable"




	//host     = "ec2-54-217-225-16.eu-west-1.compute.amazonaws.com"
	//dbName   = "d19oo3ef47p7ao"
	//user     = "jluiokridztech"
	//port     = "5432"
	//password = "5faae49acd2a0c16bf72aedc7a033df6635fe8b21f7dcc28ec5369cc1e8aa13e"
	//uri      = "postgres://jluiokridztech:5faae49acd2a0c16bf72aedc7a033df6635fe8b21f7dcc28ec5369cc1e8aa13e@ec2-54-217-225-16.eu-west-1.compute.amazonaws.com:5432/d19oo3ef47p7ao"
)

func main() {
	http.HandleFunc("/", handler)

	//fmt.Println("listening...")
	//err := http.ListenAndServe(GetPort(), nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	dbInit()
}


func dbInit()  {
	connectionURL := os.Getenv("DATABASE_URL")
	parsed, err := pq.ParseURL(connectionURL)
	fmt.Println(connectionURL)
	fmt.Println(err)
	fmt.Println(parsed)

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
	return
}


// ------------------------

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello. This is our first Go web app on Heroku!")
}

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
