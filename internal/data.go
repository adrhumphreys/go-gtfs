package internal

import (
	"gtfs-cli/internal/dbx"
	"log"
)

func GetRoutes() []Route {
	db := dbx.Connect()

	var routes []Route

	err := db.Select(&routes, "SELECT * FROM route ORDER BY long_name")

	if err != nil {
		log.Fatal(err)
	}

	return routes
}