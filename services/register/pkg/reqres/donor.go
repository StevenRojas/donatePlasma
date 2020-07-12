package reqres

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	entities "github.com/StevenRojas/donatePlasma/services/register/pkg/service"
	"github.com/gorilla/mux"
)

// CreateDonorRequest create donor request
type CreateDonorRequest struct {
	Donor entities.Donor `json:"donor"`
}

// CreateDonorResponse create donor response
type CreateDonorResponse struct {
	ID  int64 `json:"id"`
	Err error `json:"error,omitempty"`
}

// GetDonorsResponse Get a list of donors
type GetDonorsResponse struct {
	Donors []entities.Donor `json:"donors"`
	Total  int64            `json:"total_records"`
	Err    error            `json:"error,omitempty"`
}

// UpdateDonorRequest update donor request
type UpdateDonorRequest struct {
	Donor entities.Donor `json:"donor"`
}

// UpdateDonorResponse create donor request
type UpdateDonorResponse struct {
	Donor entities.Donor `json:"donor"`
	Err   error          `json:"error,omitempty"`
}

// VerifyDonorResquest vefiry donor request
type VerifyDonorResquest struct {
	ID       int64 `json:"id,omitempty"`
	Verified bool  `json:"verified"`
}

// PublicDonorResquest public donor request
type PublicDonorResquest struct {
	ID     int64 `json:"id,omitempty"`
	Public bool  `json:"public"`
}

// DeleteDonorResquest delete donor request
type DeleteDonorResquest struct {
	ID int64 `json:"id,omitempty"`
}

// ActivateDonorResquest activate donor request
type ActivateDonorResquest struct {
	ID int64 `json:"id,omitempty"`
}

// GetRecipientsRequest recipient list request
type GetDonorsRequest struct {
	Query   string
	Page    int64
	PerPage int64
}

// DecodeCreateDonorRequest decode create donor request
func DecodeCreateDonorRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateDonorRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// DecodeUpdateDonorRequest decode update donor request
func DecodeUpdateDonorRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	if !validateAPIKey(r) {
		return nil, errors.New("Invalid access")
	}
	var req UpdateDonorRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		return nil, errors.New("Invalid donor ID")
	}
	req.Donor.ID = id
	return req, nil
}

// DecodeVerifyDonorRequest decode verify donor request
func DecodeVerifyDonorRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	if !validateAPIKey(r) {
		return nil, errors.New("Invalid access")
	}
	var req VerifyDonorResquest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		return nil, errors.New("Invalid donor ID")
	}
	req.ID = id
	return req, nil
}

// DecodePublicDonorRequest decode public donor request
func DecodePublicDonorRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	if !validateAPIKey(r) {
		return nil, errors.New("Invalid access")
	}
	var req PublicDonorResquest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		return nil, errors.New("Invalid donor ID")
	}
	req.ID = id
	return req, nil
}

// DecodeDeleteDonorRequest decode delete donor request
func DecodeDeleteDonorRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	if !validateAPIKey(r) {
		return nil, errors.New("Invalid access")
	}
	var req DeleteDonorResquest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		return nil, errors.New("Invalid donor ID")
	}
	req.ID = id
	return req, nil
}

// DecodeActivateDonorRequest decode activate donor request
func DecodeActivateDonorRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	if !validateAPIKey(r) {
		return nil, errors.New("Invalid access")
	}
	var req ActivateDonorResquest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		return nil, errors.New("Invalid donor ID")
	}
	req.ID = id
	return req, nil
}

// DecodePublicDonorListRequest decode public donor list request
func DecodePublicDonorListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetDonorsRequest

	vars := r.URL.Query()
	page, ok := getIntOrNull(vars, "page")
	if ok {
		req.Page = page
	} else {
		req.Page = 1
	}
	perPage, ok := getIntOrNull(vars, "per_page")
	if ok {
		req.PerPage = perPage
	} else {
		req.PerPage = 30
	}

	q, ok := vars["q"]
	if ok {
		req.Query = q[0]
	}
	return req, nil
}

func getIntOrNull(vars url.Values, field string) (int64, bool) {
	values, ok := vars[field]
	if ok {
		if len(values) >= 1 {
			value := values[0]
			if value != "" {
				id, err := strconv.ParseInt(value, 10, 32)
				if err == nil {
					return id, true
				}
			}
		}
	}
	return 0, false
}
