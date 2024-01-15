package entity

import "github.com/api-sekejap/internal/entity/base"

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	Salt         string `json:"salt"`
	IsVerified   bool   `json:"is_verified"`

	Social  []UserSocialAccount  `json:"social_accounts"`
	Payment []UserPaymentAccount `json:"payment_accounts"`
	Meta    base.Metadata
	Extra   base.ExtraAttribute
}

type UserSocialAccount struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Provider string `json:"provider"`
}

type UserPaymentAccount struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`

	Meta  base.Metadata
	Extra base.ExtraAttribute
}
