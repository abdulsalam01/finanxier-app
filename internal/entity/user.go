package entity

import "github.com/api-sekejap/internal/entity/base"

type User struct {
	ID int `json:"id"`

	Meta base.Metadata
}
