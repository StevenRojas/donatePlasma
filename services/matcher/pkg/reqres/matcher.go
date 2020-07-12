package reqres

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	entities "github.com/StevenRojas/donatePlasma/services/matcher/pkg/service"
	"github.com/gorilla/mux"
)

// GetRecipientsRequest recipient list request
type GetRecipientsRequest struct {
	BloodTypeID *int64
	CityID      *int64
	Query       string
	Page        int64
	PerPage     int64
}

// GetRecipientsResponse Get a list of recipients
type GetRecipientsResponse struct {
	Recipients []entities.Recipient `json:"recipients"`
	Total      int64                `json:"total_records"`
	Err        error                `json:"error,omitempty"`
}

// BloodTypeRequest blood type request
type BloodTypeRequest struct {
	BloodTypeID int64
}

// CompatibleBloodTypeResponse blood type request
type CompatibleBloodTypeResponse struct {
	Compatible []entities.CompatibleBloodCount `json:"compatible_types"`
	Err        error                           `json:"error,omitempty"`
}

// GetDonorResponse Get a list of recipients
type GetDonorResponse struct {
	Donors []entities.Donor `json:"donors"`
	Err    error            `json:"error,omitempty"`
}

// DecodeBloodTypeRequest decode blood type request
func DecodeBloodTypeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req BloodTypeRequest

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["typeId"], 10, 32)
	if err != nil {
		return nil, errors.New("Invalid blood type ID")
	}
	req.BloodTypeID = id
	return req, nil
}

// DecodePublicRecipientListRequest decode public recipient list request
func DecodePublicRecipientListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetRecipientsRequest

	vars := r.URL.Query()
	cityID, ok := getIntOrNull(vars, "city")
	if ok {
		req.CityID = &cityID
	}
	bloodTypeID, ok := getIntOrNull(vars, "compatible_with")
	if ok {
		req.BloodTypeID = &bloodTypeID
	}
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
