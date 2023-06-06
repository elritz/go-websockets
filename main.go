package main

import (
	"api/api"
	"net/http"
)

func main() {
	srv := api.NewServer()

	http.ListenAndServe(":8000", srv)

}
