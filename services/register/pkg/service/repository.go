package service

import (
	"context"
	"database/sql"

	"github.com/go-kit/kit/log"
)

// Repository interface
type Repository interface {
	CreateRecipient(ctx context.Context, recipient Recipient) (int64, error)
	GetRecipientList(ctx context.Context, publicOnly bool) ([]Recipient, error)
	UpdateRecipient(ctx context.Context, recipient Recipient) (Recipient, error)
	VerifyRecipient(ctx context.Context, recipientID int64, verified bool) error
	PublicRecipient(ctx context.Context, recipientID int64, public bool) error
	DeleteRecipient(ctx context.Context, recipientID int64, answer *bool, comment *string) error
	ActivateRecipient(ctx context.Context, recipientID int64) error

	CreateDonor(ctx context.Context, donor Donor) (int64, error)
	GetDonorList(ctx context.Context, publicOnly bool, q string, page int64, perPage int64) ([]Donor, int64, error)
	UpdateDonor(ctx context.Context, donor Donor) (Donor, error)
	VerifyDonor(ctx context.Context, donorID int64, verified bool) error
	PublicDonor(ctx context.Context, donorID int64, public bool) error
	DeleteDonor(ctx context.Context, donorID int64, answer *bool, comment *string) error
	ActivateDonor(ctx context.Context, donorID int64) error
}

type repository struct {
	db     *sql.DB
	logger log.Logger
}

// NewRepository creates a new repository instance
func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &repository{
		db:     db,
		logger: log.With(logger, "Repository", "sql"),
	}
}
