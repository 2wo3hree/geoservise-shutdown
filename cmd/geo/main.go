// @title GeoService API
// @version 1.0
// @description This is a simple geo service using DaData.
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
	_ "geoservise-jwt/docs"
	"geoservise-jwt/internal/config"
	"geoservise-jwt/internal/router"
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
