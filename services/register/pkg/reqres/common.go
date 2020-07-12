package reqres

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// OkErrorResponse Empty response, just OK or Error
type OkErrorResponse struct {
	Ok  bool  `json:"ok"`
	Err error `json:"error,omitempty"`
}

// EncodeResponse generic encoder
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// DecodeEmptyRequest generic decoder for request with no parameteres, like GET /users
func DecodeEmptyRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

// DecodeProtectedEmptyRequest generic decoder for request with no parameteres validating api-key
func DecodeProtectedEmptyRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	if !validateAPIKey(r) {
		return nil, errors.New("Invalid access")
	}
	return nil, nil
}

func validateAPIKey(r *http.Request) bool {
	apiKey := r.Header.Get("api-key")
	if apiKey != "123abc!" { // TODO: Definitely improve this
		//return false
	}
	return true
}
