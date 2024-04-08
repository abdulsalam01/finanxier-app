package user

import (
	"context"

	"github.com/finanxier-app/internal/entity"
	"github.com/finanxier-app/pkg/database"
)

func (c *UsersRepo) Create(ctx context.Context, params entity.User) (entity.User, error) {
	var (
		result entity.User
		id     string
		err    error
	)

	paramsParser := parserParams(params)
	err = c.database.
		QueryRow(
			ctx,
			queryInsertWithReturning,
			&paramsParser.ID, &paramsParser.Username, &paramsParser.PasswordHash).
		Scan(&id)
	if err != nil {
		return result, database.WrapDuplicateKeyValueErr(err)
	}

	// Re-set the success data.
	result = entity.User{
		ID:       id,
		Username: params.Username,
	}
	return result, nil
}
