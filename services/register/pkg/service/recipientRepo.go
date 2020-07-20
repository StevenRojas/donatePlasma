package service

import (
	"context"
	"strconv"
	"time"
)

// CreateRecipient creates a new recipient
func (repo *repository) CreateRecipient(ctx context.Context, recipient Recipient) (int64, error) {
	sql := `
	INSERT INTO recipient (id, blood_type_id, name, cell_numbers, email, photo_path, city_id, verified, public)
	VALUES (null, ?, ?, ?, ?, ?, ?, 1, ?);`
	stmt, err := repo.db.Prepare(sql)
	res, err := stmt.ExecContext(ctx, strconv.Itoa(recipient.BloodTypeID), recipient.Name, recipient.CellPhones, recipient.Email,
		recipient.PhotoPath, recipient.CityID, recipient.Public)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId() // Returns ID and err
}

// GetRecipientList get a list of verified and public recipients. TODO: Implement pagination
func (repo *repository) GetRecipientList(ctx context.Context, publicOnly bool) ([]Recipient, error) {
	var sql string
	if publicOnly {
		sql = `SELECT id, blood_type_id, name, cell_numbers, email, photo_path, city_id, verified, public, created_at, updated_at, deleted_at
		 FROM recipient WHERE public = 1 and verified = 1 and deleted_at IS NULL;`
	} else {
		sql = `SELECT id, blood_type_id, name, cell_numbers, email, photo_path, city_id, verified, public, created_at, updated_at, deleted_at
		 FROM recipient;`
	}

	rows, err := repo.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	list := []Recipient{}
	for rows.Next() {
		var recipient Recipient
		err = rows.Scan(&recipient.ID, &recipient.BloodTypeID, &recipient.Name, &recipient.CellPhones,
			&recipient.Email, &recipient.PhotoPath, &recipient.CityID, &recipient.Verified, &recipient.Public,
			&recipient.CreatedAt, &recipient.UpdatedAt, &recipient.DeletedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, recipient)
	}
	return list, nil
}

// UpdateRecipient update a recipient
func (repo *repository) UpdateRecipient(ctx context.Context, recipient Recipient) (Recipient, error) {
	sql := `
	UPDATE recipient SET blood_type_id = ?, name = ?, cell_numbers = ?, email = ?, photo_path = ?, 
		city_id = ?, public = ? WHERE id = ?;`
	stmt, err := repo.db.Prepare(sql)
	_, err = stmt.ExecContext(ctx, strconv.Itoa(recipient.BloodTypeID), recipient.Name, recipient.CellPhones, recipient.Email,
		recipient.PhotoPath, recipient.CityID, recipient.Public, recipient.ID)
	if err != nil {
		return recipient, err
	}
	return recipient, nil
}

// UpdateRecipient verify a recipient
func (repo *repository) VerifyRecipient(ctx context.Context, recipientID int64, verified bool) error {
	sql := `
	UPDATE recipient SET verified = ? WHERE id = ?;`
	stmt, err := repo.db.Prepare(sql)
	_, err = stmt.ExecContext(ctx, verified, recipientID)
	if err != nil {
		return err
	}
	return nil
}

// PublicRecipient set public flag for a recipient
func (repo *repository) PublicRecipient(ctx context.Context, recipientID int64, public bool) error {
	sql := `
	UPDATE recipient SET public = ? WHERE id = ?;`
	stmt, err := repo.db.Prepare(sql)
	_, err = stmt.ExecContext(ctx, public, recipientID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRecipient set delete a recipient
func (repo *repository) DeleteRecipient(ctx context.Context, recipientID int64, answer *bool, comment *string) error {
	sql := `
	UPDATE recipient SET deleted_at = ?, answer = ?, comment = ? WHERE id = ?;`
	stmt, err := repo.db.Prepare(sql)
	_, err = stmt.ExecContext(ctx, time.Now(), answer, comment, recipientID)
	if err != nil {
		return err
	}
	return nil
}

// ActivateRecipient set activate a recipient
func (repo *repository) ActivateRecipient(ctx context.Context, recipientID int64) error {
	sql := `
	UPDATE recipient SET deleted_at = null WHERE id = ?;`
	stmt, err := repo.db.Prepare(sql)
	_, err = stmt.ExecContext(ctx, recipientID)
	if err != nil {
		return err
	}
	return nil
}
