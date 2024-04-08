package product

import "github.com/finanxier-app/internal/repository/base"

const (
	tableDefinition = "products"

	// Select statements.
	querySelect      = `SELECT id, name, price, ` + base.GetBaseAttrQuery + `FROM ` + tableDefinition
	querySelectCount = `SELECT COUNT(*) FROM ` + tableDefinition

	querySelectByID     = querySelect + ` WHERE id = $1`
	querySelectByParams = querySelect + ` LIMIT $1 OFFSET $2`

	// Insert statements.
	queryInsert              = `INSERT INTO ` + tableDefinition + `(id, name, price) VALUES ($1, $2, $3)`
	queryInsertWithReturning = queryInsert + ` RETURNING id`
)
