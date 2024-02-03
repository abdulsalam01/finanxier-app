package entity

import "github.com/api-sekejap/internal/entity/base"

type Channel struct {
	ID          int    `json:"id"`
	PackageID   int    `json:"package_id"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	AssetUrl    string `json:"asset_url"`
	Description string `json:"description"`

	base.Metadata
	base.ExtraAttribute
}
