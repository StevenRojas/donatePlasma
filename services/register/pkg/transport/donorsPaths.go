package transport

import (
	"net/http"

	"github.com/StevenRojas/donatePlasma/services/register/pkg/endpoints"
	"github.com/StevenRojas/donatePlasma/services/register/pkg/reqres"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func setDonorPaths(r *mux.Router, endpoints endpoints.Endpoints) {
	// Create a donor record
	r.Methods(http.MethodPost).Path("/api/register/donor").Handler(httptransport.NewServer(
		endpoints.CreateDonor,
		reqres.DecodeCreateDonorRequest,
		reqres.EncodeResponse,
	))

	// Get a list of all donors
	r.Methods(http.MethodGet).Path("/api/register/donor").Handler(httptransport.NewServer(
		endpoints.GetDonors,
		reqres.DecodePublicDonorListRequest,
		reqres.EncodeResponse,
	))

	// Get a list of public donors
	r.Methods(http.MethodGet).Path("/api/register/donor/public").Handler(httptransport.NewServer(
		endpoints.GetPublicDonors,
		reqres.DecodePublicDonorListRequest,
		reqres.EncodeResponse,
	))

	// Update a donor
	r.Methods(http.MethodPut).Path("/api/register/donor/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateDonor,
		reqres.DecodeUpdateDonorRequest,
		reqres.EncodeResponse,
	))

	// Verify/Unverify a donor
	r.Methods(http.MethodPatch).Path("/api/register/donor/{id}/verify").Handler(httptransport.NewServer(
		endpoints.VerifyDonor,
		reqres.DecodeVerifyDonorRequest,
		reqres.EncodeResponse,
	))

	// Set as public or not public a donor
	r.Methods(http.MethodPatch).Path("/api/register/donor/{id}/public").Handler(httptransport.NewServer(
		endpoints.PublicDonor,
		reqres.DecodePublicDonorRequest,
		reqres.EncodeResponse,
	))

	// Delete a donor
	r.Methods(http.MethodPatch).Path("/api/register/donor/{id}/delete").Handler(httptransport.NewServer(
		endpoints.DeleteDonor,
		reqres.DecodeDeleteDonorRequest,
		reqres.EncodeResponse,
	))

	// Activate a donor
	r.Methods(http.MethodPatch).Path("/api/register/donor/{id}/activate").Handler(httptransport.NewServer(
		endpoints.ActivateDonor,
		reqres.DecodeActivateDonorRequest,
		reqres.EncodeResponse,
	))
}
