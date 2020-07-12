package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"

	"github.com/go-kit/kit/log"
)

// Repository interface
type Repository interface {
	GetRecipientList(ctx context.Context, cityID *int64, bloodTypeID *int64, q string, page int64, perPage int64) ([]Recipient, int64, error)
	CanReceiveFrom(ctx context.Context, bloodTypeID int64) ([]CompatibleBloodCount, error)
	CanDonateTo(ctx context.Context, bloodTypeID int64) ([]CompatibleBloodCount, error)
	GetDonorList(ctx context.Context, bloodTypeID int64) ([]Donor, error)
	getTotalRecipients(ctx context.Context, cityID *int64, bloodTypeID *int64, q string) (int64, error)
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

// GetRecipientList get a list of verified and public recipients
func (repo *repository) GetRecipientList(ctx context.Context, cityID *int64, bloodTypeID *int64, q string, page int64, perPage int64) ([]Recipient, int64, error) {
	start := (page - 1) * perPage
	var sql string
	sql = `SELECT id, blood_type_id, name, cell_numbers, email, photo_path, city_id, verified, public, created_at, updated_at, deleted_at,
			(SELECT CONCAT('[', GROUP_CONCAT(JSON_OBJECT(
				'blood_type_id', donor_blood_type_id,
				'count', (SELECT COUNT(id) FROM donor WHERE blood_type_id = donor_blood_type_id AND public = 1 AND verified = 1 AND deleted_at))),
				']') AS donors FROM compatibility WHERE recipient_blood_type_id = blood_type_id) AS donors
			FROM recipient WHERE public = 1 AND verified = 1 AND deleted_at IS NULL`
	if cityID != nil {
		sql = sql + " AND city_id = " + strconv.FormatInt(*cityID, 10)
	}
	if bloodTypeID != nil {
		sql = sql + " AND blood_type_id = " + strconv.FormatInt(*bloodTypeID, 10)
	}
	if q != "" {
		sql = sql + " AND (name LIKE '%" + q + "%' OR cell_numbers LIKE '" + q + "%')"
	}

	sql = sql + " ORDER BY updated_at desc LIMIT " + strconv.FormatInt(start, 10) + ", " + strconv.FormatInt(perPage, 10)
	rows, err := repo.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, 0, err
	}

	list := []Recipient{}
	for rows.Next() {
		var recipient Recipient
		var potentialDonors string
		err = rows.Scan(&recipient.ID, &recipient.BloodTypeID, &recipient.Name, &recipient.CellPhones,
			&recipient.Email, &recipient.PhotoPath, &recipient.CityID, &recipient.Verified, &recipient.Public,
			&recipient.CreatedAt, &recipient.UpdatedAt, &recipient.DeletedAt, &potentialDonors)
		if err != nil {
			return nil, 0, err
		}
		donors := make([]CompatibleBloodCount, 0)
		json.Unmarshal([]byte(potentialDonors), &donors)
		recipient.PotentialDonors = donors
		list = append(list, recipient)
	}
	total, _ := repo.getTotalRecipients(ctx, cityID, bloodTypeID, q)
	return list, total, nil
}

// CanReceiveFrom returns donor counts for a recipient
func (repo *repository) CanReceiveFrom(ctx context.Context, bloodTypeID int64) ([]CompatibleBloodCount, error) {
	sql := `SELECT donor_blood_type_id, (SELECT count(id) FROM donor WHERE blood_type_id = donor_blood_type_id 
		AND public = 1 AND verified = 1 AND deleted_at) as c 
		FROM compatibility WHERE recipient_blood_type_id = ` + strconv.FormatInt(bloodTypeID, 10)
	rows, err := repo.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	list := []CompatibleBloodCount{}
	for rows.Next() {
		var count CompatibleBloodCount
		err = rows.Scan(&count.BloodTypeID, &count.Count)
		if err != nil {
			return nil, err
		}
		list = append(list, count)
	}
	return list, nil
}

// CanDonateTo returns recipient counts for a donor
func (repo *repository) CanDonateTo(ctx context.Context, bloodTypeID int64) ([]CompatibleBloodCount, error) {
	sql := `SELECT recipient_blood_type_id, (SELECT count(id) FROM recipient WHERE blood_type_id = recipient_blood_type_id 
		AND public = 1 AND verified = 1 AND deleted_at) as c 
		FROM compatibility WHERE donor_blood_type_id= ` + strconv.FormatInt(bloodTypeID, 10)
	rows, err := repo.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	list := []CompatibleBloodCount{}
	for rows.Next() {
		var count CompatibleBloodCount
		err = rows.Scan(&count.BloodTypeID, &count.Count)
		if err != nil {
			return nil, err
		}
		list = append(list, count)
	}
	return list, nil
}

// GetDonorList returns public donors compatible with a blood type
func (repo *repository) GetDonorList(ctx context.Context, bloodTypeID int64) ([]Donor, error) {
	sql := `SELECT * FROM donor WHERE blood_type_id = ` + strconv.FormatInt(bloodTypeID, 10) + `
		AND public = 1 AND verified = 1 AND deleted_at IS NULL ORDER BY updated_at DESC`
	rows, err := repo.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	list := []Donor{}
	for rows.Next() {
		var donor Donor
		err = rows.Scan(&donor.ID, &donor.BloodTypeID, &donor.Name, &donor.Cell,
			&donor.Email, &donor.CityID, &donor.Verified, &donor.Public,
			&donor.CreatedAt, &donor.UpdatedAt, &donor.DeletedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, donor)
	}
	return list, nil
}

// getTotalRecipients get a count of verified and public recipients
func (repo *repository) getTotalRecipients(ctx context.Context, cityID *int64, bloodTypeID *int64, q string) (int64, error) {
	var sql string
	sql = `SELECT count(id) as c FROM recipient WHERE public = 1 AND verified = 1 AND deleted_at IS NULL`
	if cityID != nil {
		sql = sql + " AND city_id = " + strconv.FormatInt(*cityID, 10)
	}
	if bloodTypeID != nil {
		sql = sql + " AND blood_type_id IN (SELECT recipient_blood_type_id FROM compatibility WHERE donor_blood_type_id = " +
			strconv.FormatInt(*bloodTypeID, 10) + ")"
	}
	if q != "" {
		sql = sql + " AND (name LIKE '%" + q + "%' OR cell_numbers LIKE '" + q + "%')"
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
