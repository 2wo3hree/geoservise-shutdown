package router

import (
	"geoservise/internal/handler"
	"geoservise/internal/service"
	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
)

func SetupRouter(apiKey, secretKey string) *chi.Mux {
	creds := client.Credentials{
		ApiKeyValue:    apiKey,
		SecretKeyValue: secretKey,
	}
	api := dadata.NewSuggestApi(
		client.WithCredentialProvider(&creds),
	)
	if api == nil {
		log.Fatal("Dadata API не инициализирован")
	}

	s := service.NewService(api)
	h := handler.NewAddressHandler(s)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/address/search", h.Search)
	r.Post("/api/address/geocode", h.Geocode)

	return r
}
