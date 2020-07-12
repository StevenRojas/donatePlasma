package transport

import (
	"context"
	"net/http"

	"github.com/StevenRojas/donatePlasma/services/register/pkg/endpoints"
	"github.com/gorilla/mux"
)

// NewHTTPServer Create new HTTP server instance
func NewHTTPServer(ctx context.Context, endpoints endpoints.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware)
	setRecipientPaths(r, endpoints)
	setDonorPaths(r, endpoints)
	return r
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			if r.Method == "OPTIONS" {
				return
			}
			next.ServeHTTP(w, r)
		})
}
