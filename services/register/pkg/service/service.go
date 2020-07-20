package service

import (
	"context"
	"net/smtp"
	"os"

	"github.com/go-kit/kit/log"
)

// Service definition
type Service interface {
	CreateRecipient(ctx context.Context, recipient Recipient) (int64, error)
	GetRecipientList(ctx context.Context, publicOnly bool) ([]Recipient, error)
	UpdateRecipient(ctx context.Context, recipient Recipient) (Recipient, error)
	VerifyRecipient(ctx context.Context, recipientID int64, verified bool) error
	PublicRecipient(ctx context.Context, recipientID int64, public bool) error
	DeleteRecipient(ctx context.Context, recipientID int64, answer *bool, comment *string) error
	ActivateRecipient(ctx context.Context, recipientID int64) error
	SendComment(ctx context.Context, email string, comment string, isRecluter bool) error

	CreateDonor(ctx context.Context, donor Donor) (int64, error)
	GetDonorList(ctx context.Context, publicOnly bool, q string, page int64, perPage int64) ([]Donor, int64, error)
	UpdateDonor(ctx context.Context, donor Donor) (Donor, error)
	VerifyDonor(ctx context.Context, donorID int64, verified bool) error
	PublicDonor(ctx context.Context, donorID int64, public bool) error
	DeleteDonor(ctx context.Context, donorID int64, answer *bool, comment *string) error
	ActivateDonor(ctx context.Context, donorID int64) error
}

type service struct {
	repository Repository
	logger     log.Logger
}

// NewService create new service instance
func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s service) CreateRecipient(ctx context.Context, recipient Recipient) (int64, error) {
	logger := log.With(s.logger, "msg", "CreateRecipient")

	logger.Log("msg", "Creating Recipient")
	response, err := s.repository.CreateRecipient(ctx, recipient)
	return response, err
}

func (s service) GetRecipientList(ctx context.Context, publicOnly bool) ([]Recipient, error) {
	logger := log.With(s.logger, "msg", "GetRecipientList")

	logger.Log("msg", "Getting Recipient list")
	response, err := s.repository.GetRecipientList(ctx, publicOnly)
	return response, err
}

func (s service) UpdateRecipient(ctx context.Context, recipient Recipient) (Recipient, error) {
	logger := log.With(s.logger, "msg", "UpdateRecipient")

	logger.Log("msg", "Updating Recipient")
	response, err := s.repository.UpdateRecipient(ctx, recipient)
	return response, err
}

func (s service) VerifyRecipient(ctx context.Context, recipientID int64, verified bool) error {
	logger := log.With(s.logger, "msg", "VerifyRecipient")

	logger.Log("msg", "Verifiying Recipient")
	err := s.repository.VerifyRecipient(ctx, recipientID, verified)
	return err
}

func (s service) PublicRecipient(ctx context.Context, recipientID int64, public bool) error {
	logger := log.With(s.logger, "msg", "PublicRecipient")

	logger.Log("msg", "Setting public Recipient")
	err := s.repository.PublicRecipient(ctx, recipientID, public)
	return err
}

func (s service) DeleteRecipient(ctx context.Context, recipientID int64, answer *bool, comment *string) error {
	logger := log.With(s.logger, "msg", "DeleteRecipient")

	logger.Log("msg", "Deleting Recipient")
	err := s.repository.DeleteRecipient(ctx, recipientID, answer, comment)
	return err
}

func (s service) ActivateRecipient(ctx context.Context, recipientID int64) error {
	logger := log.With(s.logger, "msg", "ActivateRecipient")

	logger.Log("msg", "Activating Recipient")
	err := s.repository.ActivateRecipient(ctx, recipientID)
	return err
}

func (s service) CreateDonor(ctx context.Context, donor Donor) (int64, error) {
	logger := log.With(s.logger, "msg", "CreateDonor")

	logger.Log("msg", "Creating Donor")
	response, err := s.repository.CreateDonor(ctx, donor)
	return response, err
}

func (s service) GetDonorList(ctx context.Context, publicOnly bool, q string, page int64, perPage int64) ([]Donor, int64, error) {
	logger := log.With(s.logger, "msg", "GetDonorList")

	logger.Log("msg", "Getting Donor list")
	response, total, err := s.repository.GetDonorList(ctx, publicOnly, q, page, perPage)
	return response, total, err
}

func (s service) UpdateDonor(ctx context.Context, donor Donor) (Donor, error) {
	logger := log.With(s.logger, "msg", "UpdateDonor")

	logger.Log("msg", "Updating Donor")
	response, err := s.repository.UpdateDonor(ctx, donor)
	return response, err
}

func (s service) VerifyDonor(ctx context.Context, donorID int64, verified bool) error {
	logger := log.With(s.logger, "msg", "VerifyDonor")

	logger.Log("msg", "Verifiying Donor")
	err := s.repository.VerifyDonor(ctx, donorID, verified)
	return err
}

func (s service) PublicDonor(ctx context.Context, donorID int64, public bool) error {
	logger := log.With(s.logger, "msg", "PublicDonor")

	logger.Log("msg", "Setting public Donor")
	err := s.repository.PublicDonor(ctx, donorID, public)
	return err
}

func (s service) DeleteDonor(ctx context.Context, donorID int64, answer *bool, comment *string) error {
	logger := log.With(s.logger, "msg", "DeleteDonor")

	logger.Log("msg", "Deleting Donor")
	err := s.repository.DeleteDonor(ctx, donorID, answer, comment)
	return err
}

func (s service) ActivateDonor(ctx context.Context, donorID int64) error {
	logger := log.With(s.logger, "msg", "ActivateDonor")

	logger.Log("msg", "Activating Donor")
	err := s.repository.ActivateDonor(ctx, donorID)
	return err
}

func (s service) SendComment(ctx context.Context, email string, comment string, isRecluter bool) error {
	logger := log.With(s.logger, "msg", "SendComment")
	var subject string
	if isRecluter {
		subject = "Subject: Nuevo recluta\n"
	} else {
		subject = "Subject: Nuevo commentario\n"
	}
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	from, _ := os.LookupEnv("EMAIL_FROM")
	pass, _ := os.LookupEnv("EMAIL_PASS")
	to, _ := os.LookupEnv("EMAIL_TO")
	body := comment + "<br><br>" + email
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		subject + mime + body
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
	if err != nil {
		logger.Log("msg", "Error sending comment sent "+err.Error())
	}
	logger.Log("msg", "Comment sent")
	return nil
}
