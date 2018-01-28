package middlewares

import (
	"net/http"

	"github.com/rodrigo-brito/bus-api-go/lib/context"
)

func contextMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		BPC := false
		r.ParseForm()
		if bpc := r.FormValue("bpc"); len(bpc) != 0 {
			BPC = true
		}
		ctx := context.DefaultContext(BPC)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
