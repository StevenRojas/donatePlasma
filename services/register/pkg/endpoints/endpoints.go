package endpoints

import (
	"context"
	"errors"

	reqres "github.com/StevenRojas/donatePlasma/services/register/pkg/reqres"
	service "github.com/StevenRojas/donatePlasma/services/register/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints list of endpoints
type Endpoints struct {
	CreateRecipient     endpoint.Endpoint
	GetRecipients       endpoint.Endpoint
	GetPublicRecipients endpoint.Endpoint
	UpdateRecipient     endpoint.Endpoint
	VerifyRecipient     endpoint.Endpoint
	PublicRecipient     endpoint.Endpoint
	DeleteRecipient     endpoint.Endpoint
}

// MakeEndpoints create endpoints
func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		CreateRecipient:     makeCreateRecipientEndpoint(s),
		GetRecipients:       makeGetRecipientsEndpoint(s),
		GetPublicRecipients: makeGetPublicRecipientsEndpoint(s),
		UpdateRecipient:     makeUpdateRecipientEndpoint(s),
		VerifyRecipient:     makeVerifyRecipientEndpoint(s),
		PublicRecipient:     makePublicRecipientEndpoint(s),
		DeleteRecipient:     makeDeleteRecipientEndpoint(s),
	}
}

func makeCreateRecipientEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.CreateRecipientRequest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		id, err := s.CreateRecipient(ctx, req.Recipient)

		return reqres.CreateRecipientResponse{
			ID:  id,
			Err: err,
		}, err
	}
}

func makeGetRecipientsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		recipients, err := s.GetRecipientList(ctx, false)

		return reqres.GetRecipientsResponse{
			Recipients: recipients,
			Err:        err,
		}, err
	}
}

func makeGetPublicRecipientsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		recipients, err := s.GetRecipientList(ctx, true)

		return reqres.GetRecipientsResponse{
			Recipients: recipients,
			Err:        err,
		}, err
	}
}

func makeUpdateRecipientEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.UpdateRecipientRequest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		recipient, err := s.UpdateRecipient(ctx, req.Recipient)

		return reqres.UpdateRecipientResponse{
			Recipient: recipient,
			Err:       err,
		}, err
	}
}

func makeVerifyRecipientEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.VerifyRecipientResquest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		err := s.VerifyRecipient(ctx, req.ID, req.Verified)
		ok = (err == nil)
		return reqres.OkErrorResponse{
			Ok:  ok,
			Err: err,
		}, err
	}
}

func makePublicRecipientEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.PublicRecipientResquest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		err := s.PublicRecipient(ctx, req.ID, req.Public)
		ok = (err == nil)
		return reqres.OkErrorResponse{
			Ok:  ok,
			Err: err,
		}, err
	}
}

func makeDeleteRecipientEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.DeleteRecipientResquest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		err := s.DeleteRecipient(ctx, req.ID)
		ok = (err == nil)
		return reqres.OkErrorResponse{
			Ok:  ok,
			Err: err,
		}, err
	}
}
