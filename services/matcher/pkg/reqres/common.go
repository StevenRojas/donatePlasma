package reqres

import (
	"context"
	"encoding/json"
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
