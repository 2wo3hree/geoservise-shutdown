package main

import (
	"geoservise/internal/config"
	"geoservise/internal/router"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	r := router.SetupRouter(cfg.ApiKey, cfg.SecretKey)

	log.Println("Сервер запущен на порту 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
