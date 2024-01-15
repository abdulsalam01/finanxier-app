package entity

import "github.com/api-sekejap/internal/entity/base"

type SubscriptionPlan struct {
	ID                   int     `json:"id"`
	PackageID            int     `json:"package_id"`
	Type                 int     `json:"type"` // Indicate the type of package, could be: Basic, Silver, Premium, etc.
	Price                float64 `json:"price"`
	IsFreeTrialAvailable bool    `json:"is_free_trial_available"`
	Period               string  `json:"period"` // Could be: Monthly, 3/6 Month, Annualy.
	PeriodFreeTrial      int     `json:"period_free_trial"`

	Meta  base.Metadata
	Extra base.ExtraAttribute
}
