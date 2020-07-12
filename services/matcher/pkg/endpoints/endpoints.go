package endpoints

import (
	"context"

	"github.com/StevenRojas/donatePlasma/services/matcher/pkg/reqres"
	service "github.com/StevenRojas/donatePlasma/services/matcher/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints list of endpoints
type Endpoints struct {
	GetPublicRecipients endpoint.Endpoint
	CanReceiveFrom      endpoint.Endpoint
	CanDonateTo         endpoint.Endpoint
	GetPublicDonors     endpoint.Endpoint
}

// MakeEndpoints create endpoints
func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		GetPublicRecipients: makeGetPublicRecipientsEndpoint(s),
		CanReceiveFrom:      makeCanReceiveFromEndpoint(s),
		CanDonateTo:         makeCanDonateToEndpoint(s),
		GetPublicDonors:     makeGetPublicDonorsEndpoint(s),
	}
}

func makeGetPublicRecipientsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, _ := request.(reqres.GetRecipientsRequest)
		recipients, count, err := s.GetRecipientList(ctx, req.CityID, req.BloodTypeID, req.Query, req.Page, req.PerPage)

		return reqres.GetRecipientsResponse{
			Recipients: recipients,
			Total:      count,
			Err:        err,
		}, err
	}
}

func makeCanReceiveFromEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, _ := request.(reqres.BloodTypeRequest)
		count, err := s.CanReceiveFrom(ctx, req.BloodTypeID)

		return reqres.CompatibleBloodTypeResponse{
			Compatible: count,
			Err:        err,
		}, err
	}
}

func makeCanDonateToEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, _ := request.(reqres.BloodTypeRequest)
		count, err := s.CanDonateTo(ctx, req.BloodTypeID)

		return reqres.CompatibleBloodTypeResponse{
			Compatible: count,
			Err:        err,
		}, err
	}
}

func makeGetPublicDonorsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, _ := request.(reqres.BloodTypeRequest)
		donors, err := s.GetDonorList(ctx, req.BloodTypeID)

		return reqres.GetDonorResponse{
			Donors: donors,
			Err:    err,
		}, err
	}
}
