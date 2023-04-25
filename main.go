// main.go

package main

import (
	"os"
	"net/http"
)

func main() {
	a := App{}

	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	http.Handle("/", a.Router)
	http.ListenAndServe(":8010", nil)
}
