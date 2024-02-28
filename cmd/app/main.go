package main

import (
	_ "github.com/lib/pq"
	"log"
	"online_store/internal/app"
)

func main() {

	cfg := app.GetConfig()

	s, err := app.NewApiServer(cfg)
	if err = s.Start(); err != nil {
		log.Fatal("failed to raise the server", err)
	}
}
