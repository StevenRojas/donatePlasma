package endpoints

import (
	"context"
	"errors"

	"github.com/StevenRojas/donatePlasma/services/register/pkg/reqres"
	service "github.com/StevenRojas/donatePlasma/services/register/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

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

func makeActivateRecipientEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.ActivateRecipientResquest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		err := s.ActivateRecipient(ctx, req.ID)
		ok = (err == nil)
		return reqres.OkErrorResponse{
			Ok:  ok,
			Err: err,
		}, err
	}
}

func makeUploaderEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.UploaderResquest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		return reqres.UploaderResponse{
			Filename: req.Filename,
			Err:      nil,
		}, nil
	}
}

func makeCommentsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.CommentsResquest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		err := s.SendComment(ctx, req.Email, req.Comment, false)
		ok = (err == nil)
		return reqres.OkErrorResponse{
			Ok:  ok,
			Err: err,
		}, err
	}
}

func makeRecruitEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.CommentsResquest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		err := s.SendComment(ctx, req.Email, req.Comment, true)
		ok = (err == nil)
		return reqres.OkErrorResponse{
			Ok:  ok,
			Err: err,
		}, err
	}
}
