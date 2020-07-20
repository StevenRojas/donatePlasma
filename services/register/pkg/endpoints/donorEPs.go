package endpoints

import (
	"context"
	"errors"

	"github.com/StevenRojas/donatePlasma/services/register/pkg/reqres"
	service "github.com/StevenRojas/donatePlasma/services/register/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

func makeCreateDonorEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.CreateDonorRequest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		id, err := s.CreateDonor(ctx, req.Donor)

		return reqres.CreateDonorResponse{
			ID:  id,
			Err: err,
		}, err
	}
}

func makeGetDonorsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.GetDonorsRequest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		donors, total, err := s.GetDonorList(ctx, false, req.Query, req.Page, req.PerPage)

		return reqres.GetDonorsResponse{
			Donors: donors,
			Total:  total,
			Err:    err,
		}, err
	}
}

func makeGetPublicDonorsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.GetDonorsRequest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		donors, total, err := s.GetDonorList(ctx, true, req.Query, req.Page, req.PerPage)

		return reqres.GetDonorsResponse{
			Donors: donors,
			Total:  total,
			Err:    err,
		}, err
	}
}

func makeUpdateDonorEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.UpdateDonorRequest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		donor, err := s.UpdateDonor(ctx, req.Donor)

		return reqres.UpdateDonorResponse{
			Donor: donor,
			Err:   err,
		}, err
	}
}

func makeVerifyDonorEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.VerifyDonorResquest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		err := s.VerifyDonor(ctx, req.ID, req.Verified)
		ok = (err == nil)
		return reqres.OkErrorResponse{
			Ok:  ok,
			Err: err,
		}, err
	}
}

func makePublicDonorEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.PublicDonorResquest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		err := s.PublicDonor(ctx, req.ID, req.Public)
		ok = (err == nil)
		return reqres.OkErrorResponse{
			Ok:  ok,
			Err: err,
		}, err
	}
}

func makeDeleteDonorEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.DeleteDonorResquest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		err := s.DeleteDonor(ctx, req.ID, &req.Answer, &req.Comment)
		ok = (err == nil)
		return reqres.OkErrorResponse{
			Ok:  ok,
			Err: err,
		}, err
	}
}

func makeActivateDonorEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(reqres.ActivateDonorResquest)
		if !ok {
			return nil, errors.New("Wrong request message")
		}

		err := s.ActivateDonor(ctx, req.ID)
		ok = (err == nil)
		return reqres.OkErrorResponse{
			Ok:  ok,
			Err: err,
		}, err
	}
}
