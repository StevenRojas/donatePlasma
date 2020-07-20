package service

import (
	"context"
	"strconv"
	"time"
)

// CreateDonor creates a new donor
func (repo *repository) CreateDonor(ctx context.Context, donor Donor) (int64, error) {
	sql := `
	INSERT INTO donor (id, blood_type_id, name, cell, email, city_id, verified, public)
	VALUES (null, ?, ?, ?, ?, ?, 1, ?);`
	stmt, err := repo.db.Prepare(sql)
	res, err := stmt.ExecContext(ctx, strconv.Itoa(donor.BloodTypeID), donor.Name, donor.Cell, donor.Email,
		donor.CityID, donor.Public)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId() // Returns ID and err
}

// GetDonorList get a list of verified and public donors. TODO: Implement pagination
func (repo *repository) GetDonorList(ctx context.Context, publicOnly bool, q string, page int64, perPage int64) ([]Donor, int64, error) {
	start := (page - 1) * perPage
	var sql string
	if publicOnly {
		sql = `SELECT id, blood_type_id, name, cell, email, city_id, verified, public, created_at, updated_at, deleted_at
		 FROM donor WHERE public = 1 and verified = 1 and deleted_at IS NULL`
	} else {
		sql = `SELECT id, blood_type_id, name, cell, email, city_id, verified, public, created_at, updated_at, deleted_at
		 FROM donor`
	}
	if q != "" {
		sql = sql + " AND (name LIKE '%" + q + "%' OR cell LIKE '" + q + "%')"
	}

	sql = sql + " ORDER BY updated_at desc LIMIT " + strconv.FormatInt(start, 10) + ", " + strconv.FormatInt(perPage, 10)

	rows, err := repo.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, 0, err
	}
	list := []Donor{}
	for rows.Next() {
		var donor Donor
		err = rows.Scan(&donor.ID, &donor.BloodTypeID, &donor.Name, &donor.Cell,
			&donor.Email, &donor.CityID, &donor.Verified, &donor.Public,
			&donor.CreatedAt, &donor.UpdatedAt, &donor.DeletedAt)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, donor)
	}
	total, _ := repo.getTotalDonors(ctx, q)
	return list, total, nil
}

// UpdateDonor update a donor
func (repo *repository) UpdateDonor(ctx context.Context, donor Donor) (Donor, error) {
	sql := `
	UPDATE donor SET blood_type_id = ?, name = ?, cell = ?, email = ?, city_id = ?, public = ? WHERE id = ?;`
	stmt, err := repo.db.Prepare(sql)
	_, err = stmt.ExecContext(ctx, strconv.Itoa(donor.BloodTypeID), donor.Name, donor.Cell, donor.Email,
		donor.CityID, donor.Public, donor.ID)
	if err != nil {
		return donor, err
	}
	return donor, nil
}

// UpdateDonor verify a donor
func (repo *repository) VerifyDonor(ctx context.Context, donorID int64, verified bool) error {
	sql := `
	UPDATE donor SET verified = ? WHERE id = ?;`
	stmt, err := repo.db.Prepare(sql)
	_, err = stmt.ExecContext(ctx, verified, donorID)
	if err != nil {
		return err
	}
	return nil
}

// PublicDonor set public flag for a donor
func (repo *repository) PublicDonor(ctx context.Context, donorID int64, public bool) error {
	sql := `
	UPDATE donor SET public = ? WHERE id = ?;`
	stmt, err := repo.db.Prepare(sql)
	_, err = stmt.ExecContext(ctx, public, donorID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDonor set delete a donor
func (repo *repository) DeleteDonor(ctx context.Context, donorID int64, answer *bool, comment *string) error {
	sql := `
	UPDATE donor SET deleted_at = ?, answer = ?, comment = ? WHERE id = ?;`
	stmt, err := repo.db.Prepare(sql)
	_, err = stmt.ExecContext(ctx, time.Now(), answer, comment, donorID)
	if err != nil {
		return err
	}
	return nil
}

// ActivateDonor set activate a donor
func (repo *repository) ActivateDonor(ctx context.Context, donorID int64) error {
	sql := `
	UPDATE donor SET deleted_at = null WHERE id = ?;`
	stmt, err := repo.db.Prepare(sql)
	_, err = stmt.ExecContext(ctx, donorID)
	if err != nil {
		return err
	}
	return nil
}

// getTotalDonors get a count of verified and public recipients
func (repo *repository) getTotalDonors(ctx context.Context, q string) (int64, error) {
	var sql string
	sql = `SELECT count(id) as c FROM donor WHERE public = 1 AND verified = 1 AND deleted_at IS NULL`
	if q != "" {
		sql = sql + " AND (name LIKE '%" + q + "%' OR cell LIKE '" + q + "%')"
	}
	result, err := repo.db.QueryContext(ctx, sql)
	if err != nil {
		return 0, err
	}
	var count int64
	result.Next()
	err = result.Scan(&count)
	return count, nil
}
