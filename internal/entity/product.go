package entity

import (
	"time"

	"github.com/api-sekejap/internal/entity/base"
)

type Product struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	ChannelID   int     `json:"channel_id"`
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Stock       int64   `json:"stock"`
	MinStock    int64   `json:"min_stock"` // For alert and warning when the product reach min-stock.
	Description string  `json:"description"`
	NormalPrice float64 `json:"normal_price"`

	PriceEvent []ProductPrice `json:"price_events"` // Holder for price event or discount.
	Asset      []ProductAsset `json:"assets"`
	Meta       base.Metadata
	Extra      base.ExtraAttribute
}

type ProductPrice struct {
	ID         int       `json:"id"`
	ProductID  int       `json:"product_id"`
	SlashPrice float64   `json:"slash_price"`
	EventPrice float64   `json:"event_price"`
	Discount   int       `json:"discount"`
	ExpiredAt  time.Time `json:"expired_at"`

	Meta  base.Metadata
	Extra base.ExtraAttribute
}

type ProductAsset struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Type      int    `json:"type"` // Indicate the asset type, could be: Cover, Slider, Etc.
	AssetUrl  string `json:"asset_url"`

	Meta base.Metadata
}
