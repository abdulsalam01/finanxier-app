package user

import "github.com/finanxier-app/internal/repository/base"

const (
	tableDefinition = "users"

	// Select statements.
	querySelect      = `SELECT id, username, ` + base.GetBaseAttrQuery + `FROM ` + tableDefinition
	querySelectCount = `SELECT COUNT(*) FROM ` + tableDefinition

	querySelectByID       = querySelect + ` WHERE id = $1`
	querySelectByUsername = querySelect + ` WHERE username = $1` // Indexing on Column.
	querySelectByParams   = querySelect + ` LIMIT $1 OFFSET $2`

	// Insert statements.
	queryInsert              = `INSERT INTO ` + tableDefinition + `(id, username, password_hash) VALUES ($1, $2, $3)`
	queryInsertWithReturning = queryInsert + ` RETURNING id`
)
