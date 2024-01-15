package entity

import "github.com/api-sekejap/internal/entity/base"

type Feature struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	Meta  base.Metadata
	Extra base.ExtraAttribute
}

type FeaturePoint struct {
	PackageID int `json:"package_id"`
	FeatureID int `json:"feature_id"`

	Meta  base.Metadata
	Extra base.ExtraAttribute
}
