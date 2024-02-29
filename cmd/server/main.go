package main

import (
	_ "github.com/lib/pq"
	"log"
	"online_store/internal/app"
	"online_store/internal/config"
)

func main() {

	cfg := config.GetConfig()

	s, err := app.NewApiServer(cfg)
	if err = s.Start(); err != nil {
		log.Fatal("не удалось подключиться к серверу", err)
	}
}
