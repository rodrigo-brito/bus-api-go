package middlewares

import (
	"net/http"

	"github.com/rodrigo-brito/bus-api-go/lib/context"
)

func contextMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.DefaultContext()
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
