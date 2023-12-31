package entity

import "github.com/api-sekejap/internal/entity/base"

type SubscriptionPlan struct {
	ID        int `json:"id"`
	PackageID int `json:"package_id"`

	Meta base.Metadata
}
