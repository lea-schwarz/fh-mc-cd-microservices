// main.go

package main

import (
	"os"
	"net/http"
)

func main() {
	a := App{}

	// $env:APP_DB_USERNAME="postgres"
	// $env:APP_DB_PASSWORD="geheim"
	// $env:APP_DB_NAME="postgres"
	// localhost:8010

	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	http.Handle("/", a.Router)
	http.ListenAndServe(":8010", nil)
}