package main

import (
	"net/http"
)

func main() {
	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUserHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
