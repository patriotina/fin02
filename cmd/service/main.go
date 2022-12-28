package main

import (
	service "fin02/internal/service/http"
	"net/http"
)

func main() {
	server := service.New()
	err := http.ListenAndServe("localhost:3000", server.BuildRoutes())
	if err != nil {
		panic(err)
	}
}
