package reqres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	entities "github.com/StevenRojas/donatePlasma/services/register/pkg/service"
	"github.com/gorilla/mux"
)

// CreateRecipientRequest create recipient request
type CreateRecipientRequest struct {
	Recipient entities.Recipient `json:"recipient"`
}

// CreateRecipientResponse create recipient response
type CreateRecipientResponse struct {
	ID  int64 `json:"id"`
	Err error `json:"error,omitempty"`
}

// GetRecipientsResponse Get a list of recipients
type GetRecipientsResponse struct {
	Recipients []entities.Recipient `json:"recipients"`
	Err        error                `json:"error,omitempty"`
}

// UpdateRecipientRequest update recipient request
type UpdateRecipientRequest struct {
	Recipient entities.Recipient `json:"recipient"`
}

// UpdateRecipientResponse create recipient request
type UpdateRecipientResponse struct {
	Recipient entities.Recipient `json:"recipient"`
	Err       error              `json:"error,omitempty"`
}

// VerifyRecipientResquest vefiry recipient request
type VerifyRecipientResquest struct {
	ID       int64 `json:"id,omitempty"`
	Verified bool  `json:"verified"`
}

// PublicRecipientResquest public recipient request
type PublicRecipientResquest struct {
	ID     int64 `json:"id,omitempty"`
	Public bool  `json:"public"`
}

// DeleteRecipientResquest delete recipient request
type DeleteRecipientResquest struct {
	ID      int64  `json:"id,omitempty"`
	Answer  bool   `json:"answer"`
	Comment string `json:"comment"`
}

// ActivateRecipientResquest activate recipient request
type ActivateRecipientResquest struct {
	ID int64 `json:"id,omitempty"`
}

// UploaderResquest uploader request
type UploaderResquest struct {
	Filename string `json:"filename"`
}

// UploaderResponse uploader request
type UploaderResponse struct {
	Filename string `json:"filename"`
	Err      error  `json:"error,omitempty"`
}

// CommentsResquest uploader request
type CommentsResquest struct {
	Email   string `json:"email"`
	Comment string `json:"comment"`
}

// DecodeCreateRecipientRequest decode create recipient request
func DecodeCreateRecipientRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateRecipientRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// DecodeUpdateRecipientRequest decode update recipient request
func DecodeUpdateRecipientRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	if !validateAPIKey(r) {
		return nil, errors.New("Invalid access")
	}
	var req UpdateRecipientRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		return nil, errors.New("Invalid recipient ID")
	}
	req.Recipient.ID = id
	return req, nil
}

// DecodeVerifyRecipientRequest decode verify recipient request
func DecodeVerifyRecipientRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	if !validateAPIKey(r) {
		return nil, errors.New("Invalid access")
	}
	var req VerifyRecipientResquest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		return nil, errors.New("Invalid recipient ID")
	}
	req.ID = id
	return req, nil
}

// DecodePublicRecipientRequest decode public recipient request
func DecodePublicRecipientRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	if !validateAPIKey(r) {
		return nil, errors.New("Invalid access")
	}
	var req PublicRecipientResquest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		return nil, errors.New("Invalid recipient ID")
	}
	req.ID = id
	return req, nil
}

// DecodeDeleteRecipientRequest decode delete recipient request
func DecodeDeleteRecipientRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	if !validateAPIKey(r) {
		return nil, errors.New("Invalid access")
	}
	var req DeleteRecipientResquest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		return nil, errors.New("Invalid recipient ID")
	}
	req.ID = id
	return req, nil
}

// DecodeActivateRecipientRequest decode activate recipient request
func DecodeActivateRecipientRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	if !validateAPIKey(r) {
		return nil, errors.New("Invalid access")
	}
	var req ActivateRecipientResquest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		return nil, errors.New("Invalid recipient ID")
	}
	req.ID = id
	return req, nil
}

// DecodeUploaderRequest decode uploader request
func DecodeUploaderRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UploaderResquest
	r.ParseMultipartForm(10 << 20) // Set 10Mb as max size
	file, handler, err := r.FormFile("file_uploader")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()
	contentType := handler.Header["Content-Type"][0]
	if !strings.Contains(contentType, "image/") {
		return nil, errors.New("Tipo de archivo no soportado")
	}
	ext := strings.Trim(filepath.Ext(handler.Filename), ".")
	tempFile, err := ioutil.TempFile("../../frontend/static/images", "img-*."+ext)
	tempFile.Chmod(0755)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = tempFile.Chmod(0755)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	tempFile.Write(fileBytes)

	req.Filename = filepath.Base(tempFile.Name())
	return req, nil
}

// DecodeCommentsRequest decode create recipient request
func DecodeCommentsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CommentsResquest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
