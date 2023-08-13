package middleware

import (
	"net/http"
	"os"

	"github.com/rs/cors"
)

type CorsMiddleware struct{}

func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}

func (middleware *CorsMiddleware) MiddlewareFunc(next http.Handler) http.Handler {

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONTEND_DOMAIN")},
		AllowCredentials: true,
		Debug:            true,
	})

	return c.Handler(next)
}
