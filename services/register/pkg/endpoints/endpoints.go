package endpoints

import (
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
	ActivateRecipient   endpoint.Endpoint
	Uploader            endpoint.Endpoint
	Comments            endpoint.Endpoint
	Recruit             endpoint.Endpoint

	CreateDonor     endpoint.Endpoint
	GetDonors       endpoint.Endpoint
	GetPublicDonors endpoint.Endpoint
	UpdateDonor     endpoint.Endpoint
	VerifyDonor     endpoint.Endpoint
	PublicDonor     endpoint.Endpoint
	DeleteDonor     endpoint.Endpoint
	ActivateDonor   endpoint.Endpoint
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
		ActivateRecipient:   makeActivateRecipientEndpoint(s),
		Uploader:            makeUploaderEndpoint(s),
		Comments:            makeCommentsEndpoint(s),
		Recruit:             makeRecruitEndpoint(s),

		CreateDonor:     makeCreateDonorEndpoint(s),
		GetDonors:       makeGetDonorsEndpoint(s),
		GetPublicDonors: makeGetPublicDonorsEndpoint(s),
		UpdateDonor:     makeUpdateDonorEndpoint(s),
		VerifyDonor:     makeVerifyDonorEndpoint(s),
		PublicDonor:     makePublicDonorEndpoint(s),
		DeleteDonor:     makeDeleteDonorEndpoint(s),
		ActivateDonor:   makeActivateDonorEndpoint(s),
	}
}
