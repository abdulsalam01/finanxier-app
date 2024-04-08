package user

import (
	"context"

	"github.com/finanxier-app/internal/entity"
	"github.com/finanxier-app/pkg/database"
)

func (c *UsersRepo) GetByID(ctx context.Context, id string) (entity.User, error) {
	var (
		result entity.User
		user   User
		err    error
	)

	err = c.database.QueryRow(ctx, querySelectByID, id).
		Scan(
			&user.ID,
			&user.Username,
			&user.IsActive,
			&user.CreatedBy,
			&user.UpdatedBy,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	if err != nil {
		return result, database.WrapDuplicateKeyValueErr(err)
	}

	return user.Normalize(), nil
}

func (c *UsersRepo) GetFirstOne(ctx context.Context) (entity.User, error) {
	var (
		result entity.User
		user   User
		err    error
	)

	err = c.database.QueryRow(ctx, querySelectByParams, 1, 0).
		Scan(
			&user.ID,
			&user.Username,
			&user.IsActive,
			&user.CreatedBy,
			&user.UpdatedBy,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	if err != nil {
		return result, database.WrapDuplicateKeyValueErr(err)
	}

	return user.Normalize(), nil
}

func (c *UsersRepo) GetCurrentUser(ctx context.Context, token entity.User) (entity.User, error) {
	var (
		result entity.User
		user   User
		err    error
	)

	err = c.database.QueryRow(ctx, querySelectByUsername, token.Username).
		Scan(
			&user.ID,
			&user.Username,
			&user.IsActive,
			&user.CreatedBy,
			&user.UpdatedBy,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	if err != nil {
		return result, database.WrapDuplicateKeyValueErr(err)
	}

	return user.Normalize(), nil
}
