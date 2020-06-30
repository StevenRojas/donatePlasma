package transport

import (
	"context"
	"net/http"

	"github.com/StevenRojas/donatePlasma/services/matcher/pkg/endpoints"
	"github.com/StevenRojas/donatePlasma/services/matcher/pkg/reqres"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer Create new HTTP server instance
func NewHTTPServer(ctx context.Context, endpoints endpoints.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware)
	setPaths(r, endpoints)
	return r
}

func setPaths(r *mux.Router, endpoints endpoints.Endpoints) {
	// Get a list of public recipients with plasma compatibility
	r.Methods(http.MethodGet).Path("/api/matcher/recipients").Handler(httptransport.NewServer(
		endpoints.GetPublicRecipients,
		reqres.DecodePublicRecipientListRequest,
		reqres.EncodeResponse,
	))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, api-key")

			if r.Method == "OPTIONS" {
				return
			}
			next.ServeHTTP(w, r)
		})
}
