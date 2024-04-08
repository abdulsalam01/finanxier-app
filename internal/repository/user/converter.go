package user

import (
	"github.com/finanxier-app/internal/entity"
	baseEnt "github.com/finanxier-app/internal/entity/base"
	"github.com/finanxier-app/internal/repository/base"
)

func (p *User) Normalize() entity.User {
	return entity.User{
		ID:       p.ID,
		Username: p.Username,
		Meta:     baseEnt.Metadata(p.Metadata),
		Extra:    baseEnt.ExtraAttribute(p.ExtraAttribute),
	}
}

func parserParams(params entity.User) User {
	hashed, err := params.HashPassword()
	if err != nil {
		hashed = params.PasswordHash
	}

	c := User{
		ID:             params.ID,
		Username:       params.Username,
		PasswordHash:   hashed,
		Metadata:       base.Metadata(params.Meta),
		ExtraAttribute: base.ExtraAttribute(params.Extra),
	}

	return c
}
