package entity

import "github.com/api-sekejap/internal/entity/base"

type Feature struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	base.Metadata
	base.ExtraAttribute
}

type FeaturePoint struct {
	PackageID int `json:"package_id"`
	FeatureID int `json:"feature_id"`

	base.Metadata
	base.ExtraAttribute
}
