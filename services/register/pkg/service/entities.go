package service

// Recipient struct for recipients
type Recipient struct {
	ID          int64      `json:"id,omitempty"`
	BloodTypeID int        `json:"blood_type_id"`
	Name        string     `json:"name"`
	CellPhones  string     `json:"cell_phones"`
	Email       NullString `json:"email"`
	PhotoPath   NullString `json:"photo_path"`
	CityID      NullInt64  `json:"city_id"`
	Public      bool       `json:"public"`
	Verified    bool       `json:"verified"`
	Answer      bool       `json:"answer,omitempty"`
	Comment     string     `json:"comment,omitempty"`
	CreatedAt   NullTime   `json:"created_at"`
	UpdatedAt   NullTime   `json:"updated_at"`
	DeletedAt   NullTime   `json:"deleted_at"`
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
	Donations   string     `json:"quantity_donations"`
	Answer      bool       `json:"answer,omitempty"`
	Comment     string     `json:"comment,omitempty"`
	CreatedAt   NullTime   `json:"created_at"`
	UpdatedAt   NullTime   `json:"updated_at"`
	DeletedAt   NullTime   `json:"deleted_at"`
}
