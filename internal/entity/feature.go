package entity

import "github.com/api-sekejap/internal/entity/base"

type Feature struct {
	ID int `json:"id"`

	Meta base.Metadata
}

type FeaturePoint struct {
	ID        int `json:"id"`
	FeatureID int `json:"feature_id"`

	Meta base.Metadata
}
