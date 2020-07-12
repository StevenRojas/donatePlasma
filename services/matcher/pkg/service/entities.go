package service

// Recipient struct for recipients
type Recipient struct {
	ID              int64                  `json:"id,omitempty"`
	BloodTypeID     int                    `json:"blood_type_id"`
	Name            string                 `json:"name"`
	CellPhones      string                 `json:"cell_phones"`
	Email           NullString             `json:"email"`
	PhotoPath       NullString             `json:"photo_path"`
	CityID          NullInt64              `json:"city_id"`
	Public          bool                   `json:"public"`
	Verified        bool                   `json:"verified"`
	CreatedAt       NullTime               `json:"created_at"`
	UpdatedAt       NullTime               `json:"updated_at"`
	DeletedAt       NullTime               `json:"deleted_at"`
	PotentialDonors []CompatibleBloodCount `json:"potential_donors,omitempty"`
}

// Donor structure for donors
type Donor struct {
	ID          int64      `json:"id,omitempty"`
	BloodTypeID int        `json:"blood_type_id"`
	Name        string     `json:"name"`
	Cell        string     `json:"cell"`
	Email       NullString `json:"email"`
	CityID      NullInt64  `json:"city_id"`
	Public      bool       `json:"public"`
	Verified    bool       `json:"verified"`
	CreatedAt   NullTime   `json:"created_at"`
	UpdatedAt   NullTime   `json:"updated_at"`
	DeletedAt   NullTime   `json:"deleted_at"`
}

// CompatibleBloodCount counts for a given blood type
type CompatibleBloodCount struct {
	BloodTypeID int64 `json:"blood_type_id"`
	Count       int64 `json:"count"`
}
