package transport

import (
	"context"
	"net/http"
	"strings"

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
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			headers := []string{"Content-Type", "Accept", "Authorization", "token"}
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
			methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
			if r.Method == "OPTIONS" {
				return
			}
			next.ServeHTTP(w, r)
		})
}
