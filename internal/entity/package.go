package entity

import "github.com/api-sekejap/internal/entity/base"

type Package struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	AssetUrl    string `json:"asset_url"`

	Meta  base.Metadata
	Extra base.ExtraAttribute
}
