package entity

import "github.com/api-sekejap/internal/entity/base"

type SubscriptionPlan struct {
	ID        int `json:"id"`
	FeatureID int `json:"feature_id"`

	Meta base.Metadata
}
