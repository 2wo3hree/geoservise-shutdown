package router

import (
	"geoservise-jwt/internal/auth"
	"geoservise-jwt/internal/handler"
	"geoservise-jwt/internal/service"
	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
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

	auth.InitJWT()

	r.Post("/api/register", auth.RegisterHandler)
	r.Post("/api/login", auth.LoginHandler)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(auth.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/api/address/search", h.Search)
		r.Post("/api/address/geocode", h.Geocode)
	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
