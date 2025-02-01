package main

import (
	"log"

	"github.com/odbalapure/social/cmd/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("PORT", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
