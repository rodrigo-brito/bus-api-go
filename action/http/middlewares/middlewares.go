package middlewares

import (
	"net/http"

	"github.com/rs/cors"
)

type Middleware func(handler http.Handler) http.Handler

func ApplyMiddlewares(handle http.Handler) http.Handler {
	middlewares := []Middleware{
		cors.Default().Handler,
		contextMiddleware,
	}
	for _, middleware := range middlewares {
		handle = middleware(handle)
	}
	return handle
}
