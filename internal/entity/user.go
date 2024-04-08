package entity

import (
	"github.com/finanxier-app/internal/constant"
	"github.com/finanxier-app/internal/entity/base"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash,omitempty"`

	Meta  base.Metadata       `json:"meta"`
	Extra base.ExtraAttribute `json:"extra"`
}

// HashPassword hashes a plain text password using bcrypt.
func (u *User) HashPassword() (string, error) {
	// GenerateFromPassword hashes the password using bcrypt.DefaultCost.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return constant.DefaultString, err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash compares a hashed password with a plain-text password to see if they match.
func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
