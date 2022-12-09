package main

import (
	api "maxinteg-admin-go/api/server"
	"net/http"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
