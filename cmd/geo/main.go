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
	"geoservise-jwt/internal/server"
	"geoservise-jwt/internal/shutdown"
)

func main() {
	cfg := config.LoadConfig()

	r := router.SetupRouter(cfg.ApiKey, cfg.SecretKey)

	s := server.NewGeoServer(":8080", r)
	shutdown.Gracefully(s)
}
