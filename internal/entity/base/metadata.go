package base

import "time"

type Metadata struct {
	CreatedBy int       `json:"created_by"`
	UpdatedBy int       `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ExtraAttribute struct {
	IsActive bool `json:"is_active"`
}
